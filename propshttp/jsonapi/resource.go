package jsonapi

import (
	"reflect"
	"github.com/lann/builder"
)

const (
	idTagName = "resourceid"
	typeTagName = "resourcetype"
)

type Resource struct {
	ResourceType string      `json:"type,omitempty"`
	ID           string      `json:"id,omitempty"`
	Data         interface{} `json:"attributes,omitempty"`
}

type ResResourceBuilder builder.Builder

func (b ResResourceBuilder) ResourceType(resourceType string) ResResourceBuilder {
	return builder.Set(b, "ResourceType", resourceType).(ResResourceBuilder)
}

func (b ResResourceBuilder) ID(id string) ResResourceBuilder {
	return builder.Set(b, "ID", id).(ResResourceBuilder)
}

func (b ResResourceBuilder) Data(data interface{}) ResResourceBuilder {
	id, resourceType := identifierOf(data)
	return b.ID(id).ResourceType(resourceType).setData(data)
}

func (b ResResourceBuilder) setData(data interface{}) ResResourceBuilder {
	return builder.Set(b, "Data", data).(ResResourceBuilder)
}

func (b ResResourceBuilder) Build() Resource{
	return builder.GetStruct(b).(Resource)
}

func ResourceBuilder() ResResourceBuilder {
	return builder.Register(ResResourceBuilder{}, Resource{}).(ResResourceBuilder)
}

func identifierOf(data interface{}) (string, string) {
	var id, resourceType string
	var idFound, resourceTypeFound, allFound bool

	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		id, idFound = idFieldLookup(t.Field(i).Tag)
		resourceType, resourceTypeFound = resourceTypeFieldLookup(t.Field(i).Tag)

		if idFound && resourceTypeFound {
			break
		}
	}

	if allFound {
		return id, resourceType
	}

	return createDefaultIdentifier(t)
}

func idFieldLookup(tag reflect.StructTag) (string, bool) {
	return tag.Lookup(idTagName)
}

func resourceTypeFieldLookup(tag reflect.StructTag) (string, bool) {
	return tag.Lookup(typeTagName)
}

