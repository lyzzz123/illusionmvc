package test

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/response"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestProtobufClient(t *testing.T) {
	client := &http.Client{Timeout: 5 * time.Second}
	student := &Student{}
	student.Address = "china"
	student.Name = "lyzzz"
	student.Age = 12
	student.Cn = ClassName_class1
	studentBytes, _ := proto.Marshal(student)
	resp, err := client.Post("http://localhost:9527/protobuf_test", "application/x-protobuf", bytes.NewBuffer(studentBytes))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	responseStudent := &Student{}
	proto.Unmarshal(result, responseStudent)
	fmt.Println(responseStudent)
}

func ProtobufTest(student *Student) (*response.ProtobufResponse, error) {
	fmt.Println(student)
	student.Age = student.Age + 10
	return &response.ProtobufResponse{Data: student}, nil
}

func TestProtobufServer(t *testing.T) {
	illusionmvc.RegisterHandler("/protobuf_test", []string{httpmethod.POST}, ProtobufTest)
	illusionmvc.StartService("9527")
}
