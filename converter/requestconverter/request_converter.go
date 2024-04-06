package requestconverter

import (
	"github.com/lyzzz123/illusionmvc/wrapper"
	"net/http"
)

type RequestConverter interface {
	Convert(writer http.ResponseWriter, request *http.Request, param interface{}, inputWrapper *wrapper.InputWrapper) error

	Support(request *http.Request) bool

	Name() string
}
