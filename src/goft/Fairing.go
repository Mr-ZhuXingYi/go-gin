package goft

import "github.com/gin-gonic/gin"

type Fairing interface {
	OnRequest(ctx *gin.Context) error
}
