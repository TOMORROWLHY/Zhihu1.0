package api

import (
	"Little_Project/ZhiHu/model"
	"Little_Project/ZhiHu/service"
	"Little_Project/ZhiHu/tools"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// RemarkAnswer 对回答进行评论
// @Summary 对问题的回答进行评论
// @Description 评论
// @Tags 评论
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query tools.Resp false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /user/remark/answer [post]
func RemarkAnswer(c *gin.Context) {
	//先判断是否登录
	username, exist := c.Get("username")
	if !exist {
		tools.NormalErr(c, 401, "请先登录！")
		return
	}
	id, _ := strconv.Atoi(c.PostForm("Id")) //回答的ID  Answer Id
	content := c.PostForm("remark")
	createtime := time.Now()
	//查找用户id
	uid, err := service.SearchUIDByUsername(username.(string))
	if err != nil {
		tools.RsepInternalErr(c)
		return
	}
	//创建评论
	remark := model.Remark{
		User: model.User{
			Id:       uid,
			Username: username.(string),
		},
		UserId:     uid,
		Content:    content,
		CreateTime: createtime,
	}
	err = service.CreateRemark(remark)
	if err != nil {
		tools.RsepInternalErr(c)
		return
	}
	err = service.InsertRemarkIntoAnswer(id, remark)
	if err != nil {
		tools.RsepInternalErr(c)
		return
	}
	tools.RespSuccess(c, 200, "成功发表评论！", nil)
}

// RemarkArticle 对文章进行评论
// @Summary 对发布的文章进行评论
// @Description 评论
// @Tags 评论
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query tools.Resp false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /user/remark/question [post]
func RemarkArticle(c *gin.Context) {
	//先判断是否登录
	username, exist := c.Get("username")
	if !exist {
		tools.NormalErr(c, 401, "请先登录！")
		return
	}
	aid, _ := strconv.Atoi(c.PostForm("Id")) //文章的ID ArticleId
	content := c.PostForm("remark")
	createtime := time.Now()
	//查找用户id
	uid, err := service.SearchUIDByUsername(username.(string))
	if err != nil {
		tools.RsepInternalErr(c)
		return
	}
	//创建评论
	remark := model.Remark{
		User: model.User{
			Id:       uid,
			Username: username.(string),
		},
		UserId:     uid,
		Content:    content,
		CreateTime: createtime,
	}
	err = service.CreateRemark(remark)
	if err != nil {
		tools.RsepInternalErr(c)
		return
	}
	err = service.InsertRemarkIntoArticle(aid, remark)
	if err != nil {
		tools.RsepInternalErr(c)
	}
	tools.RespSuccess(c, 200, "成功发布评论！", nil)
}
