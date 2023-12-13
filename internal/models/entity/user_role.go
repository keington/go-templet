package entity

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/29 22:01
 * @file: user_role.go
 * @description:
 */

type UserRole struct {
	Id     int64  `gorm:"column:id" json:"id"`
	UserId string `gorm:"column:user_id" json:"userId"`
	RoleId string `gorm:"column:role_id" json:"roleId"`
}

func (UserRole) TableName() string {
	return "user_role"
}
