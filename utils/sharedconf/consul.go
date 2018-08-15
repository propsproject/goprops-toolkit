package sharedconf

import (
	"fmt"
	"github.com/google/uuid"

	consul "github.com/hashicorp/consul/api"
)

type ConsulClient struct {
	consul   *consul.Client
	services []string
}

type ConsulRegistration struct {
	Name string
	Port int
}

func newConsulClient(endpoint string) (*ConsulClient, error) {
	config := consul.DefaultConfig()
	config.Address = endpoint
	c, err:= consul.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &ConsulClient{
		consul: c,
		services: make([]string, 0),
	}, nil
}

// Register a service with consul local agent
func (c *ConsulClient) Register(microserviceName string, reg ConsulRegistration, isDevelopment bool) (string, error) {
	id := uuid.New().String()

	env := "production"
	if isDevelopment {
		env = "development"
	}

	consulReg := &consul.AgentServiceRegistration{
		ID:   fmt.Sprintf("%s.%s", reg.Name, id),
		Name: fmt.Sprintf("%s.%s", microserviceName, id),
		Port: reg.Port,
		Tags: []string{env},
		EnableTagOverride: false,
	}
	err := c.consul.Agent().ServiceRegister(consulReg)
	c.services = append(c.services, consulReg.ID)
	return consulReg.ID, err
}

func (c *ConsulClient) Clean() {
	for _, serviceId := range c.services {
		c.DeRegister(serviceId)
	}
}

// DeRegister a service with consul local agent
func (c *ConsulClient) DeRegister(id string) error {
	return c.consul.Agent().ServiceDeregister(id)
}

// Service return a service
func (c *ConsulClient) Service(service, tag string) ([]*consul.ServiceEntry, *consul.QueryMeta, error) {
	addrs, meta, err := c.consul.Health().Service(service, tag, true, nil)
	if len(addrs) == 0 && err == nil {
		return nil, nil, fmt.Errorf("service ( %s ) was not found", service)
	}
	if err != nil {
		return nil, nil, err
	}
	return addrs, meta, nil
}
