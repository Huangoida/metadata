package dal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"metadata/conf"
)

var mysqlDB *gorm.DB

func InitMysqlDb() {
	dsn := fmt.Sprintf(conf.GetConf().MysqlTemplate, conf.GetConf().Mysql.Username, conf.GetConf().Mysql.Passwd,
		conf.GetConf().Mysql.Host, conf.GetConf().Mysql.Port, conf.GetConf().Mysql.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	mysqlDB = db
}

func GetDb() *gorm.DB {
	return mysqlDB
}
