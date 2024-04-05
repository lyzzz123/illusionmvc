package illusionmvc

import (
	"github.com/lyzzz123/illusionmvc/config"
	"github.com/lyzzz123/illusionmvc/converter/requestconverter"
	"github.com/lyzzz123/illusionmvc/converter/responsewriter"
	"github.com/lyzzz123/illusionmvc/converter/typeconverter"
	"github.com/lyzzz123/illusionmvc/filter"
	"github.com/lyzzz123/illusionmvc/handler"
	"github.com/lyzzz123/illusionmvc/handler/exceptionhandler"
	"github.com/lyzzz123/illusionmvc/listener"
	"github.com/lyzzz123/illusionmvc/log"
	"github.com/lyzzz123/illusionmvc/service"
	"net/http"
	"reflect"
)

func init() {

	log.RegisterLog(&log.LogrusLog{})

	typeconverter.RegisterTypeConverter(reflect.TypeOf(*new(int)), typeconverter.IntConvert)
	typeconverter.RegisterTypeConverter(reflect.TypeOf(*new(int8)), typeconverter.Int8Convert)
	typeconverter.RegisterTypeConverter(reflect.TypeOf(*new(int16)), typeconverter.Int16Convert)
	typeconverter.RegisterTypeConverter(reflect.TypeOf(*new(int32)), typeconverter.Int32Convert)
	typeconverter.RegisterTypeConverter(reflect.TypeOf(*new(int64)), typeconverter.Int64Convert)

	typeconverter.RegisterTypeConverter(reflect.TypeOf(*new(uint)), typeconverter.UintConvert)
	typeconverter.RegisterTypeConverter(reflect.TypeOf(*new(uint8)), typeconverter.UintConvert)
	typeconverter.RegisterTypeConverter(reflect.TypeOf(*new(uint16)), typeconverter.Uint16Convert)
	typeconverter.RegisterTypeConverter(reflect.TypeOf(*new(uint32)), typeconverter.Uint32Convert)
	typeconverter.RegisterTypeConverter(reflect.TypeOf(*new(uint64)), typeconverter.Uint64Convert)

	typeconverter.RegisterTypeConverter(reflect.TypeOf(*new(bool)), typeconverter.BoolConvert)

	typeconverter.RegisterTypeConverter(reflect.TypeOf(*new(float64)), typeconverter.Float64Convert)
	typeconverter.RegisterTypeConverter(reflect.TypeOf(*new(float32)), typeconverter.Float32Convert)

	typeconverter.RegisterTypeConverter(reflect.TypeOf(*new(string)), typeconverter.StringConvert)

	typeconverter.RegisterTypeConverter(reflect.TypeOf(new(int)), typeconverter.IntPtrConvert)
	typeconverter.RegisterTypeConverter(reflect.TypeOf(new(int8)), typeconverter.Int8PtrConvert)
	typeconverter.RegisterTypeConverter(reflect.TypeOf(new(int16)), typeconverter.Int16PtrConvert)
	typeconverter.RegisterTypeConverter(reflect.TypeOf(new(int32)), typeconverter.Int32PtrConvert)
	typeconverter.RegisterTypeConverter(reflect.TypeOf(new(int64)), typeconverter.Int64PtrConvert)

	typeconverter.RegisterTypeConverter(reflect.TypeOf(new(uint)), typeconverter.UintPtrConvert)
	typeconverter.RegisterTypeConverter(reflect.TypeOf(new(uint8)), typeconverter.Uint8PtrConvert)
	typeconverter.RegisterTypeConverter(reflect.TypeOf(new(uint16)), typeconverter.Uint16PtrConvert)
	typeconverter.RegisterTypeConverter(reflect.TypeOf(new(uint32)), typeconverter.Uint32PtrConvert)
	typeconverter.RegisterTypeConverter(reflect.TypeOf(new(uint64)), typeconverter.Uint64PtrConvert)

	typeconverter.RegisterTypeConverter(reflect.TypeOf(new(bool)), typeconverter.BoolPtrConvert)
	typeconverter.RegisterTypeConverter(reflect.TypeOf(new(float64)), typeconverter.Float64PtrConvert)
	typeconverter.RegisterTypeConverter(reflect.TypeOf(new(float32)), typeconverter.Float64PtrConvert)
	typeconverter.RegisterTypeConverter(reflect.TypeOf(new(string)), typeconverter.StringPtrConvert)

	requestconverter.RegisterRequestConverter(&requestconverter.ApplicationJSONConverter{})
	requestconverter.RegisterRequestConverter(&requestconverter.ApplicationXWWWFormUrlencodedConverter{})
	requestconverter.RegisterRequestConverter(&requestconverter.MultipartFormDataConverter{})
	requestconverter.RegisterRequestConverter(&requestconverter.GetMethodConverter{})
	requestconverter.RegisterRequestConverter(&requestconverter.ApplicationProtobufConverter{})

	responsewriter.RegisterResponseWriter(&responsewriter.FileResponseWriter{})
	responsewriter.RegisterResponseWriter(&responsewriter.JSONResponseWriter{})
	responsewriter.RegisterResponseWriter(&responsewriter.ProtobufResponseWriter{})

	exceptionhandler.RegisterExceptionHandler(&exceptionhandler.DefaultExceptionHandler{})

}

func RegisterResponseWriter(responseWriter responsewriter.ResponseWriter) {
	responsewriter.RegisterResponseWriter(responseWriter)
}

func RegisterRequestConverter(requestConverter requestconverter.RequestConverter) {
	requestconverter.RegisterRequestConverter(requestConverter)
}
func RegisterTypeConverter(typ reflect.Type, converterFunc func(string) (interface{}, error)) {
	typeconverter.RegisterTypeConverter(typ, converterFunc)
}

func RegisterFilter(serviceFilter filter.Filter) {
	filter.RegisterFilter(serviceFilter)
}

func RegisterHandler(path string, httpMethod []string, handlerMethod interface{}) {
	handler.RegisterHandler(path, httpMethod, handlerMethod)
}

func RegisterServiceListener(listen listener.Listener) {
	listener.RegisterServiceListener(listen)
}

func RegisterLog(l log.Log) {
	log.RegisterLog(l)
}

func RegisterExceptionHandler(exceptionHandler exceptionhandler.ExceptionHandler) {
	exceptionhandler.RegisterExceptionHandler(exceptionHandler)
}

func StartService() {

	if err := listener.ExecuteHttpServerStartUpListener(); err != nil {
		panic(err)
	}

	port := config.GetConfig("port", "8080")

	if err := http.ListenAndServe(":"+port.(string), &service.IllusionService{}); err != nil {
		panic(err)
	}

	if err := listener.ExecuteHttpServerShutdownListener(); err != nil {
		panic(err)
	}

}
