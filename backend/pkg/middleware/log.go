package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

var hostname, _ = os.Hostname()

func LogMiddleware(logger *logrus.Logger, pathPrefix ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		logged := len(pathPrefix) == 0

		for _, prefix := range pathPrefix {
			if strings.HasPrefix(path, prefix) {
				logged = true
				break
			}
		}

		if !logged {
			return
		}

		start := time.Now()

		defer func() {
			latency := time.Since(start)
			statusCode := c.Writer.Status()
			clientIP := c.ClientIP()
			clientUserAgent := c.Request.UserAgent()

			entry := logger.WithFields(logrus.Fields{
				"hostname":   hostname,
				"path":       path,
				"method":     c.Request.Method,
				"statusCode": statusCode,
				"clientIP":   clientIP,
				"userAgent":  clientUserAgent,
			})

			if len(c.Errors) > 0 {
				entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
			} else {
				msg := fmt.Sprintf("[%s %s] %d %v", c.Request.Method, c.Request.URL, statusCode, latency)
				if statusCode >= http.StatusInternalServerError { //nolint: gocritic
					entry.Error(msg)
				} else if statusCode >= http.StatusBadRequest {
					entry.Warn(msg)
				} else {
					entry.Info(msg)
				}
			}
		}()

		c.Next()
	}
}

func LogTaskMiddleware(logger *logrus.Logger) asynq.MiddlewareFunc {
	return func(handler asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
			start := time.Now()
			logger.Infof("Task: %s, Payload: %s", t.Type(), t.Payload())
			if err := handler.ProcessTask(ctx, t); err != nil {
				logger.Errorf("Task: %s, Error: %s", t.Type(), err.Error())
				return err
			}
			logger.Infof("Finished task: %s: Elapsed Time = %v", t.Type(), time.Since(start))

			return nil
		})
	}
}
