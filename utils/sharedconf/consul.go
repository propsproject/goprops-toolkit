package sharedconf

import (
	"fmt"
	"github.com/satori/go.uuid"

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

func newConsulClient(endpoint string) *ConsulClient {
	config := consul.DefaultConfig()
	config.Address = endpoint
	c, _ := consul.NewClient(config)
	return &ConsulClient{
		consul: c,
		services: make([]string, 0),
	}
}

// Register a service with consul local agent
func (c *ConsulClient) Register(name string, port int) (string, error) {
	id := uuid.NewV4().String()

	c.services = append(c.services, id)
	reg := &consul.AgentServiceRegistration{
		ID:   id,
		Name: name,
		Port: port,
	}
	err := c.consul.Agent().ServiceRegister(reg)
	return id, err
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
