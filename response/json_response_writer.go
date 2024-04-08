package response

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
)

type JSONResponse struct {
	Data interface{}
}

type JSONResponseWriter struct {
	//var jsonResponseType = reflect.TypeOf(new(response2.JSONResponse))
	ResponseType reflect.Type
}

func (jsonResponseWriter *JSONResponseWriter) Write(writer http.ResponseWriter, returnValue interface{}) error {
	writer.Header().Set("Content-Type", "application/json")

	if returnValue == nil {
		return nil
	}

	jsonStringResponse := returnValue.(*JSONResponse)

	if bytes, err := json.Marshal(jsonStringResponse.Data); err != nil {
		return err
	} else {
		io.WriteString(writer, string(bytes))
	}

	return nil
}

func (jsonResponseWriter *JSONResponseWriter) Support() reflect.Type {
	return jsonResponseWriter.ResponseType
}
