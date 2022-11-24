package conf

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var ConfigMysql *ConfigStruct
var ConfigMongo *ConfigStruct

type DBStruct struct {
	Host     string
	Port     string
	Username string
	Passwd   string
	Database string
}

type ConfigStruct struct {
	Db         DBStruct
	DbTemplate string
}

func GetConfMysql() *ConfigStruct {
	return ConfigMysql
}
func GetConfMongo() *ConfigStruct {
	return ConfigMongo
}

func ParseConf() {
	PareseConfHelper("env-mysql", "mysql")
	PareseConfHelper("env-mongo", "mongo")
}

func PareseConfHelper(env string, db string) {
	v := viper.New()
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	environment := os.Getenv(env)
	absolutePath := filepath.Join(pwd, "conf", environment+".yaml")
	v.SetConfigFile(absolutePath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	var config *ConfigStruct
	err = v.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	if db == "mysql" {
		ConfigMysql = config
	}
	if db == "mongo" {
		ConfigMongo = config
	}

}
