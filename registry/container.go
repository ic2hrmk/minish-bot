package registry

import (
	"fmt"
)

type FactoryMethod func() (Application, error)

type Application interface {
	Run() error
}

func NewRegistryContainer() registryContainer {
	return make(registryContainer)
}

type registryContainer map[string]FactoryMethod

func (c *registryContainer) Add(serviceName string, fabric FactoryMethod) {
	(*c)[serviceName] = fabric
}

func (c *registryContainer) Get(serviceName string) (FactoryMethod, error) {
	entry, ok := (*c)[serviceName]
	if !ok {
		return nil, fmt.Errorf("service [%s] isn't registered", serviceName)
	}

	return entry, nil
}
