package framework

import (
	"errors"
	"sync"
)

// Container define container is server container
type Container interface {
	// Bind a service provider if keyword exists will replace and return error
	Bind(provider ServiceProvider) error

	// IsBind identify check service provider has been bound?
	IsBind(key string) bool

	// Make service by ke identify
	Make(key string) (interface{}, error)

	// MustMake make service by key identify, if key does not bound service provider will throw panic
	// So using the interface must be sure keep service container already assign key identify bind service provider
	MustMake(key string) interface{}

	// MakeNew get service by key identify, but is not singleton
	// New instance by service provider register boot function pass parameters
	// This function is useful when you need to start different instances for different parameters
	MakeNew(key string, params []interface{}) (interface{}, error)
}

// GoHubContainer framework container
type GoHubContainer struct {
	// Must implementation
	Container

	// Providers store service provider, key is identify
	providers map[string]ServiceProvider

	// instances store specific instances
	instances map[string]interface{}

	// Lock container operator
	lock sync.RWMutex
}

// NewGoHubContainer new gohub instance
func NewGoHubContainer() *GoHubContainer {
	return &GoHubContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

// Bind binding provider on gohub container
func (gohub *GoHubContainer) Bind(provider ServiceProvider) error {
	gohub.lock.Lock()
	defer gohub.lock.Unlock()

	key := provider.Name()
	gohub.providers[key] = provider

	if !provider.IsDefer() {
		if err := provider.Boot(gohub); err != nil {
			return err
		}

		// Instance
		params := provider.Params(gohub)
		method := provider.Register(gohub)
		instance, err := method(params...)
		if err != nil {
			return errors.New(err.Error())
		}
		gohub.instances[key] = instance
	}
	return nil
}

// IsBind check service provider is bind
func (gohub *GoHubContainer) IsBind(key string) bool {
	return gohub.findServiceProvider(key) != nil
}

func (gohub *GoHubContainer) findServiceProvider(key string) ServiceProvider {
	gohub.lock.RLock()
	defer gohub.lock.RUnlock()
	if sp, ok := gohub.providers[key]; ok {
		return sp
	}
	return nil
}

func (gohub *GoHubContainer) Make(key string) (interface{}, error) {
	return gohub.make(key, nil, false)
}

// MustMake when error trigger panic
func (gohub *GoHubContainer) MustMake(key string) interface{} {
	serv, err := gohub.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return serv
}

func (gohub *GoHubContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return gohub.make(key, params, true)
}

func (gohub *GoHubContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	gohub.lock.RLock()
	defer gohub.lock.RUnlock()
	// Query is already registered service provider, if not register will be return error
	sp := gohub.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contact [" + key + "] have not register")
	}

	if forceNew {
		return gohub.newInstance(sp, params)
	}

	// If don't need force new instance on container already instanced
	// will straight retrieve instance
	if ins, ok := gohub.instances[key]; ok {
		return ins, nil
	}

	// If container not instance will be once instance
	inst, err := gohub.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

// newInstance
func (gohub *GoHubContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	if err := sp.Boot(gohub); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(gohub)
	}
	method := sp.Register(gohub)
	ins, err := method(params...)
	return ins, err
}
