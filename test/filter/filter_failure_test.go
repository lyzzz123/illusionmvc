package filter

import (
	"errors"
	"fmt"
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/response"
	"net/http"
	"testing"
)

func FilterFailureTest() (*response.JSONResponse, error) {
	jsonResponse := &response.JSONResponse{
		Data: "filter test",
	}

	return jsonResponse, nil
}

type FirstFilterFailure struct {
}

func (firstFilterFailure *FirstFilterFailure) PreHandle(writer http.ResponseWriter, request *http.Request) error {
	fmt.Println("FirstFilter PreHandle")
	//i := 0
	//fmt.Println(2 / i)
	return errors.New("firstFilterFailure PreHandle")
}

func (firstFilterFailure *FirstFilterFailure) PostHandle(writer http.ResponseWriter, request *http.Request) error {
	fmt.Println("FirstFilter PostHandle")
	//i := 0
	//fmt.Println(2 / i)
	//return errors.New("firstFilterFailure PostHandle")
	return nil
}

func (firstFilterFailure *FirstFilterFailure) GetPriority() int {
	return 3
}

func (firstFilterFailure *FirstFilterFailure) GetPathPattern() string {
	return "/**"
}

func TestFilterFailure(t *testing.T) {
	illusionmvc.RegisterFilter(&FirstFilterFailure{})
	illusionmvc.RegisterHandler("/test_filter_failure", []string{httpmethod.GET}, FilterFailureTest)
	illusionmvc.StartService("9527")
}
