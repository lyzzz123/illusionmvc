package controller

type Exporter struct {
	Path          string
	HttpMethod    []string
	HandlerMethod interface{}
}

type Controller interface {
	Export() []*Exporter
}
