package util

import (
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"time"
)

type Logger struct {
	TempLogger *logrus.Logger
	Ctx        *gin.Context
}

var GlobalLogger *logrus.Logger

const (
	CtxLoggerName    = "CtxLogger"
	GlobalLoggerName = "GlobalLogger"
)

var LoggerList = []string{CtxLoggerName, GlobalLoggerName}

func GetCtxLogger(c *gin.Context) Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.SetLevel(logrus.InfoLevel) // 设置日志级别
	logger.SetReportCaller(false)     // 设置在输出日志中添加文件名和方法信息 默认关闭
	writer, _ := DivisionWriter(CtxLoggerName)
	logger.SetOutput(io.MultiWriter(writer))
	return Logger{TempLogger: logger, Ctx: c}
}

func (l *Logger) DoInfo(info string) {
	l.TempLogger.WithFields(logrus.Fields{
		"tracer_id": GetTracerId(l.Ctx),
	}).Info(info)
}

func (l *Logger) DoError(err string) {
	l.TempLogger.WithFields(logrus.Fields{
		"tracer_id": GetTracerId(l.Ctx),
	}).Error(err)
}

func (l *Logger) DoDebug(err string) {
	l.TempLogger.WithFields(logrus.Fields{
		"tracer_id": GetTracerId(l.Ctx),
	}).Debug(err)
}

func GetTracerId(c *gin.Context) string {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}
	return traceId
}

func GetLogFileName(name string) string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(pwd, "logs", name+".log")
}

func DivisionWriter(name string) (*rotatelogs.RotateLogs, error) {
	writer, err := rotatelogs.New(
		GetLogFileName(name)+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(GetLogFileName(name)),
		rotatelogs.WithMaxAge(time.Duration(72)*time.Hour),       // 保留最近3天的日志文件
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour), // 每隔1天轮转一个新文件
	)
	return writer, err
}
