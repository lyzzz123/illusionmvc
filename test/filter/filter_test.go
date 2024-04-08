package filter

import (
	"fmt"
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/filter"
	"github.com/lyzzz123/illusionmvc/response"
	"net/http"
	"testing"
)

func FilterTest() (*response.JSONResponse, error) {
	jsonResponse := &response.JSONResponse{
		Data: "filter test",
	}

	return jsonResponse, nil
}

type FirstFilter struct {
}

func (firstFilter *FirstFilter) PreHandle(writer http.ResponseWriter, request *http.Request) error {
	fmt.Println("FirstFilter PreHandle:" + request.URL.Path)
	return nil
}

func (firstFilter *FirstFilter) PostHandle(writer http.ResponseWriter, request *http.Request) error {
	fmt.Println("FirstFilter PostHandle")

	return nil
}

func (firstFilter *FirstFilter) GetPriority() int {
	return 3
}

func (firstFilter *FirstFilter) GetPathPattern() string {
	return "/**"
}

type SecondFilter struct {
}

func (secondFilter *SecondFilter) PreHandle(writer http.ResponseWriter, request *http.Request) error {
	fmt.Println("SecondFilter PreHandle:" + request.URL.Path)
	return nil
}

func (secondFilter *SecondFilter) PostHandle(writer http.ResponseWriter, request *http.Request) error {
	fmt.Println("SecondFilter PostHandle")
	return nil
}

func (secondFilter *SecondFilter) GetPriority() int {
	return 4
}

func (secondFilter *SecondFilter) GetPathPattern() string {
	return "/test_filter/*/test_b"
}

func TestFilter(t *testing.T) {
	illusionmvc.RegisterFilter(&FirstFilter{})
	illusionmvc.RegisterFilter(&SecondFilter{})
	illusionmvc.RegisterFilter(&filter.CorsFilter{})

	illusionmvc.RegisterHandler("/test_filter/test_a/test_b", []string{httpmethod.GET}, FilterTest)
	illusionmvc.StartService("9527")
}
