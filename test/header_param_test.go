package test

import (
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/response"
	"testing"
)

type HeaderParam struct {
	UserAgent string `paramValue:"User-Agent"`
}

func HeaderParamTest(param *HeaderParam) (*response.JSONResponse, error) {
	message := param.UserAgent
	return &response.JSONResponse{
		Data: message,
	}, nil
}

func TestHeaderParam(t *testing.T) {

	illusionmvc.RegisterHandler("/header_param_test", []string{httpmethod.POST, httpmethod.GET}, HeaderParamTest)
	illusionmvc.StartService("9527")

}
