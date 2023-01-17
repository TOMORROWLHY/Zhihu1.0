package service

import (
	"Little_Project/ZhiHu/dao/mysql"
	"Little_Project/ZhiHu/model"
)

// 创建用户
func CreatUser(u model.User) error {
	err := mysql.InsertUser(u)
	return err
}

// 通过用户名查找用户
func SearchUserByUsername(name string) (u model.User, err error) {
	u, err = mysql.SearchUserByUsername(name)
	return u, err
}

// 通过用户名查找用户ID
func SearchUIDByUsername(name string) (id int, err error) {
	u, err := mysql.SearchUserByUsername(name)
	id = u.Id
	return id, err
}

// 修改密码
func ChangePass(username, newpass string) error {
	err := mysql.ChangePass(username, newpass)
	return err
}

// 判断用户名是否重复
func CheckExist(username string) (flag bool) {
	u, _ := mysql.SearchUserByUsername(username)
	flag = true
	if username != u.Username {
		flag = false
		return flag
	}
	return flag
}

// 修改用户的用户名
func ChangeName(username string, newname string) (err error) {
	err = mysql.ChangeName(username, newname)
	return err
}
