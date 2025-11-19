package db

import (
	"database/sql"
	"log"
	"log/slog"

	"adtmp/pkg/config"
	"adtmp/pkg/internal/domain/entities"

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
	// 迁移表
	tables := []any{&entities.User{}, &entities.Role{}, &entities.Route{}, &entities.UserRole{}, &entities.RoleRoute{}}
	for _, table := range tables {
		// log.Printf("MIGRATING table: %T", table)
		if !gormDB.Migrator().HasTable(table) {
			log.Printf("WARN 表 %T 不存在, 创建中 ...", table)
			if err := gormDB.Migrator().CreateTable(table); err != nil {
				log.Printf("ERR 创建表 %T 失败: %s", table, err.Error())
				return nil, err
			}
			log.Printf("INFO 表 %T 创建成功", table)
		}
	}
	// if err := gormDB.AutoMigrate(&entities.User{}, &entities.Role{}, &entities.Route{}, &entities.UserRole{}, &entities.RoleRoute{}); err != nil {
	// 	// slog.Error("迁移表失败", slog.String("Error", err.Error()))
	// 	log.Printf("ERR 迁移表失败: %s", err.Error())
	// 	return nil, err
	// }

	return gormDB, nil
}

func Close() {
	if sqlDB != nil {
		sqlDB.Close()
	}
}
