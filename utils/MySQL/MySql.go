package MySQL

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sky-take-out-gin/config"
	model "sky-take-out-gin/model/sql"
	"time"
)

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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=%v&loc=%v",
		cfg.DatabaseConfig.Username,
		cfg.DatabaseConfig.Password,
		cfg.DatabaseConfig.Host,
		cfg.DatabaseConfig.Port,
		cfg.DatabaseConfig.Name,
		cfg.DatabaseConfig.ParseTime,
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

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
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
