package handler

import (
	"errors"
	"github.com/lyzzz123/illusionmvc/filter"
	"github.com/lyzzz123/illusionmvc/handler/exceptionhandler"
	"github.com/lyzzz123/illusionmvc/log"
	"github.com/lyzzz123/illusionmvc/request/requestconverter"
	"github.com/lyzzz123/illusionmvc/response"
	"github.com/lyzzz123/illusionmvc/utils"
	"github.com/lyzzz123/illusionmvc/wrapper"
	"net/http"
	"reflect"
	"strings"
)

type Wrapper struct {
	Path         string
	HasPathValue bool
	ReplacedPath string
	//PathValueRegex *regexp.Regexp
	Handler     interface{}
	FilterArray []filter.Filter
	HttpMethod  []string
	OutputType  reflect.Type
	Input       *wrapper.InputWrapper

	DefaultExceptionHandler exceptionhandler.ExceptionHandler
	RequestConverterMap     map[string]requestconverter.RequestConverter
	ResponseWriter          response.Writer
}

func (wrapper *Wrapper) Handle(writer http.ResponseWriter, request *http.Request) error {
	args := make([]reflect.Value, 0)
	if wrapper.Input.InputType != nil {
		inputParam := reflect.New(wrapper.Input.InputType).Interface()
		if err := wrapper.setInputParam(writer, request, inputParam); err != nil {
			return err
		}
		args = append(args, reflect.ValueOf(inputParam))
	}
	result := reflect.ValueOf(wrapper.Handler).Call(args)
	if result[1].Interface() != nil {
		if err := wrapper.DefaultExceptionHandler.Handle(writer, result[1].Interface().(error)); err != nil {
			return err
		}
		return nil
	}
	returnValue := result[0].Interface()
	if err := wrapper.ResponseWriter.Write(writer, returnValue); err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (wrapper *Wrapper) setInputParam(writer http.ResponseWriter, request *http.Request, inputParam interface{}) error {

	var findConverter requestconverter.RequestConverter = nil
	for _, requestConverter := range wrapper.RequestConverterMap {
		if requestConverter.Support(request) {
			findConverter = requestConverter
		}
	}
	if findConverter == nil {
		panic("not support Content-Type:" + request.Header.Get("Content-Type"))
	}
	if err := findConverter.Convert(writer, request, inputParam, wrapper.Input); err != nil {
		return err
	}

	if wrapper.HasPathValue {
		if err := wrapper.setPathValue(request, inputParam); err != nil {
			return err
		}
	}

	if err := wrapper.setPathValue(request, inputParam); err != nil {
		return err
	}
	if err := wrapper.setHeaderValue(request, inputParam); err != nil {
		return err
	}
	if err := wrapper.setCookieValue(request, inputParam); err != nil {
		return err
	}
	if err := wrapper.setRequestAndResponse(writer, request, inputParam); err != nil {
		return err
	}
	return nil
}

func (wrapper *Wrapper) setRequestAndResponse(writer http.ResponseWriter, request *http.Request, param interface{}) error {
	reflectParamValue := reflect.ValueOf(param).Elem()
	if wrapper.Input.HasRequestParam {
		reflectParamValue.Field(wrapper.Input.ResponseParamIndex).Set(reflect.ValueOf(request))
	}
	if wrapper.Input.HasResponseParam {
		reflectParamValue.Field(wrapper.Input.ResponseParamIndex).Set(reflect.ValueOf(writer))
	}
	return nil
}

func (wrapper *Wrapper) setCookieValue(request *http.Request, param interface{}) error {
	reflectParamValue := reflect.ValueOf(param).Elem()
	cookies := request.Cookies()
	for _, cookie := range cookies {
		paramIndex, ok := wrapper.Input.ParamValuePositionMap[cookie.Name]
		if ok {
			field := reflectParamValue.Field(paramIndex)
			if converter, ok := wrapper.Input.TypeConverterMap[field.Type()]; ok {
				if parsedValue, err := converter.Convert(cookie.Value); err != nil {
					return err
				} else {
					field.Set(reflect.ValueOf(parsedValue))
				}
			}
		}
	}
	return nil
}

func (wrapper *Wrapper) setHeaderValue(request *http.Request, param interface{}) error {
	reflectParamValue := reflect.ValueOf(param).Elem()
	for headerName, headerValue := range request.Header {
		paramIndex, ok := wrapper.Input.ParamValuePositionMap[headerName]
		if ok {
			field := reflectParamValue.Field(paramIndex)
			if converter, ok := wrapper.Input.TypeConverterMap[field.Type()]; ok {
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

func (wrapper *Wrapper) setPathValue(request *http.Request, param interface{}) error {
	pathValueArray := wrapper.getPathValueArray(request.URL.Path)
	reflectParamValue := reflect.ValueOf(param).Elem()
	for i := 0; i < len(pathValueArray); i++ {
		if index, ok := wrapper.Input.PathValuePositionMap[i]; ok {
			field := reflectParamValue.Field(index)
			if converter, ok := wrapper.Input.TypeConverterMap[field.Type()]; ok {
				if parsedValue, err := converter.Convert(pathValueArray[i]); err != nil {
					return err
				} else {
					field.Set(reflect.ValueOf(parsedValue))
				}
			}
		}
	}
	return nil
}

func (wrapper *Wrapper) getPathValueArray(path string) []string {
	subPathArray := strings.Split(path, "/")
	replacedSubPathArray := strings.Split(wrapper.ReplacedPath, "/")
	pathValueArray := make([]string, 0)
	for i := 0; i < len(subPathArray); i++ {
		if replacedSubPathArray[i] == "*" {
			pathValueArray = append(pathValueArray, subPathArray[i])
		}
	}
	return pathValueArray
}

func CreateHandlerWrapper(path string, httpMethod []string, handlerMethod interface{}) *Wrapper {
	if path == "" {
		panic("RegisterHandler path must not be blank")
	}

	if httpMethod == nil {
		panic("RegisterHandler httpMethod must not be nil")
	}

	if len(httpMethod) == 0 {
		panic("RegisterHandler httpMethod length not be zero")
	}

	if handlerMethod == nil {
		panic("RegisterHandler handlerMethod not be nil")
	}

	hw := &Wrapper{Input: &wrapper.InputWrapper{}}
	hw.Path = path
	hw.ReplacedPath = utils.CreateReplacedPath(path)
	hw.Handler = handlerMethod
	if pt, err := getInputType(handlerMethod); err != nil {
		panic(err)
	} else {
		hw.Input.InputType = pt
	}
	if ot, err := getOutputType(handlerMethod); err != nil {
		panic(err)
	} else {
		hw.OutputType = ot
	}

	inputValueNameIndexMap := createInputValueNameIndexMap(hw.Input.InputType)
	hw.Input.ParamValuePositionMap = inputValueNameIndexMap
	if utils.HasPathValue(path) {
		pathValueNameIndexMap := createPathValueNameIndexMap(path)
		hw.Input.PathValuePositionMap = createPathValuePositionMap(pathValueNameIndexMap, inputValueNameIndexMap)
		hw.HasPathValue = true
	}
	hw.HttpMethod = httpMethod
	checkRequestAndResponse(hw.Input)
	return hw
}

func createPathValuePositionMap(pathValueNameIndexMap map[string]int, inputValueNameIndexMap map[string]int) map[int]int {
	pathValuePositionMap := make(map[int]int, len(pathValueNameIndexMap))
	for pathValueName, pathValueIndex := range pathValueNameIndexMap {
		inputValueNameIndex, ok := inputValueNameIndexMap[pathValueName]
		if ok {
			pathValuePositionMap[pathValueIndex] = inputValueNameIndex
		} else {
			panic("not find " + pathValueName)
		}
	}
	return pathValuePositionMap
}

func createInputValueNameIndexMap(inputType reflect.Type) map[string]int {

	paramMap := make(map[string]int)
	if inputType == nil {
		return paramMap
	}
	for i := 0; i < inputType.NumField(); i++ {
		paramValue := inputType.Field(i).Tag.Get("paramValue")
		if paramValue != "" {
			paramMap[paramValue] = i
		}
	}
	return paramMap
}

func checkRequestAndResponse(inputWrapper *wrapper.InputWrapper) {
	if inputWrapper.InputType == nil {
		return
	}
	for i := 0; i < inputWrapper.InputType.NumField(); i++ {
		fieldValue := inputWrapper.InputType.Field(i)
		if fieldValue.Type.String() == "http.ResponseWriter" {
			inputWrapper.HasRequestParam = true
			inputWrapper.RequestParamIndex = i
		}
		if fieldValue.Type.String() == "*http.Request" {
			inputWrapper.HasResponseParam = true
			inputWrapper.ResponseParamIndex = i
		}
	}
}

func createPathValueNameIndexMap(path string) map[string]int {
	pathValueArray := strings.Split(path, "/")
	subPathNameArray := make([]string, 0)

	for i := 0; i < len(pathValueArray); i++ {
		if strings.HasPrefix(pathValueArray[i], "{") &&
			strings.HasSuffix(pathValueArray[i], "}") {
			replacedSubPathName := strings.TrimRight(strings.TrimLeft(pathValueArray[i], "{"), "}")
			subPathNameArray = append(subPathNameArray, replacedSubPathName)
		}
	}
	r := make(map[string]int, len(subPathNameArray))
	for i := 0; i < len(subPathNameArray); i++ {
		r[subPathNameArray[i]] = i
	}

	//
	//pathValueRegex := regexp.MustCompile("{[^/]+}")
	//replacedPathValuePath := pathValueRegex.ReplaceAllString(path, "{([^/]+)}")
	//replacedPathValueRegex := regexp.MustCompile("^" + replacedPathValuePath + "$")
	//pathValueNameArray := replacedPathValueRegex.FindStringSubmatch(path)[1:]
	//r := make(map[string]int, len(pathValueNameArray)-1)
	//for i := 0; i < len(pathValueNameArray); i++ {
	//	r[pathValueNameArray[i]] = i
	//}
	return r
}

func getInputType(handler interface{}) (reflect.Type, error) {
	methodType := reflect.ValueOf(handler).Type()
	if methodType.NumIn() == 0 {
		return nil, nil
	}
	if methodType.NumIn() != 1 {
		return nil, errors.New(methodType.String() + " must has one input param")
	}
	argType := methodType.In(0)
	if argType.Kind() != reflect.Ptr {
		return nil, errors.New(methodType.String() + " must has one ptr input param")
	} else {
		if argType.Elem().Kind() != reflect.Struct {
			return nil, errors.New(methodType.String() + " must has one struct ptr input param")
		}
	}
	return argType.Elem(), nil
}

func getOutputType(handler interface{}) (reflect.Type, error) {
	methodType := reflect.ValueOf(handler).Type()
	if methodType.NumOut() != 2 {
		return nil, errors.New(methodType.String() + " must has two output param")
	}
	firstOutputType := methodType.Out(0)
	if firstOutputType.Kind() != reflect.Ptr {
		return nil, errors.New(methodType.String() + " first output param must be ptr")
	} else {
		if firstOutputType.Elem().Kind() != reflect.Struct {
			return nil, errors.New(methodType.String() + " first output param must be struct ptr")
		}
	}

	secondOutputType := methodType.Out(1)

	if secondOutputType.String() != "error" {
		return nil, errors.New(methodType.String() + " second output param must be error")
	}

	return firstOutputType.Elem(), nil
}
