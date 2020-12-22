package classes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/src/Model"
	"go-gin/src/goft"
	"time"
)

type ArticleClass struct {
	*goft.GormAdapter
}

func NewArticle() *ArticleClass {
	return &ArticleClass{}
}

func (this *ArticleClass) ArticleDetail(context *gin.Context) goft.Model {
	article := Model.NewArticleModel()
	article.Id = 1
	article.Title = "beijing"
	article.Content = "fdsfsfsafdsfsf"

	goft.Task(task, func() {
		fmt.Println("end", t)
	}, article.Id)
	return article
}

var t int

func task(param ...interface{}) {
	t++
	fmt.Println("task:", t, "   id:", param[0])
	time.Sleep(time.Second * 5)
}

func (this *ArticleClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/article", this.ArticleDetail)
}
