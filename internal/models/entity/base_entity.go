package entity

import (
	"time"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/15 23:03
 * @file: base_entity.go
 * @description:
 */

type BaseEntity struct {
	Id       int64     `gorm:"column:id" json:"id"`
	CreateAt time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateAt time.Time `gorm:"column:update_time" json:"update_time"`
	DeleteAt time.Time `gorm:"column:delete_time" json:"delete_time"`
}
