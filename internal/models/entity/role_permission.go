package entity

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/29 22:03
 * @file: role_permission.go
 * @description:
 */

type RolePermission struct {
	Id           int64  `gorm:"column:id" json:"id"`
	RoleId       string `gorm:"column:role_id" json:"roleId"`
	PermissionId string `gorm:"column:permission_id" json:"permissionId"`
}

func (RolePermission) TableName() string {
	return "role_permission"
}
