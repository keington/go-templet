package zlog

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/8 0:24
 * @file: log_rewrite.go
 * @description: 重写zap日志方法
 */

func Debug(args ...interface{}) {
	Sugared.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	Sugared.Debugf(template, args...)
}

func Info(args ...interface{}) {
	Sugared.Info(args...)
}

func Infof(template string, args ...interface{}) {
	Sugared.Infof(template, args...)
}

func Warn(args ...interface{}) {
	Sugared.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	Sugared.Warnf(template, args...)
}

func Error(args ...interface{}) {
	Sugared.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	Sugared.Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	Sugared.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	Sugared.Fatalf(template, args...)
}

func Panic(args ...interface{}) {
	Sugared.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	Sugared.Panicf(template, args...)
}
