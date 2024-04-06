package typeconverter

import (
	"reflect"
)

type Converter interface {
	Convert(param string) (interface{}, error)
	Support() reflect.Type
}
