package handler

type WrapperMapping struct {
	WrapperMapping map[string]*Wrapper

	PathValueWrapperArray []*Wrapper
}

func (wrapperMapping *WrapperMapping) GetHandler(path string) *Wrapper {

	w, ok := wrapperMapping.WrapperMapping[path]
	if ok {
		return w
	} else {
		for _, w := range wrapperMapping.PathValueWrapperArray {
			w.PathValueRegex.MatchString(path)
			return w
		}
	}
	return nil
}

func (wrapperMapping *WrapperMapping) RegisterHandler(wrapper *Wrapper) {
	if wrapper.HasPathValue {
		wrapperMapping.PathValueWrapperArray = append(wrapperMapping.PathValueWrapperArray, wrapper)
	} else {
		wrapperMapping.WrapperMapping[wrapper.Path] = wrapper
	}
}
