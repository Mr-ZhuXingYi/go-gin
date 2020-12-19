package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Goft struct {
	*gin.Engine
	g           *gin.RouterGroup
	beanFactory *BeanFactory
}

func Start() *Goft {
	g := &Goft{Engine: gin.New(), beanFactory: NewBeanFactory()}
	g.Use(ErrMid())                     //强迫加载的异常处理中间件
	g.beanFactory.setBean(InitConfig()) //整个配置加载进bean中
	return g
}

func (this *Goft) Launch() *Goft {
	var port int32 = 8080
	if config := this.beanFactory.GetBean(new(SysConfig)); config != nil {
		port = config.(*SysConfig).Server.Port
	}
	this.Run(fmt.Sprintf(":%d", port))
	return this
}

func (this *Goft) Attach(f Fairing) *Goft {
	this.Use(func(context *gin.Context) {
		err := f.OnRequest(context)
		if err != nil {
			context.AbortWithStatusJSON(400, "err")
			return
		} else {
			context.Next()
		}
	})
	return this
}

func (this *Goft) Handle(httpMethod, relativePath string, handler interface{}) *Goft {
	h := Convert(handler)
	if h != nil {
		this.g.Handle(httpMethod, relativePath, h)
	}
	return this
}

//设定数据库连接对象
func (this *Goft) Beans(beans ...interface{}) *Goft {
	this.beanFactory.setBean(beans...)
	return this
}

func (this *Goft) Mount(group string, iCalss ...IClass) *Goft {
	this.g = this.Group(group)
	for _, build := range iCalss {
		build.Build(this)
		this.beanFactory.inject(build)
	}
	return this
}
