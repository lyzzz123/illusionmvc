package response_writer

import (
	"io"
	"net/http"
	"reflect"
)

type StringResponse struct {
	Data string
}

type StringResponseWriter struct {
	ResponseType reflect.Type
}

func (stringResponseWriter *StringResponseWriter) Write(writer http.ResponseWriter, returnValue interface{}) error {

	stringReturnValue := returnValue.(*StringResponse)
	if stringReturnValue != nil {
		io.WriteString(writer, stringReturnValue.Data)
	}
	return nil
}

func (stringResponseWriter *StringResponseWriter) Support() reflect.Type {
	return stringResponseWriter.ResponseType
}
