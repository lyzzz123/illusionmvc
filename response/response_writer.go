package response

import (
	"net/http"
	"reflect"
)

type Writer interface {
	Write(writer http.ResponseWriter, returnValue interface{}) error

	Support() reflect.Type
}
