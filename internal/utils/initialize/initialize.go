package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sky-take-out-gin/pkg/common/config"
	"sky-take-out-gin/pkg/common/database/MySQL"
	"sky-take-out-gin/pkg/common/database/Redis"
	"sky-take-out-gin/pkg/common/logger"
	"sky-take-out-gin/route"
)

// Initialize 用于初始化项目
func Initialize() error {
	// 初始化配置
	if err := config.InitConfig(); err != nil {
		return fmt.Errorf("初始化配置失败: %w", err)
	}

	// 设置Gin模式
	gin.SetMode(config.GetConfig().ServerConfig.Mode)

	// 初始化数据库
	if err := MySQL.InitDB(); err != nil {
		return fmt.Errorf("初始化数据库失败: %w", err)
	}

	// 初始化Redis
	if err := Redis.InitRedis(); err != nil {
		return fmt.Errorf("初始化Redis失败: %w", err)
	}

	// 设置全局日志
	if err := logger.SetupGlobalLogger(); err != nil {
		return fmt.Errorf("初始化全局日志失败: %w", err)
	}

	// 初始化路由
	if err := route.SetupHTTPRoute(); err != nil {
		return fmt.Errorf("初始化路由失败: %w", err)
	}

	return nil
}
