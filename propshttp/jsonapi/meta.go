package jsonapi

import "github.com/lann/builder"

type Meta struct {
	//The value of each meta member MUST be an object (a “meta object”).
	Errors []ResError `json:"errors, omitempty"`
}

type ResMetaBuilder builder.Builder

func (mb ResMetaBuilder) Errors(errors []ResError) ResMetaBuilder {
	return builder.Set(mb, "Errors", errors).(ResMetaBuilder)
}

func (mb ResMetaBuilder) Build() Meta {
	return builder.GetStruct(mb).(Meta)
}

func MetaBuilder() ResMetaBuilder {
	return builder.Register(ResMetaBuilder{}, Meta{}).(ResMetaBuilder)
}