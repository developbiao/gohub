package config

import "gohub/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			// App name
			"name": config.Env("APP_NAME", "Gohub"),

			// App environment
			"env": config.Env("APP_ENV", "production"),

			// Debug mode default is false
			"debug": config.Env("APP_DEBUG", false),

			// App server port
			"port": config.Env("APP_PORT", "3000"),
			// App key
			"key": config.Env("APP_KEY", "3336699dcf9ea060a0a6532b166da32f304af0ff"),

			// Url
			"url": config.Env("APP_URL", "http://localhost:3000"),

			// API DOMAIN
			"api_domain": config.Env("API_DOMAIN"),

			// Timezone
			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),
		}
	})
}
