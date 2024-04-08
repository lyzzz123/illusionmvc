package response

import (
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
)

type FileResponse struct {
	Name string
	Size int64
	File *os.File
}

type FileResponseWriter struct {
	ResponseType reflect.Type
}

func (fileResponseWriter *FileResponseWriter) Write(writer http.ResponseWriter, returnValue interface{}) error {

	fileReturnValue := returnValue.(*FileResponse)
	if fileReturnValue != nil {
		writer.Header().Set("Content-Disposition", "attachment; filename="+fileReturnValue.Name)
		writer.Header().Set("Content-Length", strconv.FormatInt(fileReturnValue.Size, 10))
		io.Copy(writer, fileReturnValue.File)
		if err := fileReturnValue.File.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (fileResponseWriter *FileResponseWriter) Support() reflect.Type {
	return fileResponseWriter.ResponseType
}
