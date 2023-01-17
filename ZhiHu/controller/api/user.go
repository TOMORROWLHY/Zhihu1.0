package api

import (
	"Little_Project/ZhiHu/middleware"
	"Little_Project/ZhiHu/model"
	"Little_Project/ZhiHu/service"
	"Little_Project/ZhiHu/tools"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// Register 用户注册
// @Summary 用户的注册界面
// @Description 进行用户注册
// @Tags 用户注册
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query tools.Resp false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /register [post]
func Register(c *gin.Context) {
	username := c.PostForm("name")
	password := c.PostForm("password")
	location := c.PostForm("location")
	gender := c.PostForm("gender")
	introduction := c.PostForm("introduction")
	work := c.PostForm("work")
	workexp := c.PostForm("workexp")
	//判断用户名或者密码是否为空，减少不必要的资源占用
	if username == "" || password == "" {
		tools.RespParamErr(c)
		return
	}
	//判断用户名是否重复
	flag := service.CheckExist(username)
	if flag {
		tools.NormalErr(c, 500, "用户名已存在")
		return
	}
	err := service.CreatUser(model.User{
		Username:     username,
		Password:     password,
		Location:     location,
		Workexp:      workexp,
		Work:         work,
		Gender:       gender,
		Introduction: introduction,
	})
	if err != nil {
		tools.RsepInternalErr(c)
		return
	}
	tools.RespSuccess(c, 200, "注册成功！", nil)
	//直接在用户注册的时候，创建用户的收藏
	user, _ := service.SearchUserByUsername(username)
	service.CreateCollectQuestion(model.CollectQuestion{
		User:   user,
		UserId: user.Id,
	})
	service.CreateCollectArticle(model.CollectArticle{
		User:   user,
		UserId: user.Id,
	})
}

// Login 用户登录
// @Summary 用户的登录界面
// @Description 进行用户登录
// @Tags 用户登录
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query tools.Resp false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /user/login [post]
func Login(c *gin.Context) {
	username := c.PostForm("name")
	password := c.PostForm("password")
	if username == "" || password == "" {
		tools.RespParamErr(c)
		return
	}
	u, err := service.SearchUserByUsername(username)
	if err != nil {
		if err != sql.ErrNoRows {
			tools.NormalErr(c, 300, "用户不存在")
		} else {
			log.Printf("search user err: %#v", err)
			tools.RsepInternalErr(c)
			return
		}
		return
	}
	if u.Password != password {
		tools.NormalErr(c, 500, "密码错误！")
		return
	}

	tools.RespOK(c, "登录成功")
	// 创建一个我们自己的声明
	claim := model.MyClaims{
		Username: username, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // 过期时间
			Issuer:    "rollFish",                           // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	tokenString, err := token.SignedString(middleware.Secret)
	if err != nil {
		log.Printf("the err :%#v", err)
		return
	}
	tools.RespSuccess(c, 200, tokenString, u)
}

// GetUsernameFromToken 从token中获取username
func GetUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	tools.RespOK(c, username.(string))
}

// ChangeName 修改用户名
// @Summary 用户修改用户名界面
// @Description 进行用用户名的修改
// @Tags 修改用户名
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query tools.Resp false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /user/changename [put]
func ChangeName(c *gin.Context) {
	//检查是否登录
	username, exist := c.Get("username")
	if !exist {
		tools.NormalErr(c, 401, "请先登录！")
		return
	}
	newName := c.PostForm("newname")
	//防止新用户名是空的
	if newName == "" {
		tools.RespParamErr(c)
		return
	}
	err := service.ChangeName(username.(string), newName)
	if err != nil {
		tools.RsepInternalErr(c)
		return
	}
	tools.RespSuccess(c, 200, "修改用户名成功！", newName)
}

// ChangePass 修改密码
// @Summary 用户修改密码界面
// @Description 进行用户密码的修改
// @Tags 修改密码
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query tools.Resp false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /user/changenpass [put]
func ChangePass(c *gin.Context) {
	//检查是否登录
	username, exist := c.Get("username")
	if !exist {
		tools.NormalErr(c, 401, "请先登录！")
		return
	}
	newpass := c.PostForm("newpassword")
	_, err := service.SearchUserByUsername(username.(string))
	if err != nil {
		if err != sql.ErrNoRows {
			tools.NormalErr(c, 300, "用户不存在")
		} else {
			log.Printf("search user err: %#v", err)
			tools.RsepInternalErr(c)
			return
		}
		return
	}
	err = service.ChangePass(username.(string), newpass)
	if err != nil {
		tools.NormalErr(c, 300, "修改密码失败！")
		return
	}
	tools.RespSuccess(c, 200, "修改密码成功！", nil)
}

// MyCollectQuestion 查看用户收藏的问题
// @Summary 查看用户收藏的问题
// @Description 查看收藏
// @Tags 收藏
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query tools.Resp false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /user/collectquestion [get]
func MyCollectQuestion(c *gin.Context) {
	//检查是否登录
	username, exist := c.Get("username")
	if !exist {
		tools.NormalErr(c, 401, "请先登录！")
		return
	}
	qs, err := service.TakeOutQuestion(username.(string))
	if err != nil {
		tools.RsepInternalErr(c)
	}
	//判断是否收藏了问题
	if len(qs) == 0 {
		tools.NormalErr(c, 200, "您还未收藏问题")
		return
	}
	//将问题传递给前端
	tools.RespSuccess(c, http.StatusOK, "以下是您收藏的问题", qs)
}

// MyCollectArticle 查看用户收藏的文章
// @Summary 查看用户收藏的文章
// @Description 查看收藏
// @Tags 收藏
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query tools.Resp false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /user/collectarticle [get]
func MyCollectArticle(c *gin.Context) {
	//检查是否登录
	username, exist := c.Get("username")
	if !exist {
		tools.NormalErr(c, 401, "请先登录！")
		return
	}
	as, err := service.TakeOutArticle(username.(string))
	if err != nil {
		tools.RsepInternalErr(c)
	}
	//判断是否收藏了文章
	if len(as) == 0 {
		tools.NormalErr(c, 200, "您还未收藏问题")
		return
	}
	//将文章传递给前端
	tools.RespSuccess(c, http.StatusOK, "以下是您收藏的问题", as)
}

// ShowMyInformation 展示用户信息
// @Summary 向用户展示自己的个人信息
// @Tags 查看个人信息
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query tools.Resp false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /user/infomation [get]
func ShowMyInformation(c *gin.Context) {
	//检查是否登录
	username, exist := c.Get("username")
	if !exist {
		tools.NormalErr(c, 401, "请先登录！")
		return
	}
	//查看用户信息
	u, err := service.SearchUserByUsername(username.(string))
	if err != nil {
		tools.RsepInternalErr(c)
		return
	}
	tools.RespSuccess(c, 200, "个人信息：", u)
}
