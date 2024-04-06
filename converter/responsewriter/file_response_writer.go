package responsewriter

import (
	response2 "github.com/lyzzz123/illusionmvc/response"
	"io"
	"net/http"
	"reflect"
	"strconv"
)

var fileResponseType = reflect.TypeOf(new(response2.FileResponse))

type FileResponseWriter struct {
}

func (fileResponseWriter *FileResponseWriter) Write(writer http.ResponseWriter, returnValue interface{}) error {

	fileReturnValue := returnValue.(*response2.FileResponse)
	if fileReturnValue != nil {
		writer.Header().Set("Content-Disposition", "attachment; filename="+fileReturnValue.Name)
		writer.Header().Set("Content-Length", strconv.FormatInt(fileReturnValue.Size, 10))
		io.Copy(writer, fileReturnValue.Reader)
	}
	return nil
}

func (fileResponseWriter *FileResponseWriter) Support() reflect.Type {
	return fileResponseType
}
