package mysql

import (
	"Little_Project/ZhiHu/model"
	"go.uber.org/zap"
)

// 通过用户名查询回答Answer
func SearchAnswer(username string) (answer model.Answer, err error) {
	tx := DB.Preload("User").Preload("Remark").Where("username=?", username).Find(&answer)
	err = tx.Error
	if err != nil {
		zap.L().Error("search answer err: ", zap.Error(err))
		return
	}
	return answer, err
}

// 创建回答
func CreateAnswer(answer model.Answer) (err error) {
	tx := DB.Preload("User").Preload("Remark").Create(&answer)
	err = tx.Error
	if err != nil {
		zap.L().Error("create answer err: ", zap.Error(err))
		return
	}
	return err
}

// 通过answer id来查找回答
func SearchAnswerByID(id int) (answer model.Answer, err error) {
	tx := DB.Preload("User").Preload("Remark").Find(&answer)
	err = tx.Error
	if err != nil {
		zap.L().Error("search answer by id err: ", zap.Error(err))
		return
	}
	return answer, err
}

// 将评论嵌入到回答中
func InsertRemarkIntoAnswer(remark []model.Remark) (err error) {
	var answer model.Answer
	tx := DB.Model(&answer).Preload("User").Preload("Remark").Update("remark", remark)
	err = tx.Error
	if err != nil {
		zap.L().Error("insert remark into answer err: ", zap.Error(err))
		return
	}
	return err
}
