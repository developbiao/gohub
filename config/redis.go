package config

import "gohub/pkg/config"

func init() {
	config.Add("redis", func() map[string]interface{} {
		return map[string]interface{}{
			"host":     config.Env("REDIS_HOST", "127.0.0.1"),
			"port":     config.Env("REDIS_POST", "6379"),
			"password": config.Env("REDIS_PASSWORD", ""),
			// Business using with (session, sms, etc)
			"database": config.Env("REDIS_MAIN_DB", 1),
		}
	})
}
