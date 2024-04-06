package exceptionhandler

import "net/http"

type DefaultBusinessExceptionHandler struct {
}

func (defaultBusinessExceptionHandler *DefaultBusinessExceptionHandler) Handle(writer http.ResponseWriter, err error) error {
	writer.WriteHeader(500)
	return nil
}
