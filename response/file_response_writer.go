package response

import (
	"io"
	"net/http"
	"reflect"
	"strconv"
)

type FileResponse struct {
	Name   string
	Size   int64
	Reader io.Reader
}

type FileResponseWriter struct {
	ResponseType reflect.Type
}

func (fileResponseWriter *FileResponseWriter) Write(writer http.ResponseWriter, returnValue interface{}) error {

	fileReturnValue := returnValue.(*FileResponse)
	if fileReturnValue != nil {
		writer.Header().Set("Content-Disposition", "attachment; filename="+fileReturnValue.Name)
		writer.Header().Set("Content-Length", strconv.FormatInt(fileReturnValue.Size, 10))
		io.Copy(writer, fileReturnValue.Reader)
	}
	return nil
}

func (fileResponseWriter *FileResponseWriter) Support() reflect.Type {
	return fileResponseWriter.ResponseType
}
