package configs

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

var cfgFile *ini.File
var pathConfig string
var config map[string]string

func init() {
	if isInConfigsDirectory() {
		// test envionment
		pathConfig = "config.ini"
	} else {
		// go run main.go
		pathConfig = "configs/config.ini"
	}
}

func isInConfigsDirectory() bool {
	workingDirectory, _ := os.Getwd()
	return strings.Contains(workingDirectory, "/configs")
}

// New : pathConfigのファイルを読み込んで、mapを返す
func New() map[string]string {
	if config == nil {
		cfgFile = loadConfigFile()
		config = getConfig()
	}

	return config
}

// ファイルの読み込み、エラーが出たら終了。
func loadConfigFile() *ini.File {
	cfg, err := ini.Load(pathConfig)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	return cfg
}
