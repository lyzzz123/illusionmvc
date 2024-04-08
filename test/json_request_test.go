package test

import (
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/response"
	"testing"
)

type JSONParam struct {
	Hello string `json:"hello"`
	Name  string `json:"name"`
}

//Content-Type must be application/json
func JsonRequestTest(param *JSONParam) (*response.JSONResponse, error) {
	message := param.Hello + ":" + param.Name
	return &response.JSONResponse{
		Data: message,
	}, nil
}

func TestJsonRequest(t *testing.T) {

	illusionmvc.RegisterHandler("/json_request_test", []string{httpmethod.POST}, JsonRequestTest)
	illusionmvc.StartService("9527")

}
