package handler

import "github.com/lyzzz123/illusionmvc/constant/httpmethod"

type Container struct {
	GetMapping *WrapperMapping

	PostMapping *WrapperMapping

	PutMapping *WrapperMapping

	DeleteMapping *WrapperMapping
}

func (container *Container) RegisterWrapper(wrapper *Wrapper) {
	for _, method := range wrapper.HttpMethod {
		if method == httpmethod.GET {
			container.GetMapping.RegisterHandler(wrapper)
		}
		if method == httpmethod.POST {
			container.PostMapping.RegisterHandler(wrapper)
		}
		if method == httpmethod.PUT {
			container.PutMapping.RegisterHandler(wrapper)
		}
		if method == httpmethod.DELETE {
			container.DeleteMapping.RegisterHandler(wrapper)
		}
	}
}

func (container *Container) GetWrapper(method string, path string) *Wrapper {
	if method == httpmethod.GET {
		return container.GetMapping.GetHandler(path)
	}
	if method == httpmethod.POST {
		return container.PostMapping.GetHandler(path)
	}
	if method == httpmethod.PUT {
		return container.PutMapping.GetHandler(path)
	}
	if method == httpmethod.DELETE {
		return container.DeleteMapping.GetHandler(path)
	}
	return nil
}
