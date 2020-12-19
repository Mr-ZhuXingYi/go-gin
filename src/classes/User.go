package classes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/src/Model"
	"go-gin/src/goft"
)

type UserClass struct {
	*goft.GormAdapter
	Age *goft.Value `prefix:"user.source"`
}

func NewUserClass() *UserClass {
	return &UserClass{}
}

func (this *UserClass) UserList(context *gin.Context) string {
	var c int
	this.Table("t_database").Count(&c)
	fmt.Println("count:", c)
	return "user success"
}

func (this *UserClass) UserDetail(context *gin.Context) goft.Model {
	u := new(Model.UserModel)
	goft.Unwrap(context.ShouldBindUri(u), "参数不合法")
	u.Name = this.Age.String()
	return u
}

func (this *UserClass) UserDetails(context *gin.Context) goft.Models {
	userModels := []*Model.UserModel{
		{Name: "zxy", Age: 18},
		{Name: "zxy1", Age: 19},
	}
	return goft.MakeModels(userModels)
}

func (this *UserClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/user", this.UserList)
	goft.Handle("GET", "/user_detail/:id", this.UserDetail)
	goft.Handle("GET", "/user_details", this.UserDetails)
}
