package model

type User struct {
	Id           int    //用户ID
	Username     string //用户名
	Password     string //密码
	Level        int    `gorm:"default:1"` //用户等级
	Location     string //用户住址
	Introduction string //简介
	Gender       string //性别
	Work         string //所在行业
	Workexp      string //职业经历
}
