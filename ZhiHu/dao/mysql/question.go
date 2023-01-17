package mysql

import (
	"Little_Project/ZhiHu/model"
	"go.uber.org/zap"
)

// 创建问题
func CreateQuestion(q model.Question) (err error) {
	tx := DB.Preload("User").Preload("Answer").Create(&q)
	err = tx.Error
	if err != nil {
		zap.L().Error(" create question err: ", zap.Error(err))
		return
	}
	return err
}

// 通过qid查找问题
func SearchQuestionById(qid int) (question model.Question, err error) {
	tx := DB.Preload("User").Preload("Answer").Find(&question)
	err = tx.Error
	if err != nil {
		zap.L().Error("search question by qid err: ", zap.Error(err))
		return
	}
	return question, err
}

// 将回答嵌入到问题中
func InsertAnswerIntoQuestion(answer []model.Answer) (err error) {
	var question model.Question
	tx := DB.Model(&question).Preload("User").Preload("Answer").Update("answer", answer)
	err = tx.Error
	if err != nil {
		zap.L().Error("insert answer into question err: ", zap.Error(err))
		return
	}
	return err
}
