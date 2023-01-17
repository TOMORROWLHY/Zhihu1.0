package api

import (
	"Little_Project/ZhiHu/model"
	"Little_Project/ZhiHu/service"
	"Little_Project/ZhiHu/tools"
	"github.com/gin-gonic/gin"
	"time"
)

// PostQuestion 发布问题
// @Summary 发布问题界面
// @Description 用户可以自行发布自己的问题
// @Tags 问题
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query tools.Resp false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /user/question/post [post]
func PostQuestion(c *gin.Context) {
	//先判断是否登录
	username, exist := c.Get("username")
	if !exist {
		tools.NormalErr(c, 401, "请先登录！")
		return
	}
	content := c.PostForm("question")
	createtime := time.Now()
	//查找用户id
	id, err := service.SearchUIDByUsername(username.(string))
	if err != nil {
		tools.RsepInternalErr(c)
		return
	}
	//添加问题
	q := model.Question{
		Content: content,
		User: model.User{
			Id:       id,
			Username: username.(string),
		},
		UserId:     id,
		CreateTime: createtime,
	}
	err = service.CreateQuestion(q)
	if err != nil {
		tools.RsepInternalErr(c)
		return
	}
	tools.RespSuccess(c, 200, "发布问题成功！", nil)
}

//
