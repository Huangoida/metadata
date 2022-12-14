package mysql

import (
	"fmt"
	"metadata/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlDB *gorm.DB

func InitMysqlDb() {
	dsn := fmt.Sprintf(conf.GetConfMysql().DbTemplate, conf.GetConfMysql().Username, conf.GetConfMysql().Passwd,
		conf.GetConfMysql().Host, conf.GetConfMysql().Port, conf.GetConfMysql().Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	mysqlDB = db
}

func GetDb() *gorm.DB {
	return mysqlDB
}
