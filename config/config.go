package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var cfg *Config

type DatabaseConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Name      string `yaml:"name"`
	Charset   string `yaml:"charset"`
	ParseTime bool   `yaml:"parseTime"`
	Loc       string `yaml:"loc"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type SecretConfig struct {
	PasswordSecret string `yaml:"passwordSecret"`
	JWTSecret      string `yaml:"jwtSecret"`
}

type Config struct {
	DatabaseConfig `yaml:"database"`
	RedisConfig    `yaml:"redis"`
	ServerConfig   `yaml:"server"`
	SecretConfig   `yaml:"secret"`
}

// InitConfig 初始化配置
func InitConfig() error {
	cfg = &Config{}
	// 初始化配置
	file, err := os.Open("config.yaml")
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Panicln("关闭文件失败, err: ", err)
		}
	}(file)

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return err
	}
	return nil
}

// GetConfig 获取配置
// @Description 获取配置
// @Return *Config 配置
func GetConfig() *Config {
	if cfg == nil {
		log.Fatalf("配置文件还未初始化")
		return nil
	}
	return cfg
}
