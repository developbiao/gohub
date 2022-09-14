package bootstrap

import (
	"gohub/pkg/config"
	"gohub/pkg/logger"
)

// SetupLogger Initialization
func SetupLogger() {
	logger.InitLogger(
		config.GetString("log.filename"),
		config.GetInt("log.max_size"),
		config.GetInt("log_max_backup"),
		config.GetInt("log_max_age"),
		config.GetBool("log.compress"),
		config.GetString("log.type"),
		config.GetString("log.level"),
	)
}
