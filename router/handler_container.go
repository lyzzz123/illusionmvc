package router

import (
	"github.com/lyzzz123/illusionmvc/constant/httpmethod"
	"github.com/lyzzz123/illusionmvc/handler"
)

type Router struct {
	GetMapping *handler.WrapperMapping

	PostMapping *handler.WrapperMapping

	PutMapping *handler.WrapperMapping

	DeleteMapping *handler.WrapperMapping
}

func (router *Router) RegisterHandlerWrapper(wrapper *handler.Wrapper) {
	for _, method := range wrapper.HttpMethod {
		if method == httpmethod.GET {
			router.GetMapping.RegisterHandlerWrapper(wrapper)
		}
		if method == httpmethod.POST {
			router.PostMapping.RegisterHandlerWrapper(wrapper)
		}
		if method == httpmethod.PUT {
			router.PutMapping.RegisterHandlerWrapper(wrapper)
		}
		if method == httpmethod.DELETE {
			router.DeleteMapping.RegisterHandlerWrapper(wrapper)
		}
	}
}

func (router *Router) GetHandlerWrapper(method string, path string) *handler.Wrapper {
	if method == httpmethod.GET {
		return router.GetMapping.GetHandlerWrapper(path)
	}
	if method == httpmethod.POST {
		return router.PostMapping.GetHandlerWrapper(path)
	}
	if method == httpmethod.PUT {
		return router.PutMapping.GetHandlerWrapper(path)
	}
	if method == httpmethod.DELETE {
		return router.DeleteMapping.GetHandlerWrapper(path)
	}
	return nil
}
