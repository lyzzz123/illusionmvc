package test

import (
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/handler"
	"github.com/lyzzz123/illusionmvc/response"
	"testing"
)

type JSONParam struct {
	Hello       string `json:"hello"`
	Name        string `json:"name"`
	ContentType string `paramValue:"Content-Type"`
}

func Hello(param *JSONParam) (*response.JSONResponse, error) {
	message := param.Hello + ":" + param.Name
	return &response.JSONResponse{
		Data: &response.Response{Code: 0, Message: "success", Data: message},
	}, nil
}

func TestJsonRequest(t *testing.T) {

	illusionmvc.RegisterStaticHandler(&handler.DefaultStaticHandler{
		StaticPaths: "/static",
		StaticDir:   "D:\\temp",
	})
	illusionmvc.RegisterHandler("/testJson", []string{httpmethod.POST}, Hello)
	illusionmvc.StartService("8082")

}
