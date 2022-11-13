package config

import "gohub/pkg/config"

func init() {
	config.Add("jwt", func() map[string]interface{} {
		return map[string]interface{}{
			// Using config.GetString("app.key")
			// "signing_key":

			// Expire time unit is minute maximum is 2 hours
			"expire_time": config.Env("JWT_EXPIRE_TIME", 120),

			// Allow refresh time unit is minute 86400 is 2 hours
			"max_refresh_time": config.Env("JWT_MAX_REFRESH_TIME", 86400),

			// debug mode for test
			"debug_expire_time": 86400,
		}
	})
}
