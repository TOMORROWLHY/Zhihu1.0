package mysql

import (
	"Little_Project/ZhiHu/model"
	"go.uber.org/zap"
)

// 创建评论
func CreateRemark(remark model.Remark) (err error) {
	tx := DB.Preload("User").Create(&remark)
	err = tx.Error
	if err != nil {
		zap.L().Error("create remark err: ", zap.Error(err))
		return
	}
	return err
}
