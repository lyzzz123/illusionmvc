package service

import (
	"github.com/lyzzz123/illusionmvc/converter/requestconverter"
	"github.com/lyzzz123/illusionmvc/converter/typeconverter"
	"github.com/lyzzz123/illusionmvc/filter"
	"github.com/lyzzz123/illusionmvc/handler"
	"github.com/lyzzz123/illusionmvc/handler/exceptionhandler"
	"github.com/lyzzz123/illusionmvc/log"
	"github.com/lyzzz123/illusionmvc/response"
	"net/http"
	"reflect"
	"regexp"
	"sort"
)

type IllusionService struct {
	TypeConverterMap map[reflect.Type]typeconverter.Converter

	DefaultSystemExceptionHandler exceptionhandler.ExceptionHandler

	DefaultBusinessExceptionHandler exceptionhandler.ExceptionHandler

	RequestConverterArray []requestconverter.RequestConverter

	ResponseWriterMap map[reflect.Type]response.Writer

	filterArray []filter.Filter

	handlerContainer *handler.Container

	DefaultStaticHandler handler.StaticHandler
}

func (illusionService *IllusionService) RegisterResponseWriter(responseWriter response.Writer) {
	if illusionService.ResponseWriterMap == nil {
		illusionService.ResponseWriterMap = make(map[reflect.Type]response.Writer)
	}
	illusionService.ResponseWriterMap[responseWriter.Support()] = responseWriter
}

func (illusionService *IllusionService) RegisterRequestConverter(requestConverter requestconverter.RequestConverter) {
	illusionService.RequestConverterArray = append(illusionService.RequestConverterArray, requestConverter)
}

func (illusionService *IllusionService) RegisterSystemExceptionHandler(exceptionHandler exceptionhandler.ExceptionHandler) {
	illusionService.DefaultSystemExceptionHandler = exceptionHandler
}

func (illusionService *IllusionService) RegisterBusinessExceptionHandler(exceptionHandler exceptionhandler.ExceptionHandler) {
	illusionService.DefaultBusinessExceptionHandler = exceptionHandler
}

func (illusionService *IllusionService) RegisterTypeConverter(converter typeconverter.Converter) {
	if illusionService.TypeConverterMap == nil {
		illusionService.TypeConverterMap = make(map[reflect.Type]typeconverter.Converter)
	}
	illusionService.TypeConverterMap[converter.Support()] = converter
}

func (illusionService *IllusionService) RegisterFilter(filter filter.Filter) {
	illusionService.filterArray = append(illusionService.filterArray, filter)
	sort.SliceStable(illusionService.filterArray, func(i, j int) bool {
		return illusionService.filterArray[i].GetPriority() > illusionService.filterArray[j].GetPriority()
	})
}

func (illusionService *IllusionService) RegisterHandler(path string, httpMethod []string, handlerMethod interface{}) {

	if illusionService.handlerContainer == nil {
		illusionService.handlerContainer = &handler.Container{
			GetMapping: &handler.WrapperMapping{
				WrapperMapping:        make(map[string]*handler.Wrapper),
				PathValueWrapperArray: make([]*handler.Wrapper, 0),
			},
			PostMapping: &handler.WrapperMapping{
				WrapperMapping:        make(map[string]*handler.Wrapper),
				PathValueWrapperArray: make([]*handler.Wrapper, 0),
			},
			PutMapping: &handler.WrapperMapping{
				WrapperMapping:        make(map[string]*handler.Wrapper),
				PathValueWrapperArray: make([]*handler.Wrapper, 0),
			},
			DeleteMapping: &handler.WrapperMapping{
				WrapperMapping:        make(map[string]*handler.Wrapper),
				PathValueWrapperArray: make([]*handler.Wrapper, 0),
			},
		}
	}

	wrapper := handler.CreateHandlerWrapper(path, httpMethod, handlerMethod)
	for _, f := range illusionService.filterArray {
		pathPattern := f.GetPathPattern()
		doubleStarRegex := regexp.MustCompile("/\\*\\*")
		pathPattern = doubleStarRegex.ReplaceAllString(pathPattern, "/.+")
		singleStarRegex := regexp.MustCompile("/\\*")
		pathPattern = singleStarRegex.ReplaceAllString(pathPattern, "/[a-zA-Z\\d]+")
		pathPatternRegex := regexp.MustCompile(pathPattern)
		pathPatternRegex.MatchString(path)
		wrapper.FilterArray = append(wrapper.FilterArray, f)
	}
	wrapper.Input.TypeConverterMap = illusionService.TypeConverterMap
	wrapper.DefaultExceptionHandler = illusionService.DefaultBusinessExceptionHandler
	wrapper.RequestConverterArray = illusionService.RequestConverterArray
	responseWriter, ok := illusionService.ResponseWriterMap[wrapper.OutputType]
	if ok {
		wrapper.ResponseWriter = responseWriter
	} else {
		panic("not support response type:" + wrapper.OutputType.String())
	}
	illusionService.handlerContainer.RegisterWrapper(wrapper)

}

func (illusionService *IllusionService) RegisterStaticHandler(staticHandler handler.StaticHandler) {
	illusionService.DefaultStaticHandler = staticHandler
}

func (illusionService *IllusionService) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	defer func() {
		if msg := recover(); msg != nil {
			log.Error("%v", msg)
			illusionService.DefaultSystemExceptionHandler.Handle(writer, nil)
		}
	}()

	if illusionService.DefaultStaticHandler.Match(request) {
		illusionService.DefaultStaticHandler.HandleStatic(writer, request)
	} else {
		wrapper := illusionService.handlerContainer.GetWrapper(request.Method, request.URL.Path)
		for _, f := range wrapper.FilterArray {
			if err := f.PreHandle(writer, request); err != nil {
				log.Error("execute filter preHandle error:%s", err.Error())
				writer.WriteHeader(500)
			}
		}
		if err := wrapper.Handle(writer, request); err != nil {
			log.Error("execute handler error:%s", err.Error())
			writer.WriteHeader(500)
			return
		}

		for _, f := range wrapper.FilterArray {
			if err := f.PostHandle(writer); err != nil {
				log.Error("execute filter postHandle error:%s", err.Error())
				writer.WriteHeader(500)
			}
		}
	}

}

func (illusionService *IllusionService) Start(port string) {
	if port == "" {
		port = "8080"
	}
	if err := http.ListenAndServe(":"+port, illusionService); err != nil {
		panic(err)
	}
}
