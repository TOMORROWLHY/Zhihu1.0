package service

import (
	"Little_Project/ZhiHu/dao/mysql"
	"Little_Project/ZhiHu/model"
)

// 查看自己收藏的问题 取出其中的问题
func TakeOutQuestion(username string) (qs []model.Question, err error) {
	mycollect, err := mysql.FindCollectQuestion(username)
	qs = mycollect.CollectQuestion
	return qs, err
}

// 查看自己收藏的文章 取出其中的文章
func TakeOutArticle(username string) (as []model.Article, err error) {
	mycollect, err := mysql.FindCollectArticle(username)
	as = mycollect.CollectArticle
	return as, err
}

// 创建收藏的问题
func CreateCollectQuestion(cq model.CollectQuestion) (err error) {
	err = mysql.CreateCollectQuestion(cq)
	return err
}

// 创建收藏的问题
func CreateCollectArticle(ca model.CollectArticle) (err error) {
	err = mysql.CreateCollectArticle(ca)
	return err
}

// 更新收藏的问题
func UpdateColLectQuestion(username string, qid int) (err error) {
	question, _ := mysql.SearchQuestionById(qid)
	collectquestion, err := mysql.FindCollectQuestion(username)
	collectquestion.CollectQuestion = append(collectquestion.CollectQuestion, question)
	return err
}

// 更新收藏的文章
func UpdateColLectArticle(username string, aid int) error {
	article, _ := mysql.SearchArticleByID(aid)
	collectarticle, err := mysql.FindCollectArticle(username)
	collectarticle.CollectArticle = append(collectarticle.CollectArticle, article)
	return err
}
