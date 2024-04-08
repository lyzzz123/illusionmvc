package request_converter

import (
	"encoding/xml"
	"github.com/lyzzz123/illusionmvc/wrapper"
	"io/ioutil"
	"net/http"
)

type XmlConverter struct {
}

func (xmlConverter *XmlConverter) Convert(writer http.ResponseWriter, request *http.Request, param interface{}, inputWrapper *wrapper.InputWrapper) error {

	if bodyBytes, err := ioutil.ReadAll(request.Body); err != nil {
		return err
	} else {
		if err := xml.Unmarshal(bodyBytes, param); err != nil {
			return err
		}
	}
	return nil
}

func (xmlConverter *XmlConverter) Support(request *http.Request) bool {

	if request.Header.Get("Content-Type") == "application/xml" {
		return true
	}

	return false
}

func (xmlConverter *XmlConverter) Name() string {
	return "XmlConverter"
}
