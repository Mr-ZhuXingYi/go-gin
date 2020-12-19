package goft

import (
	"github.com/gin-gonic/gin"
)

func ErrMid() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				context.AbortWithStatusJSON(405, gin.H{"err": e})
			}
		}()
		context.Next()
	}
}

func Unwrap(err error, msg ...interface{}) {
	if err != nil {
		if len(msg) > 0 {
			panic(msg[0])
		} else {
			panic(err)
		}
	}
}
