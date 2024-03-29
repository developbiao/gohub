package config

import "gohub/pkg/config"

func init() {
	config.Add("redis", func() map[string]interface{} {
		return map[string]interface{}{
			"host":     config.Env("REDIS_HOST", "127.0.0.1"),
			"port":     config.Env("REDIS_POST", "6379"),
			"username": config.Env("REDIS_USERNAME", ""),
			"password": config.Env("REDIS_PASSWORD", ""),
			// Business using with (session, sms, etc)
			"database": config.Env("REDIS_MAIN_DB", 1),
			// Cache using 0, Clearing the cache should be not affect the business
			"database_cache": config.Env("REDIS_CACHE_DB", 0),
		}
	})
}
