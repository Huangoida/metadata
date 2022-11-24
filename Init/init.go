package Init

import (
	"io"
	"metadata/conf"
	"metadata/dal/mongo"
	"metadata/dal/mysql"
	"metadata/util"

	"github.com/sirupsen/logrus"
)

func InitConfig() {
	conf.ParseConf()
	mysql.InitMysqlDb()
	mongo.InitMangoDb()
	initGlobalLogger()
}

func initGlobalLogger() {
	util.GlobalLogger = logrus.New()
	util.GlobalLogger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	util.GlobalLogger.SetLevel(logrus.InfoLevel) // 设置日志级别
	util.GlobalLogger.SetReportCaller(false)     // 设置在输出日志中添加文件名和方法信息 默认关闭
	writer, err := util.DivisionWriter(util.GlobalLoggerName)
	if err != nil {
		panic(err)
	}
	util.GlobalLogger.SetOutput(io.MultiWriter(writer))
}
