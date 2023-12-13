package osx

import "os"

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/11 22:15
 * @file: osx.go
 * @description:
 */

// GetEnv 返回环境变量的值，或返回提供的后备值
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
