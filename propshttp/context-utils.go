package propshttp

import (
	"net/http"
	"context"
)

type ContextValue struct {
	key string
	value interface{}
}

type extension struct {
	values []ContextValue
}

func (c *extension) WithValue(value ...ContextValue) *extension {
	c.values = append(c.values, value...)
	return c
}

func (c *extension) setContextValue(parent context.Context, ctxValue ContextValue) context.Context {
	return context.WithValue(parent, ctxValue.key, ctxValue.value)
}

func (c *extension) Extend(r *http.Request) *http.Request {
	parent := r.Context()
	for _, v := range c.values {
		parent = c.setContextValue(parent, v)
	}

	return r.WithContext(parent)
}

func NewContext() *extension {
	return &extension{}
}

func NewContextValue(key string, v interface{}) ContextValue {
	return ContextValue{key, v}
}
