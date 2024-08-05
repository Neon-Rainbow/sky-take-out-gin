package main

import (
	"go.uber.org/zap"
	"sky-take-out-gin/internal/utils/init"
)

func main() {
	err := init.Initialize()
	if err != nil {
		zap.L().Fatal("初始化失败", zap.Error(err))
	}
}
