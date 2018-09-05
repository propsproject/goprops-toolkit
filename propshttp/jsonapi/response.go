package jsonapi

import (
	"encoding/json"
	"fmt"
	"github.com/kc1116/builder"
	"net/http"
)

const (
	DEFAULT_CONTENT_TYPE = "application/json"
	ContentType          = "Content-Type"

	errMarshallingResponseJSON = "error marshalling http Response ( %v ): %v"
)

var (
	errContentTypeNotSet    = fmt.Errorf("content type header has not been set")
	errHTTPStatusNotSet     = fmt.Errorf("status code has not been set")
)

type Response struct {
	Data   interface{} `json:"data,omitempty"`
	Errors []ResError  `json:"errors,omitempty"`

	ContentType    string              `json:"-"`
	Status         int                 `json:"-"`
	Headers        [][]string          `json:"-"`
	ContentTypeSet bool `json:"-"`
	StatusSet      bool `json:"-"`
}

type HTTPResponse struct {
	Response Response
}

type ResBuilder builder.Builder

func (b ResBuilder) Data(data interface{}) ResBuilder {
	return builder.Set(b, "Data", data).(ResBuilder)
}

func (b ResBuilder) Error(errors []ResError) ResBuilder {
	return builder.Extend(b, "Errors", errors).(ResBuilder)
}

func (b ResBuilder) Meta(meta Meta) ResBuilder {
	return builder.Set(b, "Meta", meta).(ResBuilder)
}

func (b ResBuilder) Header(key, value string) ResBuilder {
	return builder.Extend(b, "Headers", []string{key, value}).(ResBuilder)
}

func (b ResBuilder) ContentType(contentType string) ResBuilder {
	return builder.Set(b, "ContentType", contentType).(ResBuilder).contentTypeSet()
}

func (b ResBuilder) Status(status int) ResBuilder {
	return builder.Set(b, "Status", status).(ResBuilder).statusSet()
}

func (b ResBuilder) statusSet() ResBuilder {
	return builder.Set(b, "StatusSet", true).(ResBuilder)
}

func (b ResBuilder) contentTypeSet() ResBuilder {
	return builder.Set(b, "ContentTypeSet", true).(ResBuilder)
}

func (b ResBuilder) Build() *HTTPResponse {
	return &HTTPResponse{builder.GetStruct(b).(Response)}
}

func (r *HTTPResponse) validateResponse() error {
	if !r.Response.StatusSet {
		return errHTTPStatusNotSet
	} else if !r.Response.ContentTypeSet {
		return errContentTypeNotSet
	}

	return nil
}

func (r *HTTPResponse) writeHeaders(w http.ResponseWriter) error {
	if err := r.validateResponse(); err != nil {
		return err
	}

	w.Header().Set(ContentType, r.Response.ContentType)
	for _, header := range r.Response.Headers {
		w.Header().Set(header[0], header[1])
	}
	w.WriteHeader(r.Response.Status)

	return nil
}

func (r *HTTPResponse) Respond(w http.ResponseWriter) error {
	if err := r.writeHeaders(w); err != nil {
		return err
	}
	
	if err := json.NewEncoder(w).Encode(&r.Response); err != nil {
		return fmt.Errorf(errMarshallingResponseJSON, r, err)
	}

	return nil
}

func ResponseBuilder() ResBuilder {
	return builder.Register(ResBuilder{}, Response{}).(ResBuilder)
}
