package tools

import (
	"os"
	"path/filepath"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/10/30 22:51
 * @file: file.go
 * @description:
 */

func SelfPath() string {
	path, _ := filepath.Abs(os.Args[0])
	return path
}

func SelfDir() string {
	return filepath.Dir(SelfPath())
}
