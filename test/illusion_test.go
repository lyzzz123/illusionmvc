package test

import (
	"fmt"
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/filter"
	"reflect"
	"testing"
)

func TestToRegex(t *testing.T) {
	//th := &TestHandler{}
	illusion.RegisterHandler("/getTest/{t1}/{t2}", []string{httpmethod.GET}, new(TestHandler).GetTest)
	//illusion.RegisterHandler("/getTest/{t1}/{t2}", []string{httpmethod.POST}, new(TestHandler1).GetTest1)

	illusion.RegisterHandler("/getTest", []string{httpmethod.GET}, new(TestHandler2).GetTest2)
	illusion.RegisterHandler("/protobuf", []string{httpmethod.POST}, new(TestHandler).Protobuf)
	illusion.RegisterServiceListener(&TestListener{})

	illusion.RegisterFilter(&filter.CorsFilter{})

	illusion.StartService()
}

func TestSplit(t *testing.T) {
	ef := reflect.ValueOf(*new([]int)).Type().String()
	fmt.Println(ef)
}

type TS struct {
	Hello []string `json:"hello"`
	World string   `json:"world"`
}
