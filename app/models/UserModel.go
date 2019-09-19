package models

import (
	. "../../app/database"
)

type User struct {
	Id       int    `json:"id" from "id"`
	UserName string `json:"username" from "username"`
	PassWd   string `json:"passwd" from "passwd"`
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
		_ = rows.Scan(&user.Id, &user.UserName, &user.PassWd)
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

//select by id
func (u *User) GetUserInfoById() (users []User, err error) {
	_ = SqlDB.QueryRow("select id, username, passwd from user where id = ?", u.Id).Scan(&u.Id, &u.UserName, &u.PassWd)
	return
}

//add a user
func (u *User) AddUser() (lastId int64, err error) {
	rs, err := SqlDB.Exec("INSERT INTO user(id,username,passwd) VALUES (0,?,?)", u.UserName, u.PassWd)
	if err != nil {
		return
	}
	lastId, err = rs.LastInsertId()
	return lastId, err
}

//update a user
func (u *User) UpdateUser() (effect_num int64, err error) {
	rs, err := SqlDB.Exec("UPDATE user set username = ? , passwd = ? where id = ?", u.UserName, u.PassWd, u.Id)
	if err != nil {
		return
	}
	effect_num, err = rs.RowsAffected()
	return effect_num, err
}

//del a user
func (u *User) DelUser() (effort_num int64, err error) {
	rs, err := SqlDB.Exec("DELETE FROM user where id= ?", u.Id)
	if err != nil {
		return
	}
	effort_num, err = rs.RowsAffected()
	return
}
