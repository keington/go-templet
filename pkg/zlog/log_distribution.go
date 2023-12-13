package zlog

import (
	"fmt"
	"os"
	"time"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/9 22:34
 * @file: log_distribution.go
 * @description: 切割日志
 */

func distributionLogFile(path, name string) {
	fileOutputPath := fmt.Sprintf("%s/%s.log", path, name)
	fileInfo, err := os.Stat(fileOutputPath)
	if err == nil && fileInfo.ModTime().Before(time.Now().Add(-24*time.Hour)) {
		// 日志文件是昨天或更早的，进行分割
		err := os.Rename(fileOutputPath, fmt.Sprintf("%s/%s_%s.log", path, name, time.Now().Format("2006-01-02")))
		if err != nil {
			Errorf("error when distribution the log file: %v", err)
		}
	}
}
