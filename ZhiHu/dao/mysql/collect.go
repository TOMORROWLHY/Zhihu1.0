package mysql

import (
	"Little_Project/ZhiHu/model"
	"go.uber.org/zap"
)

// 通过用户名查找该用户收藏的问题
func FindCollectQuestion(username string) (mycollect model.CollectQuestion, err error) {
	tx := DB.Preload("User").Preload("CollectQuestion").Where("username=?", username).Find(&mycollect)
	err = tx.Error
	if err != nil {
		zap.L().Error("search mycollection err: ", zap.Error(err))
		return
	}
	return mycollect, err
}

// 通过用户名查找该用户收藏的文章
func FindCollectArticle(username string) (mycollect model.CollectArticle, err error) {
	tx := DB.Preload("User").Preload("CollectArticle").Where("username=?", username).Find(&mycollect)
	err = tx.Error
	if err != nil {
		zap.L().Error("search mycollection err: ", zap.Error(err))
		return
	}
	return mycollect, err
}

// 创建用户收藏的问题
func CreateCollectQuestion(cq model.CollectQuestion) (err error) {
	tx := DB.Preload("User").Preload("CollectArticle").Create(&cq)
	err = tx.Error
	if err != nil {
		zap.L().Error("create collect question failed, err: ", zap.Error(err))
		return
	}
	return err
}

// 创建用户收藏的文章
func CreateCollectArticle(ca model.CollectArticle) (err error) {
	tx := DB.Preload("User").Preload("CollectArticle").Create(&ca)
	err = tx.Error
	if err != nil {
		zap.L().Error("create collect article failed, err: ", zap.Error(err))
		return
	}
	return err
}

// 更新用户收藏的问题
func UpdateCollectQuestion(username string, questions []model.Question) (err error) {
	var cq model.CollectQuestion
	tx := DB.Model(&cq).Preload("User").Preload("CollectArticle").Where("username=?", username).Update("collect_question=?", questions)
	err = tx.Error
	if err != nil {
		zap.L().Error("update collect questions failed, err: ", zap.Error(err))
		return
	}
	return err
}

// 更新用户收藏的文章
func UpdateCollectArticle(username string, articles []model.Article) (err error) {
	var ca model.CollectArticle
	tx := DB.Model(&ca).Preload("User").Preload("CollectArticle").Where("username=?", username).Update("collect_question=?", articles)
	err = tx.Error
	if err != nil {
		zap.L().Error("update collect articles failed, err: ", zap.Error(err))
		return
	}
	return err
}
