package test

import (
	"fmt"
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/log"
	response2 "github.com/lyzzz123/illusionmvc/response"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)

type TestHandler struct {
}

func (testHandler *TestHandler) GetTest(testGetParam *TestGetParam) (*response2.JSONResponse, error) {
	log.Info("%v", testGetParam)
	return &response2.JSONResponse{"{\"hello\":\"world\"}"}, nil
}

type TestHandler1 struct {
}

func (testHandler1 *TestHandler1) GetTest1(testGetParam *TestGetParam) (*response2.JSONResponse, error) {
	log.Info("%v", testGetParam)
	return &response2.JSONResponse{"{\"hello\":\"world\"}"}, nil
}

type TestHandler2 struct {
}

func (testHandler2 *TestHandler2) GetTest2(testGetParam *TestGetParam) (*response2.JSONResponse, error) {
	log.Info("%v", testGetParam)
	//s := "{\"hello\":\"world\"}"
	return &response2.JSONResponse{&W{"world"}}, nil
}

type W struct {
	Hello string `json:"hello"`
}

type TestGetParam struct {
	T1 string `paramValue:"t1"`
	T2 int    `paramValue:"t2"`
}

type TestHandlerParam struct {
	HelloWorld string `paramValue:"helloWorld"`

	Asddf *multipart.FileHeader `paramValue:"asdddf"`

	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

func (testHandler *TestHandler) Upload(testHandlerParam *TestHandlerParam) (*response2.JSONResponse, error) {

	i := 0
	j := 3 / i
	fmt.Println(j)
	file, _ := testHandlerParam.Asddf.Open()

	newFile, _ := os.Create("D:\\temp\\" + testHandlerParam.Asddf.Filename)
	io.Copy(newFile, file)
	newFile.Close()
	file.Close()
	return &response2.JSONResponse{Data: "{\"hello\":\"upload success !!!!\"}"}, nil
}

func (testHandler *TestHandler) Download(testHandlerParam *TestHandlerParam) (*response2.FileResponse, error) {

	fr := &response2.FileResponse{}
	fileInfo, _ := os.Stat("D:\\ebook\\Kubernetes权威指南：从Docker到Kubernetes实践全接触（第4版）.pdf")
	file, _ := os.Open("D:\\ebook\\Kubernetes权威指南：从Docker到Kubernetes实践全接触（第4版）.pdf")
	fr.Name = fileInfo.Name()
	fr.Size = fileInfo.Size()
	fr.Reader = file

	return fr, nil
}

func (testHandler *TestHandler) PostJSON(postJSONParam *PostJSONParam) (*response2.JSONResponse, error) {

	return &response2.JSONResponse{"{\"hello\":\"" + postJSONParam.T2 + "\"}"}, nil

}

type PostJSONParam struct {
	T1 int    `pathValue:"t1"`
	T2 string `pathValue:"t2"`
	T3 string `json:"t3"`
}

func (testHandler *TestHandler) Protobuf(student *Student) (*response2.ProtobufResponse, error) {
	fmt.Println(student)

	student.Age = student.Age + 10
	return &response2.ProtobufResponse{Data: student}, nil
}

func TestToRegexd(t *testing.T) {
	ttt := &TestHandler{}
	illusionmvc.RegisterHandler("/getTest", []string{httpmethod.POST}, ttt.Upload)
	illusionmvc.StartService()
}
