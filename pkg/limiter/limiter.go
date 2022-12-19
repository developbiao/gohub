package limiter

import (
	"github.com/gin-gonic/gin"
	limiterlib "github.com/ulule/limiter/v3"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
	"gohub/pkg/config"
	"gohub/pkg/logger"
	"gohub/pkg/redis"
	"strings"
)

// GetKeyIP get Limiter key, ip
func GetKeyIP(c *gin.Context) string {
	return c.ClientIP()
}

// GetKeyRouteWithIP get Limiter key route + IP
func GetKeyRouteWithIP(c *gin.Context) string {
	return routeToKeyString(c.FullPath() + c.ClientIP())
}

// CheckRate check request is limited
func CheckRate(c *gin.Context, key string, formatted string) (limiterlib.Context, error) {
	// New instance limiter.Rate object
	var context limiterlib.Context
	rate, err := limiterlib.NewRateFromFormatted(formatted)
	if err != nil {
		logger.LogIf(err)
		return context, err
	}

	// Initialization store, here using redis.Redis object
	store, err := sredis.NewStoreWithOptions(redis.Redis.Client, limiterlib.StoreOptions{
		Prefix: config.GetString("app.name") + ":limiter",
	})
	if err != nil {
		logger.LogIf(err)
		return context, err
	}

	// Initialization limiter.Rate object
	limiterObj := limiterlib.New(store, rate)

	// Get limiter result
	if c.GetBool("limiter-once") {
		// Peek() get result
		return limiterObj.Peek(c, key)
	} else {
		// limiter route combination increment once
		c.Set("limiter-once", true)

		// Get limit count
		return limiterObj.Get(c, key)
	}
}

// routeToKeyString helper function, subtitle URL "/" to "-"
func routeToKeyString(routeName string) string {
	routeName = strings.ReplaceAll(routeName, "/", "-")
	routeName = strings.ReplaceAll(routeName, ":", "_")
	return routeName
}
