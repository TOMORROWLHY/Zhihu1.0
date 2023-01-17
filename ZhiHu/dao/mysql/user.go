package mysql

import (
	"Little_Project/ZhiHu/model"
	"go.uber.org/zap"
)

// 插入用户
func InsertUser(u model.User) (err error) {
	tx := DB.Create(&u)
	err = tx.Error
	if err != nil {
		zap.L().Error("create user err: ", zap.Error(err))
		return
	}
	return err
}

// 通过用户名查找用户
func SearchUserByUsername(name string) (u model.User, err error) {
	tx := DB.Where("username=?", name).Find(&u)
	err = tx.Error
	if err != nil {
		zap.L().Error("search user error: ", zap.Error(err))
		panic(err)
	}
	return u, err
}

// 修改用户密码
func ChangePass(username string, newpass string) (err error) {
	var u model.User
	tx := DB.Model(&u).Where("username=?", username).Update("password", newpass)
	err = tx.Error
	if err != nil {
		zap.L().Error("change password err: ", zap.Error(err))
		return
	}
	return err
}

// 修改用户的用户名
func ChangeName(username string, newname string) (err error) {
	var u model.User
	tx := DB.Model(&u).Where("username=?", username).Update("username", newname)
	err = tx.Error
	if err != nil {
		zap.L().Error("change name err: ", zap.Error(err))
		return
	}
	return err
}
