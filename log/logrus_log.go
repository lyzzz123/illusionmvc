package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

type DefaultLog struct {
}

func (defaultLog *DefaultLog) Init() error {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	return nil
}

func (defaultLog *DefaultLog) Debug(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func (defaultLog *DefaultLog) Info(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func (defaultLog *DefaultLog) Warn(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func (defaultLog *DefaultLog) Error(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}
