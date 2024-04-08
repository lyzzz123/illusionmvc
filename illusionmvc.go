package illusionmvc

import (
	"github.com/lyzzz123/illusionmvc/converter/typeconverter"
	"github.com/lyzzz123/illusionmvc/filter"
	"github.com/lyzzz123/illusionmvc/handler"
	"github.com/lyzzz123/illusionmvc/handler/exceptionhandler"
	"github.com/lyzzz123/illusionmvc/log"
	requestconverter2 "github.com/lyzzz123/illusionmvc/request/requestconverter"
	"github.com/lyzzz123/illusionmvc/response"
	"github.com/lyzzz123/illusionmvc/service"
	"reflect"
)

var illusionService = &service.IllusionService{}

func init() {
	log.RegisterLog(&log.LogrusLog{})
	illusionService.RegisterTypeConverter(&typeconverter.BoolConvert{})
	illusionService.RegisterTypeConverter(&typeconverter.BoolPtrConvert{})
	illusionService.RegisterTypeConverter(&typeconverter.Float32Convert{})
	illusionService.RegisterTypeConverter(&typeconverter.Float32PtrConvert{})
	illusionService.RegisterTypeConverter(&typeconverter.Float64Convert{})
	illusionService.RegisterTypeConverter(&typeconverter.Float64PtrConvert{})
	illusionService.RegisterTypeConverter(&typeconverter.Int8Converter{})
	illusionService.RegisterTypeConverter(&typeconverter.Int8PtrConverter{})
	illusionService.RegisterTypeConverter(&typeconverter.Int16Converter{})
	illusionService.RegisterTypeConverter(&typeconverter.Int16PtrConverter{})
	illusionService.RegisterTypeConverter(&typeconverter.Int32Converter{})
	illusionService.RegisterTypeConverter(&typeconverter.Int32PtrConverter{})
	illusionService.RegisterTypeConverter(&typeconverter.Int64Converter{})
	illusionService.RegisterTypeConverter(&typeconverter.Int64PtrConverter{})
	illusionService.RegisterTypeConverter(&typeconverter.IntConverter{})
	illusionService.RegisterTypeConverter(&typeconverter.IntPtrConverter{})
	illusionService.RegisterTypeConverter(&typeconverter.StringConvert{})
	illusionService.RegisterTypeConverter(&typeconverter.StringPtrConvert{})
	illusionService.RegisterTypeConverter(&typeconverter.Uint8Converter{})
	illusionService.RegisterTypeConverter(&typeconverter.Uint8PtrConverter{})
	illusionService.RegisterTypeConverter(&typeconverter.Uint16Converter{})
	illusionService.RegisterTypeConverter(&typeconverter.Uint16PtrConverter{})
	illusionService.RegisterTypeConverter(&typeconverter.Uint32Converter{})
	illusionService.RegisterTypeConverter(&typeconverter.Uint32PtrConverter{})
	illusionService.RegisterTypeConverter(&typeconverter.Uint64Converter{})
	illusionService.RegisterTypeConverter(&typeconverter.Uint64PtrConverter{})
	illusionService.RegisterTypeConverter(&typeconverter.UintConverter{})
	illusionService.RegisterTypeConverter(&typeconverter.UintPtrConverter{})

	illusionService.RegisterRequestConverter(&requestconverter2.ApplicationJSONConverter{})
	illusionService.RegisterRequestConverter(&requestconverter2.ApplicationXWWWFormUrlencodedConverter{})
	illusionService.RegisterRequestConverter(&requestconverter2.MultipartFormDataConverter{})
	illusionService.RegisterRequestConverter(&requestconverter2.ApplicationProtobufConverter{})

	illusionService.RegisterResponseWriter(&response.FileResponseWriter{ResponseType: reflect.TypeOf(*new(response.FileResponse))})
	illusionService.RegisterResponseWriter(&response.JSONResponseWriter{ResponseType: reflect.TypeOf(*new(response.JSONResponse))})
	illusionService.RegisterResponseWriter(&response.ProtobufResponseWriter{ResponseType: reflect.TypeOf(*new(response.ProtobufResponse))})

	illusionService.RegisterBusinessExceptionHandler(&exceptionhandler.DefaultBusinessExceptionHandler{})
	illusionService.RegisterSystemExceptionHandler(&exceptionhandler.DefaultSystemExceptionHandler{})

}

func RegisterRequestConverter(requestConverter requestconverter2.RequestConverter) {
	illusionService.RegisterRequestConverter(requestConverter)
}

func RegisterTypeConverter(converter typeconverter.Converter) {
	illusionService.RegisterTypeConverter(converter)
}

func RegisterFilter(serviceFilter filter.Filter) {
	illusionService.RegisterFilter(serviceFilter)
}

func RegisterHandler(path string, httpMethod []string, handlerMethod interface{}) {
	illusionService.RegisterHandler(path, httpMethod, handlerMethod)
}

func RegisterStaticHandler(staticHandler handler.StaticHandler) {
	illusionService.RegisterStaticHandler(staticHandler)
}

func RegisterLog(l log.Log) {
	log.RegisterLog(l)
}

func StartService(port string) {
	illusionService.Start(port)
}
