package logger

import (
	"fmt"
	"go.uber.org/zap"
	"sky-take-out-gin/pkg/common/config"
	"time"
)

// SetupGlobalLogger 用于初始化全局日志
func SetupGlobalLogger() error {
	currentTime := time.Now()
	// 生成日志文件名
	logFileName := fmt.Sprintf("./%s/%s.log",
		config.GetConfig().ServerConfig.LogFilePath,
		currentTime.Format("2006-01-02 15:04:05"))
	logger, err := NewLogger(logFileName)
	if err != nil {
		return err
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			logger.Error("Failed to sync logger", zap.Error(err))
		}
	}(logger)

	zap.ReplaceGlobals(logger)
	return nil
}
