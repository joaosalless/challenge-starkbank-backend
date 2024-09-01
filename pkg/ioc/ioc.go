package ioc

import (
	"go.uber.org/dig"
	"log"
)

type Dependency struct {
	Constructor interface{}
	Interface   interface{}
	Name        string
}

type In struct {
	dig.In
}

var container *dig.Container

func New(deps []Dependency) *dig.Container {
	container = dig.New()

	for _, dep := range deps {
		var err error

		if dep.Interface != nil {
			err = container.Provide(dep.Constructor, dig.As(dep.Interface), dig.Name(dep.Name))
		} else {
			err = container.Provide(dep.Constructor, dig.Name(dep.Name))
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	return container
}
