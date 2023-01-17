package api

import (
	"Little_Project/ZhiHu/service"
	"Little_Project/ZhiHu/tools"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ColectQuestion 收藏问题
// @Summary 对该问题感兴趣可以进行收藏
// @Description 对该问题感兴趣可以进行收藏
// @Tags 收藏问题
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query tools.Resp false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /user/question/collect [post]
func ColectQuestion(c *gin.Context) {
	//先判断是否登录
	username, exist := c.Get("username")
	if !exist {
		tools.NormalErr(c, 401, "请先登录！")
		return
	}
	qidString := c.PostForm("questionID")
	qid, _ := strconv.Atoi(qidString) //转为int类型
	err := service.UpdateColLectQuestion(username.(string), qid)
	if err != nil {
		tools.RsepInternalErr(c)
		return
	}
	tools.RespSuccess(c, 200, "收藏问题成功！", nil)
}

// ColectArticle 收藏文章
// @Summary 对该文章感兴趣可以进行收藏
// @Description 对该文章感兴趣可以进行收藏
// @Tags 收藏文章
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query tools.Resp false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /user/article/collect [post]
func ColectArticle(c *gin.Context) {
	//先判断是否登录
	username, exist := c.Get("username")
	if !exist {
		tools.NormalErr(c, 401, "请先登录！")
		return
	}
	aidString := c.PostForm("articleID")
	aid, _ := strconv.Atoi(aidString) //转为int类型
	err := service.UpdateColLectArticle(username.(string), aid)
	if err != nil {
		tools.RsepInternalErr(c)
		return
	}
	tools.RespSuccess(c, 200, "收藏文章成功！", nil)
}
