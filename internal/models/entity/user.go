package entity

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/keington/go-templet/pkg/str"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/15 22:58
 * @file: user.go
 * @description:
 */

type User struct {
	BaseEntity
	UserId    string `gorm:"column:user_id" json:"userId"`
	UserName  string `gorm:"column:user_name" json:"userName"`
	NickName  string `gorm:"column:nick_name" json:"NickName"`
	Password  string `gorm:"column:password" json:"password"`
	RoleId    string `gorm:"column:role_id" json:"roleId"`
	Phone     string `gorm:"column:phone" json:"phone"`
	Email     string `gorm:"column:email" json:"email"`
	UserState int    `gorm:"column:user_state" json:"userState"`
}

func (User) TableName() string {
	return "user"
}

func (u User) Verify() error {
	u.UserName = strings.TrimSpace(u.UserName)

	if u.UserName == "" {
		return errors.New("username is blank")
	}

	if str.Dangerous(u.UserName) {
		return errors.New("username has invalid characters")
	}

	if str.Dangerous(u.NickName) {
		return errors.New("nickname has invalid characters")
	}

	if u.Phone != "" && !str.IsPhone(u.Phone) {
		return errors.New("phone invalid")
	}

	if u.Email != "" && !str.IsMail(u.Email) {
		return errors.New("email invalid")
	}

	return nil
}

func GetUserById(id int) (user User, err error) {
	err = DB().Where("id = ?", id).First(&user).Error
	return
}

func GetUserByUserName(userId string) (user User, err error) {
	err = DB().WithContext(context.TODO()).Select("user_name", "user_state").Where("user_id = ?", userId).First(&user).Error
	return
}

func GetUserByPasswd(userName string) (User, error) {

	var user User
	err := DB().Table("users").Select("password").Where("user_name = ?", "admin").First(&user).Error
	if err != nil {
		_ = fmt.Errorf(err.Error())
	}
	return user, err
}
