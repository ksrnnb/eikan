package configs

import (
	"gopkg.in/ini.v1"
)

//TODO: もうちょい見やすく分割？
func getConfig(cfgFile *ini.File) map[string]string {
	appEnv := cfgFile.Section("").Key("APP_ENV").String()

	var dbEnd string
	switch appEnv {
	case "development":
		dbEnd = cfgFile.Section("DB").Key("DEV_ENDPOINT").String()
	case "production":
		dbEnd = cfgFile.Section("DB").Key("PROD_ENDPOINT").String()
	}

	config := map[string]string{
		"APP_ENV":      appEnv,
		"DB.ENDPOINT":  dbEnd,
		"DB.USER_NAME": cfgFile.Section("DB").Key("USER_NAME").String(),
		"DB.PASSWORD":  cfgFile.Section("DB").Key("PASSWORD").String(),
	}

	return config
}
