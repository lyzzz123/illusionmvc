package requestconverter

import (
	"github.com/lyzzz123/illusionmvc/wrapper"
	"net/http"
)

type GetMethodConverter struct {
}

func (getMethodConverter *GetMethodConverter) Convert(writer http.ResponseWriter, request *http.Request, param interface{}, inputWrapper *wrapper.InputWrapper) error {
	request.ParseForm()
	return nil
}

func (getMethodConverter *GetMethodConverter) Support(request *http.Request) bool {
	if request.Method == "GET" {
		return true
	}

	return false
}

func (getMethodConverter *GetMethodConverter) Name() string {
	return "GetMethodConverter"
}
