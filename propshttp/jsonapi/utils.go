package jsonapi

import (
	"fmt"
	"reflect"
	"crypto/sha1"
)


func createDefaultIdentifier(t reflect.Type) (string, string) {
	resourceType := getTypeName(t)
	id := sha1Hash(resourceType)
	return resourceType, id
}

func getTypeName(t reflect.Type) string {
	if t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	}
	return t.Name()
}

func sha1Hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}