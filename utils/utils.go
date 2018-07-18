package utils

import (
	"net/http"
	"context"
)

func SetRequestContext(r *http.Request, key string, value interface{}) *http.Request {
	ctx := context.WithValue(r.Context(), key, value)
	r.WithContext(ctx)
	return r
}

func RequestWithContext(r *http.Request, pairs map[string]interface{}) *http.Request  {
	for k, v := range pairs {
		r = SetRequestContext(r, k, v)
	}

	return r
}