package hello

import (
	"fmt"
	"gohub/framework"
)

// HelloService implement Service interface
type HelloService struct {
	Service
	c framework.Container
}

func newHelloService(params ...interface{}) (interface{}, error) {
	// Here need to expand the parameters
	c := params[0].(framework.Container)
	fmt.Println("New [HelloService]")
	return &HelloService{c: c}, nil
}

func (s HelloService) SayHello() Foo {
	return Foo{
		Name: "Hi, I am a say hello facade component!",
	}
}
