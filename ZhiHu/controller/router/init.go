package router

import (
	"Little_Project/ZhiHu/config/logger"
	"Little_Project/ZhiHu/controller/api"
	_ "Little_Project/ZhiHu/docs"
	"Little_Project/ZhiHu/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//使用自定义的zap日志
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//cors中间接解决跨域问题
	r.Use(middleware.CORS())
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	//用户路由组
	r.POST("/register", api.Register)
	u := r.Group("/user")
	{
		u.Use(middleware.JWTAuthMiddleware())
		u.GET("/get", api.GetUsernameFromToken)
	}
	//u.POST("/register", api.Register)
	u.POST("/login", api.Login)
	u.PUT("/changename", api.ChangeName)
	u.PUT("/changepass", api.ChangePass)
	u.GET("/infomation", api.ShowMyInformation)
	u.GET("/collectquestion", api.MyCollectQuestion)
	u.GET("/collectarticle", api.MyCollectArticle)

	//发布问题、文章、评论
	question := u.Group("/question")
	question.POST("/post", api.PostQuestion)
	question.POST("/answer", api.PostQuestionAnswer)
	question.POST("/collect", api.ColectQuestion)

	article := u.Group("/article")
	article.POST("/post", api.PostArticle)
	article.POST("/collect", api.ColectArticle)

	remark := u.Group("/remark")
	remark.POST("/answer", api.RemarkAnswer)
	remark.POST("/article", api.RemarkArticle)

	r.Run()
	return r
}
