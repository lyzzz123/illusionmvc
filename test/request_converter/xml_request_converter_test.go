package request_converter

import (
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/response"
	"testing"
)

type XMLParam struct {
	Hello string `xml:"hello"`
	Name  string `xml:"name"`
}

//Content-Type must be application/xml
func XmlRequestTest(param *XMLParam) (*response.JSONResponse, error) {
	message := param.Hello + ":" + param.Name
	return &response.JSONResponse{
		Data: message,
	}, nil
}

func TestJsonRequest(t *testing.T) {
	illusionmvc.RegisterRequestConverter(&XmlConverter{})
	illusionmvc.RegisterHandler("/xml_request_test", []string{httpmethod.POST}, XmlRequestTest)
	illusionmvc.StartService("9527")

}
