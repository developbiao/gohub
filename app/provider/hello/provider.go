package hello

import (
	"fmt"
	"gohub/framework"
)

// HelloServiceProvider implement ServiceProvider
type HelloServiceProvider struct {
	framework.ServiceProvider
}

func (sp *HelloServiceProvider) Name() string {
	return KEY
}

func (sp *HelloServiceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

func (sp *HelloServiceProvider) Register(c framework.Container) framework.NewInstance {
	return newHelloService
}

func (sp *HelloServiceProvider) IsDefer() bool {
	return true
}

func (sp *HelloServiceProvider) Boot(c framework.Container) error {
	fmt.Println("HelloService boot...")
	return nil
}
