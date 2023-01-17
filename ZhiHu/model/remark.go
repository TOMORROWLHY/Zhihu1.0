package model

import "time"

type Remark struct {
	Id         int       //评论的主键
	User       User      `gorm:"ForeignKey:Id;references:UserId"` //用户
	UserId     int       //用户主键
	Content    string    //评论的内容
	CreateTime time.Time //创建时间
}
