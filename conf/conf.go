package conf

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
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
	environment := os.Getenv("environment")
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
	Config = config
}
