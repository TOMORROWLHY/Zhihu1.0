package api

import (
	"Little_Project/ZhiHu/model"
	"Little_Project/ZhiHu/service"
	"Little_Project/ZhiHu/tools"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// PostQuestionAnswer 对问题进行回答
// @Summary 对问题进行回答
// @Description 用户可以自行回答问题
// @Tags 回答
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query tools.Resp false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /user/question/answer [post]
func PostQuestionAnswer(c *gin.Context) {
	//先判断是否登录
	username, exist := c.Get("username")
	if !exist {
		tools.NormalErr(c, 401, "请先登录！")
		return
	}
	qidString := c.PostForm("questionID")
	qid, _ := strconv.Atoi(qidString) //转为int类型
	content := c.PostForm("answer")
	createtime := time.Now()
	//查找用户id
	uid, err := service.SearchUIDByUsername(username.(string))
	if err != nil {
		tools.RsepInternalErr(c)
		return
	}
	//添加回答
	answer := model.Answer{
		Content: content,
		User: model.User{
			Id:       uid,
			Username: username.(string),
		},
		UserId:     uid,
		CreateTime: createtime,
	}
	err = service.CreateAnswer(answer)
	if err != nil {
		tools.RsepInternalErr(c)
		return
	}
	err = service.InsertAnswerIntoQuestion(qid, answer)
	if err != nil {
		tools.RsepInternalErr(c)
		return
	}
	tools.RespSuccess(c, 200, "回答问题成功！", nil)
}
