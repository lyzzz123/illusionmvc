package typeconvert

import (
	"reflect"
	"time"
)

type TimeConvert struct {
}

func (timeConvert *TimeConvert) Convert(param string) (interface{}, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, param)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (timeConvert *TimeConvert) Support() reflect.Type {
	return reflect.TypeOf(&time.Time{})
}
