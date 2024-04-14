package test

import (
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/handler"
	"testing"
)

func TestStaticResource(t *testing.T) {

	illusionmvc.RegisterStaticHandler(&handler.DefaultStaticHandler{
		StaticPath: "/static/",
		StaticDir:  "F:\\temp",
	})
	illusionmvc.StartService("9527")

}
