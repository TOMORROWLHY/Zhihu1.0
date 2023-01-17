package api

import (
	"Little_Project/ZhiHu/model"
	"Little_Project/ZhiHu/service"
	"Little_Project/ZhiHu/tools"
	"github.com/gin-gonic/gin"
	"time"
)

// PostArticle 发布文章
// @Summary 发布文章
// @Description 用户可以发布自己的所思所想
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query tools.Resp false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /user/article/post [post]
func PostArticle(c *gin.Context) {
	//先判断是否登录
	username, exist := c.Get("username")
	if !exist {
		tools.NormalErr(c, 401, "请先登录！")
		return
	}
	content := c.PostForm("article")
	createtime := time.Now()
	//查找用户id
	id, err := service.SearchUIDByUsername(username.(string))
	if err != nil {
		tools.RsepInternalErr(c)
		return
	}
	//创建文章
	article := model.Article{
		User: model.User{
			Id:       id,
			Username: username.(string),
		},
		UserId:     id,
		Content:    content,
		CreateTime: createtime,
	}
	err = service.CreateArticle(article)
	if err != nil {
		tools.RsepInternalErr(c)
		return
	}
	tools.RespOK(c, "发布文章成功！")
}
