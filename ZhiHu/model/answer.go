package model

import "time"

type Answer struct {
	Id         int       //问题的ID
	Content    string    //问题的内容
	User       User      `gorm:"ForeignKey:Id;references:UserId;"` //用户
	UserId     int       //用户主键
	CreateTime time.Time //创建时间
	Remark     []Remark  `gorm:"ForeignKey:Id;references:RemarkId;"` //多条评论
	RemarkId   int       //评论的主键
}
