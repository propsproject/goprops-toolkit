package jsonapi

import (
	"github.com/lann/builder"
)

type ResError struct {
	Code   string      `json:"code"`
	Detail string      `json:"detail"`
	Source interface{} `json:"source"`
	Status int         `json:"status"`
}

type ErrorBuilder builder.Builder

func (eb ErrorBuilder) Code(code string) ErrorBuilder {
	return builder.Set(eb, "Code", code).(ErrorBuilder)
}

func (eb ErrorBuilder) Detail(detail string) ErrorBuilder {
	return builder.Set(eb, "Detail", detail).(ErrorBuilder)
}

func (eb ErrorBuilder) Source(source interface{}) ErrorBuilder {
	return builder.Set(eb, "Source", source).(ErrorBuilder)
}

func (eb ErrorBuilder) Status(status int) ErrorBuilder {
	return builder.Set(eb, "Status", status).(ErrorBuilder)
}

func (eb ErrorBuilder) Build() ResError {
	return builder.GetStruct(eb).(ResError)
}

func (eb ErrorBuilder) BuildSlice() []ResError {
	return []ResError{builder.GetStruct(eb).(ResError)}
}

func Error() ErrorBuilder {
	return builder.Register(ErrorBuilder{}, ResError{}).(ErrorBuilder)
}
