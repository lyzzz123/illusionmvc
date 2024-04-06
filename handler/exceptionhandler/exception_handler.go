package exceptionhandler

import "net/http"

type ExceptionHandler interface {
	Handle(writer http.ResponseWriter, err error) error
}
