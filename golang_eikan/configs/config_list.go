package configs

// Configを追加する時はここに追加していく。
func getConfig() map[string]string {
	config := map[string]string{
		"APP_ENV":     getEnvionment(),
		"DB.ENDPOINT": getDbEndpoint(),
		"DB.USERNAME": cfgFile.Section("DB").Key("USERNAME").String(),
		"DB.PASSWORD": cfgFile.Section("DB").Key("PASSWORD").String(),
	}

	return config
}

// アプリケーション環境の取得
func getEnvionment() string {
	return cfgFile.Section("").Key("APP_ENV").String()
}

// Databaseのエンドポイントを取得
func getDbEndpoint() string {
	var dbEnd string
	switch getEnvionment() {
	case "development":
		dbEnd = cfgFile.Section("DB").Key("DEV_ENDPOINT").String()
	case "production":
		dbEnd = cfgFile.Section("DB").Key("PROD_ENDPOINT").String()
	}
	return dbEnd
}
