package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
)

type Formatter struct {
}

func (t Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	//字节缓冲区
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:06")
	//自定义文件路径
	pc, file, line, _ := runtime.Caller(10)
	//自定义输出格式
	fmt.Fprintf(b, "[%s][%s][%s][%s:%d] - %s\n", timestamp, entry.Level, runtime.FuncForPC(pc).Name(), path.Base(file), line, entry.Message)
	return b.Bytes(), nil
}

type DefaultLog struct {
}

func (defaultLog *DefaultLog) InitLog() error {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&Formatter{})
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
