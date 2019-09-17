package models

import (
	. "../../app/database"
)

type User struct {
	Id int `json:"id" from "id"`
	UserName string `json:"username from "username"`
	PassWd string `json:"passwd" from "passwd"`
}

//select
func (u *User) GetUserInfo() (users []User, err error) {
	//创建一个切片
	users = make([]User, 0)
	rows, err := SqlDB.Query("select id, username, passwd from user")
	if err != nil {
		return
	}
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.UserName, &user.PassWd)
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
