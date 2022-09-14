package tlog

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	logger_category *logrus.Logger
)

func LoggerCategory(env, owner, NodeId string) {
	if logger_category != nil {
		return
	}
	logger_category = logrus.New()

	writerDebug, errDebug := rotatelogs.New(
		filePath+"-debug-full-response-request-"+NodeId+"-%Y%m%d.log",
		rotatelogs.WithMaxAge(time.Second*60*60*10), // 文件最大保存时间10天
		rotatelogs.WithRotationTime(time.Second*60), // 日志切割时间间隔
	)
	if errDebug != nil {
		panic("Init log failed, err:")
	}

	writer, err := rotatelogs.New(
		filePath+"-full-response-request-"+NodeId+"-%Y%m%d.log",
		rotatelogs.WithMaxAge(time.Second*60*60*10), // 文件最大保存时间10天
		rotatelogs.WithRotationTime(time.Second*60), // 日志切割时间间隔
	)
	if err != nil {
		panic("Init log failed, err:")
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{ // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
		logrus.DebugLevel: writerDebug,
	}, &logrus.JSONFormatter{})
	logger_category.AddHook(lfHook)

	logger_category.WithFields(logrus.Fields{
		"env":   env,
		"owner": owner,
	})
}

func TraceCategory(v ...interface{}) {
	logger_category.Trace(addMsg(v)...)
}

func InfoCategory(v ...interface{}) {
	logger_category.Info(addMsg(v)...)
}

func DebugCategory(v ...interface{}) {
	logger_category.Debug(addMsg(v)...)
}
func WarnCategory(v ...interface{}) {
	logger_category.Warn(addMsg(v)...)
}

func ErrorCategory(v ...interface{}) {
	logger_category.Error(addMsg(v)...)
}

func FatalCategory(v ...interface{}) {
	logger_category.Fatal(addMsg(v)...)
}

func PanicCategory(v ...interface{}) {
	logger_category.Panic(addMsg(v)...)
}

func SetLevelCategory(debug bool) {
	if debug {
		logger_category.SetLevel(logrus.DebugLevel)
	} else {
		logger_category.SetLevel(logrus.InfoLevel)
	}
}

func GetLoggerCategory() *logrus.Logger {
	return logger_category
}
