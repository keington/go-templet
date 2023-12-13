package view

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/15 23:33
 * @file: user.go
 * @description:
 */

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
