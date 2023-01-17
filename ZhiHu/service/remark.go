package service

import (
	"Little_Project/ZhiHu/dao/mysql"
	"Little_Project/ZhiHu/model"
)

// 创建评论
func CreateRemark(remark model.Remark) (err error) {
	err = mysql.CreateRemark(remark)
	return err
}

// 从回答中提取所有的评论，并添加新的评论
func InsertRemarkIntoAnswer(id int, remark model.Remark) (err error) {
	answer, _ := mysql.SearchAnswerByID(id)
	answer.Remark = append(answer.Remark, remark)
	err = mysql.InsertRemarkIntoAnswer(answer.Remark)
	return err
}

// 从文章中提取所有的评论，并添加新的评论
func InsertRemarkIntoArticle(id int, remark model.Remark) (err error) {
	article, _ := mysql.SearchArticleByID(id)
	article.Remark = append(article.Remark, remark)
	err = mysql.InsertRemarkIntoArticle(article.Remark)
	return err
}
