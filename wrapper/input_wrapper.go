package wrapper

import (
	"github.com/lyzzz123/illusionmvc/converter/typeconverter"
	"reflect"
)

type InputWrapper struct {
	PathValuePositionMap  map[int]int
	ParamValuePositionMap map[string]int
	InputType             reflect.Type
	TypeConverterMap      map[reflect.Type]typeconverter.Converter

	HasRequestParam    bool
	RequestParamIndex  int
	HasResponseParam   bool
	ResponseParamIndex int
}
