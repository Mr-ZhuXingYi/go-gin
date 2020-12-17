package goft

import "github.com/gin-gonic/gin"

type Goft struct {
	*gin.Engine
	g *gin.RouterGroup
}

func Start() *Goft {
	return &Goft{Engine: gin.New()}
}

func (this *Goft) Launch() *Goft {
	this.Run(":8080")
	return this
}

func (this *Goft) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) *Goft {
	this.g.Handle(httpMethod, relativePath, handlers...)
	return this
}

func (this *Goft) Mount(group string, iCalss ...IClass) *Goft {
	this.g = this.Group(group)
	for _, build := range iCalss {
		build.Build(this)
	}
	return this
}
