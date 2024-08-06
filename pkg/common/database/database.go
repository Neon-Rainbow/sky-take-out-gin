package database

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"sky-take-out-gin/pkg/common/database/PostgreSQL"
	"sky-take-out-gin/pkg/common/database/Redis"
	"sync"
)

var instance *DatabaseManager
var once sync.Once

type DatabaseManager struct {
	CombinedDatabase
	once       sync.Once
	sqlError   error
	redisError error
}

func (d *DatabaseManager) GetDB() *gorm.DB {
	return d.CombinedDatabase.GetDB()
}

func (d *DatabaseManager) GetRedis() *redis.Client {
	return d.CombinedDatabase.GetRedis()
}

// GetDatabaseManager 获取数据库管理器
func GetDatabaseManager() *DatabaseManager {
	// 保证只有一个实例
	once.Do(func() {
		instance = &DatabaseManager{}
	})
	return instance
}

// InitDB 初始化数据库
func (d *DatabaseManager) InitDB() error {
	d.once.Do(func() {
		d.sqlError = d.CombinedDatabase.InitDB()
	})
	if d.sqlError != nil {
		return d.sqlError
	}
	return nil
}

func (d *DatabaseManager) InitRedis() error {
	d.once.Do(func() {
		d.redisError = d.CombinedDatabase.InitRedis()
	})
	if d.redisError != nil {
		return d.redisError
	}
	return nil
}

// Database 定义数据库接口
type DatabaseInterface interface {
	InitDB() error
	InitRedis() error
	GetDB() *gorm.DB
	GetRedis() *redis.Client
}

type CombinedDatabase struct {
	Redis.RedisDB
	PostgreSQL.PostgresDB
}

func (d *CombinedDatabase) GetDB() *gorm.DB {
	return d.PostgresDB.GetDB()
}

func (d *CombinedDatabase) GetRedis() *redis.Client {
	return d.RedisDB.GetRedis()
}

// InitDB 初始化数据库
func (d *CombinedDatabase) InitDB() error {
	return d.PostgresDB.InitDB()
}

func (d *CombinedDatabase) InitRedis() error {
	return d.RedisDB.InitRedis()
}
