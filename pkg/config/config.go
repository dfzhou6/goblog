package config

import (
	"github.com/dfzhou6/goblog/pkg/logger"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var Viper *viper.Viper

type StrMap map[string]interface{}

func init() {
	Viper = viper.New()
	Viper.SetConfigName(".env")
	Viper.SetConfigType("env")
	Viper.AddConfigPath(".")

	err := Viper.ReadInConfig()
	logger.LogError(err)

	Viper.SetEnvPrefix("appenv")
	Viper.AutomaticEnv()
}

func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return Get(envName, defaultValue[0])
	}
	return Get(envName)
}

func Get(path string, defaultValue ...interface{}) interface{} {
	if !Viper.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return Viper.Get(path)
}

func Add(name string, configuration map[string]interface{}) {
	viper.Set(name, configuration)
}

func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(Get(path, defaultValue...))
}

func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(Get(path, defaultValue...))
}

func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(Get(path, defaultValue...))
}
