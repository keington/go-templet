package database

import (
	"context"
	"errors"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/keington/go-templet/pkg/zlog"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/11 0:04
 * @file: orm_log.go
 * @description: gorm日志实现
 */

type GormLog struct {
	SugaredLog    *zap.SugaredLogger
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration
}

// LogMode 实现 gorm logger.Interface 的 LogMode 方法
func (l *GormLog) LogMode(level logger.LogLevel) logger.Interface {

	l.LogLevel = level

	return &GormLog{
		SugaredLog:    zlog.Sugared,
		SlowThreshold: l.SlowThreshold,
	}
}

func NewGormLog() *GormLog {
	return &GormLog{
		SugaredLog:    zlog.Sugared,
		SlowThreshold: 200 * time.Millisecond,
	}
}

func (l *GormLog) Info(ctx context.Context, s string, i ...interface{}) {

	if l.logger() != nil {
		l.logger().Infof(s, i...)
	}
}

func (l *GormLog) Warn(ctx context.Context, s string, i ...interface{}) {
	if l.logger() != nil {
		l.logger().Warnf(s, i...)
	}
}

func (l *GormLog) Error(ctx context.Context, s string, i ...interface{}) {
	if l.logger() != nil {
		l.logger().Errorf(s, i...)
	}
}

func (l *GormLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.logger() != nil {
		return
	}

	// 获取运行时间
	elapsed := time.Since(begin)
	// 获取 SQL 请求和返回条数
	sql, rows := fc()

	logFields := []interface{}{
		"sql", sql,
		"rowsAffected", rows,
		"elapsed", elapsed.String(),
	}

	// Gorm 错误
	if err != nil {
		// 记录未找到的错误使用 warning 等级
		if errors.Is(err, gorm.ErrRecordNotFound) {
			l.logger().Warnf("not found record: %s", logFields...)
		} else {
			// 其他错误使用 error 等级
			logFields = append(logFields, "error", err.Error())
			l.logger().Errorf("exec sql error: %s", logFields...)
		}
	}

	// 慢查询日志
	if l.SlowThreshold > 0 && elapsed > l.SlowThreshold {
		l.logger().Warnf("slow sql: %s", logFields...)
	}

	// 记录所有 SQL 请求
	l.logger().Infof("exec sql: %s", logFields...)
}

// logger 内用的辅助方法，确保 Zap 内置信息 Caller 的准确性（如 paginator/paginator.go:148）
func (l *GormLog) logger() *zap.SugaredLogger {
	if l.SugaredLog != nil {
		return l.SugaredLog
	}

	// 跳过 gorm 内置的调用
	var (
		gormPackage = filepath.Join("gorm.io", "gorm")
		// projectPackage = filepath.Join("github.com", "lark") // 项目包路径
	)

	// 减去一次封装，以及一次在 logger 初始化里添加 zap.AddCallerSkip(1)
	clone := l.SugaredLog.Desugar().WithOptions(zap.AddCallerSkip(-2)).Sugar()

	// 从当前调用栈中找到第一个不是 gorm 和项目包的调用
	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.Contains(file, gormPackage):
		// case strings.Contains(file, projectPackage):
		case strings.HasPrefix(file, "_test.go"):
		default:
			// 返回一个附带跳过 i 层调用栈的 Logger
			return clone.WithOptions(zap.AddCallerSkip(i))
		}
	}

	return zlog.Sugared
}
