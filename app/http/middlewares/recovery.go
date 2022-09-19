package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gohub/pkg/logger"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

// Recovery middleware
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Get request information
				httpRequest, _ := httputil.DumpRequest(c.Request, true)
				// Connection broken is client normal behavior, not necessary record stack information
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						errStr := strings.ToLower(se.Error())
						if strings.Contains(errStr, "broken pipe") ||
							strings.Contains(errStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				// If connection broken
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					c.Error(err.(error))
					c.Abort()
					// Connection is break, can not write status code
					return
				}

				// Recover stack information
				logger.Error("recovery from panic",
					zap.Time("time", time.Now()),               // time
					zap.Any("error", err),                      // error message
					zap.String("request", string(httpRequest)), // request
					zap.Stack("stacktrace"),                    // stack trace
				)

				// Return 500 status code
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Server internal server error",
				})
			}
		}()
		c.Next()
	}
}