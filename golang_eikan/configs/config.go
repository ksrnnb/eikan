package configs

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

var cfgFile *ini.File
var pathConfig string
var config map[string]string

func init() {
	pathConfig = "configs/config.ini"
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
