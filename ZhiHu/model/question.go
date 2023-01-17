package model

import "time"

type Question struct {
	Id         int       //问题id
	Content    string    //问题内容
	User       User      `gorm:"ForeignKey:Id;references:UserId;"` //用户
	UserId     int       //用户主键
	CreateTime time.Time //创建时间
	Answer     []Answer  `gorm:"ForeignKey:Id;references:AnswerId;"` //关于问题的回答
	AnswerId   int       //回答主键
}
