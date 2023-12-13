package entity

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/29 21:59
 * @file: role.go
 * @description:
 */

type Role struct {
	BaseEntity
	RoleId    string `gorm:"column:role_id" json:"roleId"`
	RoleName  string `gorm:"column:role_name" json:"roleName"`
	RoleState int    `gorm:"column:role_state" json:"roleState"`
}

func (Role) TableName() string {
	return "role"
}
