package hello

// Interface files and service credentials for storing services

const KEY = "gohub:hello"

type Service interface {
	SayHello() Foo
}

type Foo struct {
	Name string
}
