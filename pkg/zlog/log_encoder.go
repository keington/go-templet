package zlog

import (
	"fmt"
	"go.uber.org/zap/zapcore"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/9 22:39
 * @file: log_encoder.go
 * @description: 自定义日志格式
 */

// levelEncoder 自定义日志级别格式
func levelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var level string
	switch l {
	case zapcore.DebugLevel:
		level = "[DEBUG]"
	case zapcore.InfoLevel:
		level = "[INFO]"
	case zapcore.WarnLevel:
		level = "[WARN]"
	case zapcore.ErrorLevel:
		level = "[ERROR]"
	case zapcore.PanicLevel:
		level = "[PANIC]"
	case zapcore.FatalLevel:
		level = "[FATAL]"
	default:
		level = fmt.Sprintf("[LEVEL(%d)]", l)
	}
	enc.AppendString(level)
}
