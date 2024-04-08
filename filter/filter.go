package filter

import "net/http"

type Filter interface {
	PreHandle(writer http.ResponseWriter, request *http.Request) error

	PostHandle(writer http.ResponseWriter, request *http.Request) error

	GetPriority() int

	GetPathPattern() string
}
