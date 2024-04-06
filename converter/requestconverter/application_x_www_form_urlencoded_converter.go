package requestconverter

import (
	"github.com/lyzzz123/illusionmvc/wrapper"
	"net/http"
)

type ApplicationXWWWFormUrlencodedConverter struct {
}

func (applicationXWWWFormUrlencodedConverter *ApplicationXWWWFormUrlencodedConverter) Convert(writer http.ResponseWriter, request *http.Request, param interface{}, inputWrapper *wrapper.InputWrapper) error {
	request.ParseForm()

	return nil
}

func (applicationXWWWFormUrlencodedConverter *ApplicationXWWWFormUrlencodedConverter) Support(request *http.Request) bool {
	if request.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		return true
	}

	return false
}

func (applicationXWWWFormUrlencodedConverter *ApplicationXWWWFormUrlencodedConverter) Name() string {
	return "ApplicationXWWWFormUrlencodedConverter"
}
