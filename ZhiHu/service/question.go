package service

import (
	"Little_Project/ZhiHu/dao/mysql"
	"Little_Project/ZhiHu/model"
)

// 创建问题
func CreateQuestion(q model.Question) error {
	err := mysql.CreateQuestion(q)
	return err
}

// 查找问题by Qid
func SearchQuestionByQid(qid int) (question model.Question, err error) {
	question, err = mysql.SearchQuestionById(qid)
	return question, err
}
