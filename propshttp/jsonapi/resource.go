package jsonapi

import (
	"reflect"
)

const (
	idTagName = "resourceid"
	typeTagName = "resourcetype"
)

type Resource struct {
	ResourceType string      `json:"type"`
	ID           string      `json:"id"`
	Data         interface{} `json:"attributes"`
}

func (r *Resource) HasIdentifier() bool {
	return r.ID != "" && r.ResourceType != ""
}

type ResourceBuilder struct {
	resource *Resource
}

func (r *ResourceBuilder) Resource() *Resource {
	if ok := r.resource.HasIdentifier(); !ok {
		r.GetIdentifier()
	}

	resource := &Resource{ResourceType: r.resource.ResourceType, ID: r.resource.ID, Data: r.resource.Data}
	r.Reset()

	return resource
}

func (r *ResourceBuilder) GetIdentifier() *ResourceBuilder {
	t := reflect.TypeOf(r.resource.Data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if tag := field.Tag.Get(idTagName); tag != "" {
			v := reflect.ValueOf(r.resource.Data)
			r.resource.ID = reflect.Indirect(v).FieldByName(field.Name).String()
		}

		if tag := field.Tag.Get(typeTagName); tag != "" {
			v := reflect.ValueOf(r.resource.Data)
			r.resource.ResourceType = reflect.Indirect(v).FieldByName(field.Name).String()
		}
	}

	if r.resource.ResourceType == "" || r.resource.ID == "" {
		r.resource.ResourceType = t.Name()
		r.resource.ID = ""
	}

	return r
}

func (r *ResourceBuilder) WithDefinition (resourceType, id string) *ResourceBuilder {
	r.resource.ResourceType = resourceType
	r.resource.ID = id
	return r
}

func (r *ResourceBuilder) SetData(data interface{}) *ResourceBuilder {
	r.resource.Data = data
	return r
}

func (r *ResourceBuilder) Reset() *ResourceBuilder {
	r.resource = &Resource{}
	return r
}

