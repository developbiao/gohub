package middlewares

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"gohub/pkg/helpers"
	"gohub/pkg/logger"
	"io"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write response writer
func (r *responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// Logger record request log
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get response content
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		// Get request data
		var requestBody []byte
		if c.Request.Body != nil {
			// c.Request.Body is buffer only can read once
			requestBody, _ = io.ReadAll(c.Request.Body)
			// After reading, reassign
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// Set start time
		start := time.Now()
		c.Next()

		// Starting record log logic
		cost := time.Since(start)
		responseStatus := c.Writer.Status()

		logFields := []zap.Field{
			zap.Int("status", responseStatus),
			zap.String("request", c.Request.Method+" "+c.Request.URL.String()),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", helpers.MicrosecondsStr(cost)),
		}

		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
			// Get request content
			logFields = append(logFields, zap.String("Request Body", string(requestBody)))

			// Get response content
			logFields = append(logFields, zap.String("Response Body", w.body.String()))
		}

		if responseStatus > 400 && responseStatus <= 499 {
			// Exclude StatusBadRequest, client warning 403, 404...
			logger.Warn("HTTP Warning"+cast.ToString(responseStatus), logFields...)
		} else if responseStatus >= 500 && responseStatus <= 599 {
			// Exclude inner error, record error like server error
			logger.Error("HTTP Error" + cast.ToString(responseStatus))
		} else {
			logger.Debug("HTTP Access Log", logFields...)
		}
	}

}
