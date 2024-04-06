package requestconverter

import (
	"github.com/lyzzz123/illusionmvc/wrapper"
	"net/http"
	"reflect"
)

type ApplicationXWWWFormUrlencodedConverter struct {
}

func (applicationXWWWFormUrlencodedConverter *ApplicationXWWWFormUrlencodedConverter) Convert(writer http.ResponseWriter, request *http.Request, param interface{}, inputWrapper *wrapper.InputWrapper) error {
	request.ParseForm()
	reflectParamValue := reflect.ValueOf(param).Elem()
	for headerName, headerValue := range request.Form {
		paramIndex, ok := inputWrapper.ParamValuePositionMap[headerName]
		if ok {
			field := reflectParamValue.Field(paramIndex)
			if converter, ok := inputWrapper.TypeConverterMap[field.Type()]; ok {
				if parsedValue, err := converter.Convert(headerValue[0]); err != nil {
					return err
				} else {
					field.Set(reflect.ValueOf(parsedValue))
				}
			}
		}
	}

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
