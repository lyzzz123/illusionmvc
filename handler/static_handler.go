package handler

import (
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

type StaticHandler interface {
	Match(request *http.Request) bool
	HandleStatic(writer http.ResponseWriter, request *http.Request)
}

type DefaultStaticHandler struct {
	StaticPaths string

	StaticDir string
}

func (staticHandler *DefaultStaticHandler) Match(request *http.Request) bool {
	return strings.HasPrefix(request.URL.Path, staticHandler.StaticPaths)
}

func (staticHandler *DefaultStaticHandler) HandleStatic(writer http.ResponseWriter, request *http.Request) {

	relativePath := strings.TrimPrefix(request.URL.Path, staticHandler.StaticPaths)
	filePath := path.Join(staticHandler.StaticDir, relativePath)
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	fileStat, _ := file.Stat()
	http.ServeContent(writer, request, fileStat.Name(), fileStat.ModTime(), io.NewSectionReader(file, 0, fileStat.Size()))

}
