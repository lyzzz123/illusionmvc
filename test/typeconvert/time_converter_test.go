package typeconvert

import (
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/response"
	"testing"
	"time"
)

type TimeConverterParam struct {
	Hello    string     `paramValue:"hello"`
	TestTime *time.Time `paramValue:"testTime"`
}

func TimeConverterTest(param *TimeConverterParam) (*response.JSONResponse, error) {
	message := param.Hello + ":" + param.TestTime.String()
	return &response.JSONResponse{
		Data: message,
	}, nil
}

func TestJsonRequest(t *testing.T) {
	illusionmvc.RegisterTypeConverter(&TimeConvert{})
	illusionmvc.RegisterHandler("/time_converter_test", []string{httpmethod.POST}, TimeConverterTest)
	illusionmvc.StartService("9527")

}
