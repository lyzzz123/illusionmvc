package response_writer

import (
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"reflect"
	"testing"
)

func StringResponseTest() (*StringResponse, error) {

	return &StringResponse{
		Data: "string response",
	}, nil
}

func TestJsonRequest(t *testing.T) {
	illusionmvc.RegisterResponseWriter(&StringResponseWriter{ResponseType: reflect.TypeOf(*new(StringResponse))})
	illusionmvc.RegisterHandler("/string_response_test", []string{httpmethod.GET}, StringResponseTest)
	illusionmvc.StartService("9527")
}
