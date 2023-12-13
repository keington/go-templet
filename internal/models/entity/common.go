package entity

import (
	"github.com/keington/go-templet/pkg/database"
	"gorm.io/gorm"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/15 23:31
 * @file: common.go
 * @description:
 */

func DB() *gorm.DB {
	return database.DB
}

func Count(tx *gorm.DB) (int64, error) {
	var cnt int64
	err := tx.Count(&cnt).Error
	return cnt, err
}

func Exists(tx *gorm.DB) (bool, error) {
	num, err := Count(tx)
	return num > 0, err
}
