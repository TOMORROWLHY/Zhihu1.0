package mysql

import (
	"Little_Project/ZhiHu/model"
	"go.uber.org/zap"
)

// 创建文章
func CreateArticle(article model.Article) (err error) {
	tx := DB.Preload("User").Preload("Remark").Create(&article)
	err = tx.Error
	if err != nil {
		zap.L().Error("create article err: ", zap.Error(err))
		return
	}
	return err
}

// 通过ArticleID来查找问题
func SearchArticleByID(id int) (article model.Article, err error) {
	tx := DB.Preload("User").Preload("Remark").Find(&article)
	err = tx.Error
	if err != nil {
		zap.L().Error("search article by id err: ", zap.Error(err))
		return
	}
	return article, err
}

// 将评论嵌入到文章中
func InsertRemarkIntoArticle(remark []model.Remark) (err error) {
	var article model.Article
	tx := DB.Model(&article).Preload("User").Preload("Remark").Update("remark", remark)
	err = tx.Error
	if err != nil {
		zap.L().Error("insert remark into article err: ", zap.Error(err))
		return
	}
	return err
}
