package exceptionhandler

import "net/http"

type DefaultSystemExceptionHandler struct {
}

func (defaultSystemExceptionHandler *DefaultSystemExceptionHandler) Handle(writer http.ResponseWriter, err error) error {
	writer.WriteHeader(http.StatusInternalServerError)
	writer.Write([]byte("500 server internal error"))
	return nil
}
