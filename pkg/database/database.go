package database

import (
	"dididaren/internal/model"
	"dididaren/pkg/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %v", err)
	}

	// 自动迁移数据库表结构
	err = db.AutoMigrate(
		&model.User{},
		&model.Staff{},
		&model.Emergency{},
		&model.EmergencyContact{},
		&model.HandlingRecord{},
		&model.DangerZone{},
		&model.Rating{},
		&model.SystemConfig{},
	)
	if err != nil {
		return nil, fmt.Errorf("迁移数据库失败: %v", err)
	}

	return db, nil
}
