package test

import (
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/response"
	"os"
	"testing"
)

type DownloadParam struct {
	Hello string `paramValue:"hello"`
	Name  string `paramValue:"name"`
}

func Download(param *DownloadParam) (*response.FileResponse, error) {
	fr := &response.FileResponse{}
	fileInfo, _ := os.Stat("F:\\temp\\test.txt")
	file, _ := os.Open("F:\\temp\\test.txt")
	fr.Name = fileInfo.Name()
	fr.Size = fileInfo.Size()
	fr.File = file
	return fr, nil
}

func TestDownload(t *testing.T) {
	illusionmvc.RegisterHandler("/download", []string{httpmethod.GET}, Download)
	illusionmvc.StartService("9527")
}
