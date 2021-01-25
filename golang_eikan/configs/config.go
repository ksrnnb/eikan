package configs

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

var pathConfig string

func init() {
	pathConfig = "configs/config.ini"
}

var config map[string]string

// TODO: もうちょい整理してシングルトンにしたい
func New() map[string]string {
	if config == nil {
		cfgFile := loadConfigFile()
		config = getConfig(cfgFile)
	}

	return config
}

func loadConfigFile() *ini.File {
	cfg, err := ini.Load(pathConfig)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	return cfg
}
