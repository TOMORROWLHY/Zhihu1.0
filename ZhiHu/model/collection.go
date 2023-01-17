package model

type CollectQuestion struct {
	User            User       `gorm:"ForeignKey:Id;references:UserId;"` //用户
	UserId          int        //用户主键
	CollectQuestion []Question `gorm:"ForeignKey:Id;references:QuestionId;"` //收藏的问题
	QuestionId      int        //问题的主键
}
type CollectArticle struct {
	User           User      `gorm:"ForeignKey:Id;references:UserId;"` //用户
	UserId         int       //用户主键
	CollectArticle []Article `gorm:"ForeignKey:Id;references:ArticleId;"` //收藏的文章
	ArticleId      int       //文章主键
}
