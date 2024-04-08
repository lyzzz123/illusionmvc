package test

import (
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/response"
	"testing"
)

type CookieParam struct {
	Cookie1 string `paramValue:"Cookie_1"`
}

func CookieParamTest(param *CookieParam) (*response.JSONResponse, error) {
	message := param.Cookie1
	return &response.JSONResponse{
		Data: message,
	}, nil
}

func TestCookieParam(t *testing.T) {

	illusionmvc.RegisterHandler("/cookie_param_test", []string{httpmethod.POST, httpmethod.GET}, CookieParamTest)
	illusionmvc.StartService("9527")

}
