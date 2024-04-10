package service

import (
	"github.com/lyzzz123/illusionmvc/converter/typeconverter"
	"github.com/lyzzz123/illusionmvc/filter"
	"github.com/lyzzz123/illusionmvc/handler"
	"github.com/lyzzz123/illusionmvc/handler/exceptionhandler"
	"github.com/lyzzz123/illusionmvc/log"
	"github.com/lyzzz123/illusionmvc/request/requestconverter"
	"github.com/lyzzz123/illusionmvc/response"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"regexp"
	"sort"
	"syscall"
)

type IllusionService struct {
	TypeConverterMap map[reflect.Type]typeconverter.Converter

	DefaultSystemExceptionHandler exceptionhandler.ExceptionHandler

	DefaultBusinessExceptionHandler exceptionhandler.ExceptionHandler

	RequestConverterMap map[string]requestconverter.RequestConverter

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
	if illusionService.RequestConverterMap == nil {
		illusionService.RequestConverterMap = make(map[string]requestconverter.RequestConverter)
	}
	illusionService.RequestConverterMap[requestConverter.Name()] = requestConverter
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
		pathPattern = "^" + singleStarRegex.ReplaceAllString(pathPattern, "/[^/]+") + "$"
		pathPatternRegex := regexp.MustCompile(pathPattern)
		if pathPatternRegex.MatchString(path) {
			wrapper.FilterArray = append(wrapper.FilterArray, f)
		}

	}
	wrapper.Input.TypeConverterMap = illusionService.TypeConverterMap
	wrapper.DefaultExceptionHandler = illusionService.DefaultBusinessExceptionHandler
	wrapper.RequestConverterMap = illusionService.RequestConverterMap
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

func (illusionService *IllusionService) RegisterLog(logInstance log.Log) {
	log.RegisterLog(logInstance)
}

func (illusionService *IllusionService) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	defer func() {
		if msg := recover(); msg != nil {
			log.Error("%v", msg)
			illusionService.DefaultSystemExceptionHandler.Handle(writer, nil)
		}
	}()

	if illusionService.DefaultStaticHandler != nil && illusionService.DefaultStaticHandler.Match(request) {
		illusionService.DefaultStaticHandler.HandleStatic(writer, request)
	} else {
		wrapper := illusionService.handlerContainer.GetWrapper(request.Method, request.URL.Path)
		if wrapper == nil {
			writer.WriteHeader(http.StatusNotFound)
			writer.Write([]byte("404 Not Found"))
			return
		}
		for i := 0; i < len(wrapper.FilterArray); i++ {
			if err := wrapper.FilterArray[i].PreHandle(writer, request); err != nil {
				log.Error("execute filter preHandle error:%s", err.Error())
				panic(err)
			}
		}

		if err := wrapper.Handle(writer, request); err != nil {
			log.Error("execute handler error:%s", err.Error())
			panic(err)
		}

		for i := len(wrapper.FilterArray) - 1; i >= 0; i-- {
			if err := wrapper.FilterArray[i].PostHandle(writer, request); err != nil {
				log.Error("execute filter preHandle error:%s", err.Error())
				panic(err)
			}
		}
	}
}

func (illusionService *IllusionService) Start(port string) {

	go func() {
		if port == "" {
			port = "8080"
		}
		log.Info("service started at port %v", port)
		if err := http.ListenAndServe(":"+port, illusionService); err != nil {
			panic(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-quit
	log.Info("ShutDown Server ...")

}
