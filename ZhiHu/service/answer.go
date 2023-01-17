package service

import (
	"Little_Project/ZhiHu/dao/mysql"
	"Little_Project/ZhiHu/model"
)

// 创建回答
func CreateAnswer(answer model.Answer) (err error) {
	err = mysql.CreateAnswer(answer)
	return err
}

// 从问题中提取所有回答，并添加新的回答
func InsertAnswerIntoQuestion(qid int, answer model.Answer) (err error) {
	question, _ := mysql.SearchQuestionById(qid)
	question.Answer = append(question.Answer, answer)
	err = mysql.InsertAnswerIntoQuestion(question.Answer)
	return err
}
