package config

import "gohub/pkg/config"

func init() {
	config.Add("log", func() map[string]interface{} {
		return map[string]interface{}{
			// Log level options
			// ["debug, info, warn, error"]
			"level": config.Env("LOG_LEVEL", "debug"),
			// Log type options
			// "single" single file
			// "daily" daily log file
			"type": config.Env("LOG_TYPE", "single"),
			//* --------------- Rotate log config --------------- */
			// Log file path
			"filename": config.Env("LOG_NAME", "storage/logs/logs.log"),
			// File max size
			"max_size": config.Env("LOG_MAX_SIZE", 64),
			// Max save file number, max age expired also delete
			"max_backup": config.Env("LOG_MAX_BACKUP", 5),
			// File how many days keep
			"max_age": config.Env("LOG_MAX_AGE", 30),
			// Compress file
			"compress": config.Env("LOG_COMPRESS", false),
		}
	})
}
