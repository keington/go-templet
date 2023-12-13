package database

import (
	"fmt"
	"strings"
	"time"

	"github.com/gookit/slog"
	"github.com/keington/go-templet/pkg/zlog"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/10 23:48
 * @file: instance.go
 * @description: 数据库连接
 */

var DB *gorm.DB

// Config DataBaseConfig GORM
type Config struct {
	Type        string
	DSN         string
	MaxOpenCon  int
	MaxIdleCon  int
	MaxLifetime int
	MaxIdleTime int
	LogEnabled  bool
}

func NewDatabase(c *Config) (*gorm.DB, error) {

	var dialect gorm.Dialector

	switch strings.ToLower(c.Type) {
	case "mysql":
		dialect = mysql.Open(c.DSN)
	case "postgres":
		dialect = postgres.Open(c.DSN)
	default:
		return nil, fmt.Errorf("db type(%s) not support", c.Type)
	}

	gormCfg := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger: logger.Default.LogMode(logger.Info),
	}

	if c.LogEnabled {
		gormCfg.Logger = NewGormLog()
		zlog.Debugf("SQL statement is enabled")
	}

	db, err := gorm.Open(dialect, gormCfg)

	if err != nil {
		slog.Panicf("failed to connect to MySQL: %v", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		slog.Errorf("failed to connect to database: %v", err)
		return nil, err
	}

	// 设置连接池大小
	sqlDB.SetMaxOpenConns(c.MaxOpenCon)
	sqlDB.SetMaxIdleConns(c.MaxIdleCon)
	sqlDB.SetConnMaxLifetime(time.Duration(c.MaxLifetime) * time.Second)
	sqlDB.SetConnMaxIdleTime(time.Duration(c.MaxIdleTime) * time.Second)

	if err != nil {
		return nil, err
	}

	DB = db

	return db, err
}
