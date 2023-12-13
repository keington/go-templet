package zlog

import (
	"encoding/xml"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/10/27 0:18
 * @file: log.go
 * @description: 基于zap封装的日志实现
 */

type LogConfig struct {
	XmlName    xml.Name `xml:"log"`
	Mode       string   `xml:"mode"`
	Path       string   `xml:"path"`
	Name       string   `xml:"name"`
	Level      string   `xml:"level"`
	MaxSize    int      `xml:"maxSize"`
	MaxBackups int      `xml:"maxBackups"`
	MaxAge     int      `xml:"maxAge"`
	Compress   bool     `xml:"compress"`
}

var (
	Sugared *zap.SugaredLogger
)

func NewLogger(l *LogConfig) {

	distributionLogFile(l.Path, l.Name)

	var zapCfg zap.Config
	if l.Mode == "dev" {
		zapCfg = zap.NewDevelopmentConfig()
	} else {
		zapCfg = zap.NewProductionConfig()
	}

	zapCfg.Level = zap.NewAtomicLevelAt(LevelFromString(l.Level))
	zapCfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.00")
	zapCfg.EncoderConfig.EncodeLevel = levelEncoder
	zapCfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	// Custom encoder for the console that adds colors.
	consoleEncoderConfig := zapCfg.EncoderConfig
	consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // Use the predefined colored level encoder.

	// Set up the non-colored level encoder for file output.
	fileEncoderConfig := zapCfg.EncoderConfig
	fileEncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // Standard non-colored level encoder.

	fileOutputPath := fmt.Sprintf("%s/%s.log", l.Path, l.Name)
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileOutputPath,
		MaxSize:    l.MaxSize,
		MaxBackups: l.MaxBackups,
		MaxAge:     l.MaxAge,
		Compress:   l.Compress,
	})

	// Create a core for writing to file.
	fileCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(fileEncoderConfig), // Change to NewConsoleEncoder if you prefer plain text logs.
		fileWriter,
		zapCfg.Level,
	).With([]zap.Field{})

	// Create a core for writing to console.
	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(consoleEncoderConfig),
		zapcore.AddSync(os.Stdout),
		zapCfg.Level,
	).With([]zap.Field{})

	// Use zap-core.NewTee to log to both console and file with separate encoders.
	core := zapcore.NewTee(
		consoleCore,
		fileCore,
	)

	var logger *zap.Logger
	if l.Mode == "dev" {
		// 在开发模式下，标准输出并添加调用者信息
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(zapCfg.EncoderConfig),
			zapcore.NewMultiWriteSyncer(
				zapcore.AddSync(os.Stdout),
				zapcore.AddSync(&lumberjack.Logger{
					Filename:   fileOutputPath,
					MaxSize:    l.MaxSize,
					MaxBackups: l.MaxBackups,
					MaxAge:     l.MaxAge,
					Compress:   l.Compress,
				}),
			),
			zapCfg.Level,
		)
		logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	} else {
		// 在生产模式下，不添加调用者信息
		logger = zap.New(core)
	}

	Sugared = logger.Sugar()
}

func LevelFromString(levelStr string) zapcore.Level {
	switch strings.ToLower(levelStr) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	case "panic":
		return zapcore.PanicLevel
	default:
		return zapcore.DebugLevel
	}
}
