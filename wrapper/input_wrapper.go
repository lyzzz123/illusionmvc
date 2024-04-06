package wrapper

import "reflect"

type InputWrapper struct {
	PathValuePositionMap  map[int]int
	ParamValuePositionMap map[string]int
	InputType             reflect.Type
}
