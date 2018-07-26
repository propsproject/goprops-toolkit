package jsonapi

import (
	"github.com/propsproject/goprops-toolkit/logger"
	"net/http"
	"fmt"
	"encoding/json"
)

const (
	DefaultContentType = "application/json"
	ContentType = "Content-Type"
)

type ResponseBuilder struct {
	Response       *Response
	responseWriter http.ResponseWriter
	MultiResponse  bool
	verbose bool
	logger  *logger.Wrapper
}

func (r *ResponseBuilder) makeResponse() *ResponseBuilder {
	r.Response = &Response{
		Data: &PrimaryData{},
		Errors: Errors{},
		Meta: &Meta{},
	}
	return r
}


func (r *ResponseBuilder) NewResponse(writer http.ResponseWriter) *ResponseBuilder {
	if exists := r.Response != nil; !exists {
		r.makeResponse()
	}
	r.responseWriter = writer
	return r
}

func (r *ResponseBuilder) SetResponseWriter(writer http.ResponseWriter) *ResponseBuilder {
	r.responseWriter = writer
	return r
}

func (r *ResponseBuilder) ResponseStatus(status int) *ResponseBuilder {
	if exists := r.Response != nil; !exists {
		r.makeResponse()
	}
	r.Response.Status = status
	return r
}

func (r *ResponseBuilder) WithData(key string, data interface{}) *ResponseBuilder {
	if exists := r.Response != nil; !exists {
		r.makeResponse()
	}
	var builder ResourceBuilder
	resource := builder.SetData(data).Resource()
	r.Response.Data.SetResource(key ,resource)
	return r
}

func (r *ResponseBuilder) NewError(code, detail string, source interface{}) *ResponseBuilder {
	if exists := r.Response != nil; !exists {
		r.makeResponse()
	}
	err := ResponseError().SetCode(code).SetDetail(detail).SetSource(source)
	r.Response.Meta.Errors = append(r.Response.Meta.Errors, err)
	return r
}

func (r *ResponseBuilder) Error(err *Error) *ResponseBuilder {
	if exists := r.Response != nil; !exists {
		r.makeResponse()
	}
	r.Response.Meta.Errors = append(r.Response.Meta.Errors, err)
	return r
}

func (r *ResponseBuilder) contentType(contentType string) *ResponseBuilder {
	r.responseWriter.Header().Set(ContentType, contentType)
	return r
}

func (r *ResponseBuilder) writeContentType() *ResponseBuilder {
	if r.responseWriter.Header().Get(ContentType) == "" {
		r.responseWriter.Header().Set(ContentType, DefaultContentType)
	}

	return r
}

func (r *ResponseBuilder) hasResponseStatus() (*ResponseBuilder, bool)  {
	return r, r.Response.Status != 0
}

func (r *ResponseBuilder) writeResponse() (*ResponseBuilder, error) {

	if _, ok := r.hasResponseStatus(); !ok {
		return nil, fmt.Errorf("response status not set")
	}

	if err := json.NewEncoder(r.responseWriter).Encode(r.Response); err != nil {
		r.logger.Error(fmt.Errorf("error marshalling http response %v", err))
		http.Error(r.responseWriter, "Internal server error", http.StatusInternalServerError)
		return r, err
	}

	return r.Reset(), nil
}

func (r *ResponseBuilder) Respond() (*ResponseBuilder, error) {
	return r.writeContentType().writeResponse()
}

func (r *ResponseBuilder) Verbose() *ResponseBuilder {
	r.verbose = true
	return r
}

func (r *ResponseBuilder) WithLogger(logger  *logger.Wrapper) *ResponseBuilder {
	r.logger = logger
	return r
}

func (r *ResponseBuilder) logResponse() *ResponseBuilder {
	if r.verbose && r.logger != nil {
		msg := fmt.Sprintf("http response: %v", r.Response)
		r.logger.Info(msg)
	}

	return r
}

func (r *ResponseBuilder) Reset() *ResponseBuilder {
	r.responseWriter = nil
	r.Response = &Response{}
	r.MultiResponse = false
	return r
}

// multi response or not ? [x]
// set response status [x]
// set response primary data and resource identifier [x]
// set Resource [x]
// set errors [x]
// write status header & write content type header [x]
// log request [x]
// write response to response writer [x]