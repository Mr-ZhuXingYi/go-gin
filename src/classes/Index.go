package classes

import (
	"github.com/gin-gonic/gin"
	"go-gin/src/goft"
)

type IndexClass struct {
}

func NewIndexClass() *IndexClass {
	return &IndexClass{}
}

func (this *IndexClass) GetIndex() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, "hello world")
	}
}

func (this *IndexClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/index", this.GetIndex())
}
