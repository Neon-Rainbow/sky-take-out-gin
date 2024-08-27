package database

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"sky-take-out-gin/pkg/common/database/PostgreSQL"
	"sky-take-out-gin/pkg/common/database/Redis"
)

var dbManager *CombinedDatabase

// DatabaseInterface 定义数据库接口
type DatabaseInterface interface {
	InitDB() error
	InitRedis() error
	GetDB() *gorm.DB
	GetRedis() *redis.Client
}

type CombinedDatabase struct {
	redisClient *Redis.RedisDB
	database    *PostgreSQL.PostgresDB
}

func (d *CombinedDatabase) GetDB() *gorm.DB {
	return d.database.GetDB()
}

func (d *CombinedDatabase) GetRedis() *redis.Client {
	return d.redisClient.GetRedis()
}

// InitDB 初始化数据库
func (d *CombinedDatabase) InitDB() error {
	return d.database.InitDB()
}

func (d *CombinedDatabase) InitRedis() error {
	return d.redisClient.InitRedis()
}

func NewCombinedDatabase() error {
	postgresDB := &PostgreSQL.PostgresDB{}
	err := postgresDB.InitDB()
	if err != nil {
		return err
	}

	redisClient := &Redis.RedisDB{}
	err = redisClient.InitRedis()
	if err != nil {
		return err
	}

	db := &CombinedDatabase{
		redisClient: redisClient,
		database:    postgresDB,
	}
	dbManager = db

	fmt.Println("database", db.database)
	return nil
}

func GetDatabaseManager() *CombinedDatabase {
	return dbManager
}
