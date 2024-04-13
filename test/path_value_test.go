package test

import (
	"fmt"
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/response"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestPathValueClient(t *testing.T) {
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get("http://localhost:9527/path_value_test/1212/2323")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(result))
}

type PathValueParam struct {
	Hello string `paramValue:"hello_a"`
	Name  string `paramValue:"name_b"`
}

func PathValueTest(param *PathValueParam) (*response.JSONResponse, error) {
	message := param.Hello + " " + param.Name
	return &response.JSONResponse{
		Data: message,
	}, nil
}

func TestPathValue(t *testing.T) {

	illusionmvc.RegisterHandler("/path_value_test/{hello_a}/{name_b}", []string{httpmethod.POST, httpmethod.GET}, PathValueTest)
	illusionmvc.StartService("9527")

}
