package test

import (
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/response"
	"testing"
)

type JSONResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JsonResponseTest() (*response.JSONResponse, error) {
	r := &JSONResponse{Code: 1, Message: "success", Data: "business data, maybe a struct or a primary type data"}
	return &response.JSONResponse{
		Data: r,
	}, nil
}

func TestJsonResponse(t *testing.T) {

	illusionmvc.RegisterHandler("/json_response_test", []string{httpmethod.POST}, JsonResponseTest)
	illusionmvc.StartService("9527")

}
