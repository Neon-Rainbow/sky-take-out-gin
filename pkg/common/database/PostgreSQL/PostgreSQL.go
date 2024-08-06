package PostgreSQL

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	model "sky-take-out-gin/model/sql"
	"sky-take-out-gin/pkg/common/config"
)

type PostgresDB struct {
	*gorm.DB
}

func (p PostgresDB) InitDB() error {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		cfg.DatabaseConfig.Host,
		cfg.DatabaseConfig.Username,
		cfg.DatabaseConfig.Password,
		cfg.DatabaseConfig.Name,
		cfg.DatabaseConfig.Port,
		cfg.DatabaseConfig.Loc,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 将日志输出到终端
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			// LogLevel:      logger.Info, // 日志级别
			Colorful: true, // 彩色打印
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return err
	}

	// 迁移数据库
	err = db.AutoMigrate(
		&model.AddressBook{},
		&model.Category{},
		&model.Dish{},
		&model.DishFlavor{},
		&model.Employee{},
		&model.Order{},
		&model.OrderDetail{},
		&model.Setmeal{},
		&model.SetmealDish{},
		&model.ShoppingCart{},
		&model.User{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
		return err
	}
	p.DB = db
	return nil
}

func (p PostgresDB) GetDB() *gorm.DB {
	if p.DB == nil {
		zap.L().Fatal("数据库还未初始化")
		return nil
	}
	return p.DB
}

var db *gorm.DB

// GetDB 获取数据库连接
// @Description 获取数据库连接,在数据库连接未初始化时会报错
// @Return *gorm.DB 数据库连接
func GetDB() *gorm.DB {
	if db == nil {
		log.Fatalf("数据库还未初始化")
		return nil
	}
	return db
}

// InitDB 初始化数据库连接
func InitDB() (err error) {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		cfg.DatabaseConfig.Host,
		cfg.DatabaseConfig.Username,
		cfg.DatabaseConfig.Password,
		cfg.DatabaseConfig.Name,
		cfg.DatabaseConfig.Port,
		cfg.DatabaseConfig.Loc,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 将日志输出到终端
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			// LogLevel:      logger.Info, // 日志级别
			Colorful: true, // 彩色打印
		},
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return err
	}

	// 迁移数据库
	err = db.AutoMigrate(
		&model.AddressBook{},
		&model.Category{},
		&model.Dish{},
		&model.DishFlavor{},
		&model.Employee{},
		&model.Order{},
		&model.OrderDetail{},
		&model.Setmeal{},
		&model.SetmealDish{},
		&model.ShoppingCart{},
		&model.User{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
		return err
	}

	return nil
}
