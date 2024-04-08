package requestconverter

import (
	"github.com/golang/protobuf/proto"
	"github.com/lyzzz123/illusionmvc/wrapper"
	"io/ioutil"
	"net/http"
)

type ApplicationProtobufConverter struct {
}

func (applicationProtobufConverter *ApplicationProtobufConverter) Convert(writer http.ResponseWriter, request *http.Request, param interface{}, inputWrapper *wrapper.InputWrapper) error {

	if bodyBytes, err := ioutil.ReadAll(request.Body); err != nil {
		return err
	} else {
		protoParma := param.(proto.Message)
		if err := proto.Unmarshal(bodyBytes, protoParma); err != nil {
			return err
		}
	}

	return nil
}

func (applicationProtobufConverter *ApplicationProtobufConverter) Support(request *http.Request) bool {

	if request.Header.Get("Content-Type") == "application/x-protobuf" {
		return true
	}
	return false
}

func (applicationProtobufConverter *ApplicationProtobufConverter) Name() string {
	return "ApplicationProtobufConverter"
}
