package main

import (
	"fmt"
	"reflect"
	"strings"

	"log"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	DB struct {
		User string `mapstructure:"user"`
	} `mapstructure:"db"`
}

func BindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			BindEnvs(v.Interface(), append(parts, tv)...)
		default:
			viper.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}
func main() {

	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		switch err.(type) {
		default:
			panic(fmt.Errorf("Fatal error loading config file: %s \n", err))
		case viper.ConfigFileNotFoundError:
			log.Println("No config file found. Using defaults and environment variables")
		}
	}

	serverConfig := ServerConfig{}
	BindEnvs(serverConfig)
	viper.Unmarshal(&serverConfig)

	fmt.Println("DB USER: " + serverConfig.DB.User)
}
