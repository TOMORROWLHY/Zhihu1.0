package model

import (
	"time"
)

type Article struct {
	Id         int       //文章id
	User       User      `gorm:"ForeignKey:Id;references:UserId;"` //用户
	UserId     int       //用户主键
	Content    string    //文章内容
	CreateTime time.Time //创建时间
	Remark     []Remark  `gorm:"ForeignKey:Id;references:RemarkId;"` //多条评论
	RemarkId   int       //评论的主键
}
