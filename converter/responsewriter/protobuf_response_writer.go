package responsewriter

import (
	"github.com/golang/protobuf/proto"
	response2 "github.com/lyzzz123/illusionmvc/response"
	"net/http"
	"reflect"
)

var protobufResponseType = reflect.TypeOf(new(response2.ProtobufResponse))

type ProtobufResponseWriter struct {
}

func (jsonResponseWriter *ProtobufResponseWriter) Write(writer http.ResponseWriter, returnValue interface{}) error {
	writer.Header().Set("Content-Type", "application/x-protobuf")

	if returnValue == nil {
		return nil
	}

	protobufResponse := returnValue.(*response2.ProtobufResponse)

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
	return protobufResponseType
}
