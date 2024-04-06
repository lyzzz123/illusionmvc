package requestconverter

import (
	"github.com/lyzzz123/illusionmvc/wrapper"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"
)

type MultipartFormDataConverter struct {
}

func (multipartFormDataConverter *MultipartFormDataConverter) Convert(writer http.ResponseWriter, request *http.Request, param interface{}, inputWrapper *wrapper.InputWrapper) error {
	if err := request.ParseMultipartForm(1 << 20); err != nil {
		return err
	}
	if err := parseValue(request.MultipartForm.Value, param, inputWrapper); err != nil {
		return err
	}
	if err := parseFile(request.MultipartForm.File, param, inputWrapper); err != nil {
		return err
	}
	return nil
}

func (multipartFormDataConverter *MultipartFormDataConverter) Support(request *http.Request) bool {

	if strings.HasPrefix(request.Header.Get("Content-Type"), "multipart/form-data") {
		return true
	}

	return false
}

func (multipartFormDataConverter *MultipartFormDataConverter) Name() string {
	return "MultipartFormDataConverter"
}

func parseValue(valueMap map[string][]string, param interface{}, inputWrapper *wrapper.InputWrapper) error {
	reflectParamValue := reflect.ValueOf(param).Elem()
	for name, values := range valueMap {
		paramIndex, ok := inputWrapper.ParamValuePositionMap[name]
		if ok {
			field := reflectParamValue.Field(paramIndex)
			if converter, ok := inputWrapper.TypeConverterMap[field.Type()]; ok {
				if parsedValue, err := converter.Convert(values[0]); err != nil {
					return err
				} else {
					field.Set(reflect.ValueOf(parsedValue))
				}
			}
		}
	}
	return nil
}

func parseFile(fileMap map[string][]*multipart.FileHeader, param interface{}, inputWrapper *wrapper.InputWrapper) error {
	reflectParamValue := reflect.ValueOf(param).Elem()
	for name, files := range fileMap {
		paramIndex, ok := inputWrapper.ParamValuePositionMap[name]
		if ok {
			field := reflectParamValue.Field(paramIndex)
			field.Set(reflect.ValueOf(files[0]))
		}
	}
	return nil
}
