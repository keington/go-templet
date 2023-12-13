package entity

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/29 22:02
 * @file: permission.go
 * @description:
 */

type Permission struct {
	BaseEntity
	PermissionId   string `gorm:"column:permission_id" json:"permissionId"`
	PermissionName string `gorm:"column:permission_name" json:"permissionName"`
}

func (Permission) TableName() string {
	return "permission"
}
