package conf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var ConfigDatabase *ConfigStruct

type DBStruct struct {
	Host     string
	Port     string
	Username string
	Passwd   string
	Database string

	DbTemplate string
}

type ConfigStruct struct {
	Mysql DBStruct
	Mongo DBStruct
}

func GetConfMysql() *DBStruct {
	return &ConfigDatabase.Mysql
}
func GetConfMongo() *DBStruct {
	return &ConfigDatabase.Mongo
}

func ParseConf() {
	v := viper.New()
	pwd, err := os.Getwd()
	fmt.Println(pwd)
	if err != nil {
		panic(err)
	}
	environment := os.Getenv("environment")
	absolutePath := filepath.Join(pwd, "conf", environment+".yaml")
	v.SetConfigFile(absolutePath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	var config *ConfigStruct
	print(v)
	err = v.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	ConfigDatabase = config
}
