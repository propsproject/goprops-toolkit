package jsonapi

import (
	"net/http"
	"fmt"
	"github.com/propsproject/goprops-toolkit/logger"
)

const (
	errInternalServerError = "internal server error"
)

type ReqResponder struct {
	verbose bool
	logger  *logger.Wrapper
	*AllBuilders
}

func (r *ReqResponder) InternalServerError(responseWriter http.ResponseWriter, message ...string) {
	msg := errInternalServerError
	if len(message) > 0 {
		msg = message[0]
	}

	http.Error(responseWriter, msg, http.StatusInternalServerError)
}

func (r *ReqResponder) Verbose() *ReqResponder {
	r.verbose = true
	return r
}

func (r *ReqResponder) Logger(logger  *logger.Wrapper) *ReqResponder {
	r.logger = logger
	return r
}

func (r *ReqResponder) log(response Response) *ReqResponder {
	if r.verbose && r.logger != nil {
		msg := fmt.Sprintf("http Response: %v", response)
		r.logger.Info(msg)
	}

	return r
}

func Responder() *ReqResponder {
	return &ReqResponder{AllBuilders: Builders()}
}

// multi Response or not ? [x]
// set Response status [x]
// set Response primary data and resource identifier [x]
// set Resource [x]
// set errors [x]
// write status header & write content type header [x]
// log request [x]
// write Response to Response writer [x]