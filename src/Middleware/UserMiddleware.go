package Middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserMiddleware struct {
}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (this *UserMiddleware) OnRequest(ctx *gin.Context) error {
	fmt.Println("middle")
	fmt.Println(ctx.Query("name"))
	return nil
}
