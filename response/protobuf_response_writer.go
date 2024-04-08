package response

import (
	"github.com/golang/protobuf/proto"
	"net/http"
	"reflect"
)

type ProtobufResponse struct {
	Data proto.Message
}

type ProtobufResponseWriter struct {
	//var protobufResponseType = reflect.TypeOf(new(response2.ProtobufResponse))
	ResponseType reflect.Type
}

func (jsonResponseWriter *ProtobufResponseWriter) Write(writer http.ResponseWriter, returnValue interface{}) error {
	writer.Header().Set("Content-Type", "application/x-protobuf")

	if returnValue == nil {
		return nil
	}

	protobufResponse := returnValue.(*ProtobufResponse)

	if protoBytes, err := proto.Marshal(protobufResponse.Data); err != nil {
		return err
	} else {
		if _, err := writer.Write(protoBytes); err != nil {
			return err
		}
	}
	return nil
}

func (jsonResponseWriter *ProtobufResponseWriter) Support() reflect.Type {
	return jsonResponseWriter.ResponseType
}
