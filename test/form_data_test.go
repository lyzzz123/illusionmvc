package test

import (
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/response"
	"mime/multipart"
	"testing"
)

type FormDataParam struct {
	Hello    string                `paramValue:"hello"`
	Name     string                `paramValue:"name"`
	TestFile *multipart.FileHeader `paramValue:"testFile"`
}

//Content-Type must be multipart/form-data
func FormDataTest(param *FormDataParam) (*response.JSONResponse, error) {
	message := param.Hello + " " + param.Name + " " + param.TestFile.Filename
	return &response.JSONResponse{
		Data: message,
	}, nil
}

func TestFormData(t *testing.T) {

	illusionmvc.RegisterHandler("/form_data_test", []string{httpmethod.POST}, FormDataTest)
	illusionmvc.StartService("9527")

}
