package tlog

import (
	"fmt"
	"path"
	"runtime"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

const (
	filePath = "log/logrus"
)

var (
	logger *logrus.Logger
)

func Logger(env, owner, NodeId string) {
	if logger != nil {
		return
	}
	LoggerCategory(env, owner, NodeId)
	logger = logrus.New()

	writerDebug, errDebug := rotatelogs.New(
		filePath+"-debug-"+NodeId+"-%Y%m%d.log",
		rotatelogs.WithMaxAge(time.Second*60*60*10), // 文件最大保存时间10天
		rotatelogs.WithRotationTime(time.Second*60), // 日志切割时间间隔
	)
	if errDebug != nil {
		panic("Init log failed, err:")
	}

	writer, err := rotatelogs.New(
		filePath+"-"+NodeId+"-%Y%m%d.log",
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
	logger.AddHook(lfHook)

	logger.WithFields(logrus.Fields{
		"env":   env,
		"owner": owner,
	})
}

func Trace(v ...interface{}) {
	logger.Trace(addMsg(v)...)
}

func Info(v ...interface{}) {
	logger.Info(addMsg(v)...)
}

func Debug(v ...interface{}) {
	logger.Debug(addMsg(v)...)
}
func Warn(v ...interface{}) {
	logger.Warn(addMsg(v)...)
}

func Error(v ...interface{}) {
	logger.Error(addMsg(v)...)
}

func Fatal(v ...interface{}) {
	logger.Fatal(addMsg(v)...)
}

func Panic(v ...interface{}) {
	logger.Panic(addMsg(v)...)
}

func addMsg(v []interface{}) []interface{} {
	msg := make([]interface{}, len(v)+1)
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	_, filename := path.Split(file)
	info := "(" + filename + ":" + strconv.Itoa(line) + ") "
	msg[0] = info
	if len(v) > 0 {
		if ctx, ok := v[0].(*gin.Context); ok {
			if requestID, exist := ctx.Get("X-Request-ID"); exist {
				v[0] = fmt.Sprintf("request-id: %d ", requestID)
			}
		}
	}

	copy(msg[1:], v)
	return msg
}

func SetLevel(debug bool) {
	if debug {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
}

func GetLogger() *logrus.Logger {
	return logger
}
