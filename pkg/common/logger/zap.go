package logger

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"sky-take-out-gin/pkg/common/config"
)

// NewLogger 创建并初始化zap日志库
func NewLogger(logFilePath string) (*zap.Logger, error) {
	logDir := filepath.Dir(logFilePath)
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.MkdirAll(logDir, 0755)
		if err != nil {
			return nil, fmt.Errorf("failed to create log directory: %w", err)
		}
	}

	zapConfig := zap.NewProductionConfig()

	cfg := config.GetConfig()
	level := zapcore.InfoLevel
	if err := level.UnmarshalText([]byte(cfg.Log.Level)); err != nil {
		return nil, fmt.Errorf("invalid log level: %w", err)
	}

	zapConfig.Level = zap.NewAtomicLevelAt(level)
	zapConfig.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	zapConfig.OutputPaths = []string{
		logFilePath,
		"stdout",
	}

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}
	return logger, nil
}

// GinLogger 接收gin框架默认的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// 由于c.Request.Body只允许读一次,读完之后内容会被删除,因此在这里我们读取Body中的内容,然后存储在变量中
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.Any("head", c.Request.Header),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("body", string(bodyBytes)),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					c.Error(err.(error))
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
