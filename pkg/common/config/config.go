package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var cfg *Config

type DatabaseConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	Name      string `mapstructure:"name"`
	Charset   string `mapstructure:"charset"`
	ParseTime bool   `mapstructure:"parseTime"`
	Loc       string `mapstructure:"loc"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type ServerConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	LogFilePath string `mapstructure:"logFilePath"`
	Mode        string `mapstructure:"mode"`
}

type SecretConfig struct {
	PasswordSecret string `mapstructure:"passwordSecret"`
	JWTSecret      string `mapstructure:"jwtSecret"`
}

type Log struct {
	Level string `mapstructure:"level"`
}

type Config struct {
	DatabaseConfig `mapstructure:"database"`
	RedisConfig    `mapstructure:"redis"`
	ServerConfig   `mapstructure:"server"`
	SecretConfig   `mapstructure:"secret"`
	Log            `mapstructure:"log"`
}

// InitConfig 初始化配置
func InitConfig() error {
	// 设置默认值
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 3306)
	viper.SetDefault("database.username", "root")
	viper.SetDefault("database.password", "")
	viper.SetDefault("database.name", "test")
	viper.SetDefault("database.charset", "utf8mb4")
	viper.SetDefault("database.parseTime", true)
	viper.SetDefault("database.loc", "Local")

	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", 6379)
	viper.SetDefault("redis.username", "")
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.db", 0)

	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.logFilePath", "./logs")
	viper.SetDefault("server.mode", "release")

	viper.SetDefault("secret.passwordSecret", "default_secret")
	viper.SetDefault("secret.jwtSecret", "default_jwt_secret")

	viper.SetDefault("log.level", "info")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	cfg = &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("解析配置文件失败: %w", err)
	}

	return nil
}

// GetConfig 获取配置
// @Description 获取配置
// @Return *Config 配置
func GetConfig() *Config {
	if cfg == nil {
		zap.L().Fatal("配置未初始化")
		return nil
	}
	return cfg
}
