package responsewriter

import (
	"encoding/json"
	response2 "github.com/lyzzz123/illusionmvc/response"
	"io"
	"net/http"
	"reflect"
)

var jsonResponseType = reflect.TypeOf(new(response2.JSONResponse))

type JSONResponseWriter struct {
}

func (jsonResponseWriter *JSONResponseWriter) Write(writer http.ResponseWriter, returnValue interface{}) error {
	writer.Header().Set("Content-Type", "application/json")

	if returnValue == nil {
		return nil
	}

	jsonStringResponse := returnValue.(*response2.JSONResponse)

	if bytes, err := json.Marshal(jsonStringResponse.Data); err != nil {
		return err
	} else {
		io.WriteString(writer, string(bytes))
	}

	return nil
}

func (jsonResponseWriter *JSONResponseWriter) Support() reflect.Type {
	return jsonResponseType
}
