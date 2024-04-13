package handler

type WrapperMapping struct {
	WrapperMapping map[string]*Wrapper

	PathValueWrapperMapping *PathTreeMap
}

func (wrapperMapping *WrapperMapping) GetHandlerWrapper(path string) *Wrapper {

	w, ok := wrapperMapping.WrapperMapping[path]
	if ok {
		return w
	}
	w = wrapperMapping.PathValueWrapperMapping.GetHandlerWrapper(path)
	return w
}

func (wrapperMapping *WrapperMapping) RegisterHandlerWrapper(wrapper *Wrapper) {

	if wrapper.HasPathValue {
		wrapperMapping.PathValueWrapperMapping.PutHandlerWrapper(wrapper.ReplacedPath, wrapper)
	} else {
		wrapperMapping.WrapperMapping[wrapper.ReplacedPath] = wrapper
	}
}
