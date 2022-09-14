package app

import "gohub/pkg/config"

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}

func IsTesting() bool {
	configEnv := config.Get("app.env")
	return configEnv == "testing" || configEnv == "test"
}
