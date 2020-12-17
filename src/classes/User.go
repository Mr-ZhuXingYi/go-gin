package classes

import (
	"github.com/gin-gonic/gin"
	"go-gin/src/goft"
)

type UserClass struct {
}

func NewUserClass() *UserClass {
	return &UserClass{}
}

func (this *UserClass) GetUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, "user success")
	}
}
func (this *UserClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/user", this.GetUser())
}
