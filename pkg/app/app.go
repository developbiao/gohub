package app

import (
	"gohub/pkg/config"
	"time"
)

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

// TimenowInTimezone get current time from config timezone
func TimenowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.GetString("app.timezone"))
	return time.Now().In(chinaTimezone)
}

// URL parameter path contact URL
func URL(path string) string {
	return config.Get("app.url") + path
}

// V1URL parameter v1 identify contact URL
func V1URL(path string) string {
	return URL("/v1/" + path)
}
