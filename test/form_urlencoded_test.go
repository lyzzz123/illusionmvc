package test

import (
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/response"
	"testing"
)

type FormUrlEncodedParam struct {
	Hello string `paramValue:"hello"`
	Name  string `paramValue:"name"`
}

//Content-Type must be application/x-www-form-urlencoded
func FormUrlEncodedTest(param *FormUrlEncodedParam) (*response.JSONResponse, error) {
	message := param.Hello + " " + param.Name
	return &response.JSONResponse{
		Data: message,
	}, nil
}

func TestFormUrlEncoded(t *testing.T) {

	illusionmvc.RegisterHandler("/form_url_encoded_test", []string{httpmethod.POST, httpmethod.GET}, FormUrlEncodedTest)
	illusionmvc.StartService("9527")

}
