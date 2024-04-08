package log

import (
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/log"
	"github.com/lyzzz123/illusionmvc/response"
	"testing"
)

func LogTest() (*response.JSONResponse, error) {
	log.Info("hello %v", "lyzzz")
	return &response.JSONResponse{
		Data: "test log",
	}, nil
}

func TestPathValue(t *testing.T) {
	logInstance := &log.DefaultLog{}
	logInstance.Init()
	illusionmvc.RegisterLog(logInstance)
	illusionmvc.RegisterHandler("/log_test", []string{httpmethod.POST, httpmethod.GET}, LogTest)
	illusionmvc.StartService("9527")

}
