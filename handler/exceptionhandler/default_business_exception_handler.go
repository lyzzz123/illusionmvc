package exceptionhandler

import "net/http"

type DefaultBusinessExceptionHandler struct {
}

func (defaultBusinessExceptionHandler *DefaultBusinessExceptionHandler) Handle(writer http.ResponseWriter, err error) error {
	writer.WriteHeader(http.StatusInternalServerError)
	writer.Write([]byte("500 server internal error"))
	return nil
}
