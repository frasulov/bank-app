package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"runtime"
)

var Configuration Configurations
var ProfileConfiguration ProfileConfigurations
var ProjectPath string

var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	viper.AddConfigPath(basePath)
	viper.SetConfigType("yml")
	viper.SetConfigName("config-profile.yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf(fmt.Sprintf("Error reading config file: %v", err.Error()))
	}

	err := viper.Unmarshal(&ProfileConfiguration)

	if err != nil {
		log.Println(err.Error())
	}

	if ProfileConfiguration.Profile.Active != "" {
		activeProfile := ProfileConfiguration.Profile.Active
		configYaml := fmt.Sprintf("config-%v.yml", activeProfile)
		viper.SetConfigName(configYaml)
		if err := viper.ReadInConfig(); err != nil {
			log.Println(err.Error())
		}
		err = viper.Unmarshal(&Configuration)
		if err != nil {
			log.Println(err.Error())
		}

	}
}
