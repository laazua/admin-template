package db

import (
	"database/sql"
	"log/slog"

	"adtmp/internal/domain/entities"
	"adtmp/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	sqlDB  *sql.DB
	gormDB *gorm.DB
)

func Get() (*gorm.DB, error) {
	if gormDB != nil {
		return gormDB, nil
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.Get().DBUrl,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}
	// 设置连接池参数
	sqlDB.SetMaxIdleConns(config.Get().DBIdleConn)    // 最大空闲连接数
	sqlDB.SetMaxOpenConns(config.Get().DBOpenConn)    // 最大打开连接数
	sqlDB.SetConnMaxLifetime(config.Get().DBLifetime) // 连接的最大存活时间
	sqlDB.SetConnMaxIdleTime(config.Get().DBIdleTime) // 最大空闲连接时间
	slog.Info("数据库初始化成功...")
	// Ping 数据库检查连接是否正常
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	if err := gormDB.AutoMigrate(&entities.User{}, &entities.Role{}, &entities.Route{}, &entities.UserRole{}, &entities.RoleRoute{}); err != nil {
		slog.Error("迁移表失败", slog.String("Error", err.Error()))
		return nil, err
	}

	return gormDB, nil
}

func Close() {
	if sqlDB != nil {
		sqlDB.Close()
	}
}
