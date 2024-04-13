package illusionmvc

import (
	"github.com/lyzzz123/illusionmvc/controller"
	"github.com/lyzzz123/illusionmvc/converter/typeconverter"
	"github.com/lyzzz123/illusionmvc/filter"
	"github.com/lyzzz123/illusionmvc/listener"
	"github.com/lyzzz123/illusionmvc/log"
	"github.com/lyzzz123/illusionmvc/request/requestconverter"
	"github.com/lyzzz123/illusionmvc/response"
	"reflect"
)

type Runner struct {
	Port string `property:"server.port"`
}

func (runner *Runner) AfterRunAction(objectContainer map[reflect.Type]interface{}) error {

	for _, registerObject := range objectContainer {

		if registerConverter, ok := registerObject.(typeconverter.Converter); ok {
			RegisterTypeConverter(registerConverter)
		}

		if registerFilter, ok := registerObject.(filter.Filter); ok {
			RegisterFilter(registerFilter)
		}

		if registerListener, ok := registerObject.(listener.Listener); ok {
			RegisterListener(registerListener)
		}

		if registerLog, ok := registerObject.(log.Log); ok {
			RegisterLog(registerLog)
		}

		if registerRequestConverter, ok := registerObject.(requestconverter.RequestConverter); ok {
			RegisterRequestConverter(registerRequestConverter)
		}

		if registerResponseWriter, ok := registerObject.(response.Writer); ok {
			RegisterResponseWriter(registerResponseWriter)
		}

		if registerResponseWriter, ok := registerObject.(response.Writer); ok {
			RegisterResponseWriter(registerResponseWriter)
		}

		if controllerObject, ok := registerObject.(controller.Controller); ok {
			RegisterController(controllerObject)
		}
	}
	StartService(runner.Port)
	return nil
}
