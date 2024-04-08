package test

import (
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/response"
	"testing"
)

type PathValueParam struct {
	Hello string `paramValue:"hello"`
	Name  string `paramValue:"name"`
}

func PathValueTest(param *PathValueParam) (*response.JSONResponse, error) {
	message := param.Hello + " " + param.Name
	return &response.JSONResponse{
		Data: message,
	}, nil
}

func TestPathValue(t *testing.T) {

	illusionmvc.RegisterHandler("/path_value_test/{hello}/{name}", []string{httpmethod.POST, httpmethod.GET}, PathValueTest)
	illusionmvc.StartService("9527")

}