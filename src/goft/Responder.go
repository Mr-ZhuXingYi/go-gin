package goft

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type ResponderList []Responder

var responderList ResponderList

func init() {
	responderList = ResponderList{
		new(StringResonder),
		new(ModelResonder),
		new(ModelsResonder),
	}
}

type Responder interface {
	RespondTo() gin.HandlerFunc
}

func Convert(hander interface{}) gin.HandlerFunc {
	h_value := reflect.ValueOf(hander)
	for _, r := range responderList {
		r_value := reflect.ValueOf(r).Elem()
		if h_value.Type().ConvertibleTo(r_value.Type()) {
			r_value.Set(h_value)
			return r_value.Interface().(Responder).RespondTo()
		}
	}
	return nil
}

type StringResonder func(ctx *gin.Context) string

func (this StringResonder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(http.StatusOK, this(context))
	}
}

type ModelResonder func(ctx *gin.Context) Model

func (this ModelResonder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, this(context))
	}
}

type ModelsResonder func(ctx *gin.Context) Models

func (this ModelsResonder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("Content-type", "application/json")
		context.Writer.WriteString(string(this(context)))
	}
}
