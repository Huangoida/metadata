package mysql

import (
	"fmt"
	"metadata/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlDB *gorm.DB

func InitMysqlDb() {
	// println(conf.GetConf().Mysql.Database)
	dsn := fmt.Sprintf(conf.GetConf().MysqlTemplate, conf.GetConf().Mysql.Username, conf.GetConf().Mysql.Passwd,
		conf.GetConf().Mysql.Host, conf.GetConf().Mysql.Port, conf.GetConf().Mysql.Database)
	println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	mysqlDB = db
}

func GetDb() *gorm.DB {
	return mysqlDB
}
