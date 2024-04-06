package test

import (
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/response"
	"testing"
)

type JSONParam struct {
	Hello string `json:"hello"`
	Name  string `json:"name"`
}

func Hello(param *JSONParam) (*response.JSONResponse, error) {
	message := param.Hello + ":" + param.Name
	return &response.JSONResponse{
		Data: &response.Response{Code: 0, Message: "success", Data: message},
	}, nil
}

func TestJsonRequest(t *testing.T) {

	illusionmvc.RegisterHandler("/testJson", []string{httpmethod.POST}, Hello)
	illusionmvc.StartService()
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("D:\\temp"))))
	//
	//if err := http.ListenAndServe(":8081", nil); err != nil{
	//	fmt.Println(err)
	//}

}
