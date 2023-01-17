package service

import (
	"Little_Project/ZhiHu/dao/mysql"
	"Little_Project/ZhiHu/model"
)

// 创建文章
func CreateArticle(article model.Article) (err error) {
	err = mysql.CreateArticle(article)
	return err
}
