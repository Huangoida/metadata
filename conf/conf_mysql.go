package conf

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var Config *ConfigStruct

type MysqlStruct struct {
	Host     string
	Port     string
	Username string
	Passwd   string
	Database string
}

type ConfigStruct struct {
	Mysql         MysqlStruct
	MysqlTemplate string
}

func GetConf() *ConfigStruct {
	return Config
}

func ParseConf() {
	v := viper.New()
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	os.Setenv("env-mysql", "mysql")
	environment := os.Getenv("env-mysql")
	absolutePath := filepath.Join(pwd, "conf", environment+".yaml")
	println(absolutePath)
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
	Config = config
}
