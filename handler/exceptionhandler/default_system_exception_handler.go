package exceptionhandler

import "net/http"

type DefaultSystemExceptionHandler struct {
}

func (defaultSystemExceptionHandler *DefaultSystemExceptionHandler) Handle(writer http.ResponseWriter, err error) error {
	writer.WriteHeader(500)
	return nil
}
