package listener

import (
	"github.com/lyzzz123/illusionmvc"
	"github.com/lyzzz123/illusionmvc/log"
	"testing"
)

type ListenerTest struct {
}

func (listenerTest *ListenerTest) PreRun() error {
	log.Info("PreRun")
	return nil
}

func (listenerTest *ListenerTest) PostRun() error {
	log.Info("PostRun")
	return nil
}

func (listenerTest *ListenerTest) GetPriority() int {
	return 5
}

func TestListener(t *testing.T) {
	illusionmvc.RegisterListener(&ListenerTest{})
	illusionmvc.SetManualShutdown(true)
	illusionmvc.StartService("9527")

}
