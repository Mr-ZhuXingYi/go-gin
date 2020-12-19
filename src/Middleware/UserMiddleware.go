package Middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type UserMiddleware struct {
}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (this *UserMiddleware) OnRequest(ctx *gin.Context) error {
	fmt.Println("middle")
	fmt.Println(ctx.Query("name"))
	if time.Now().Second()%2 == 0 {
		return fmt.Errorf("err")
	}
	return nil
}
