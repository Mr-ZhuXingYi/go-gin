package main

import (
	. "go-gin/src/Middleware"
	. "go-gin/src/classes"
	"go-gin/src/goft"
)

func main() {
	//fmt.Println(*goft.InitConfig().Server)
	//fmt.Println(*goft.InitConfig().Config)
	//return
	goft.Start().Beans(goft.NewGormAdapter()).
		Attach(NewUserMiddleware()).
		Mount(
			"v1",
			NewIndexClass(),
			NewUserClass(),
			NewArticle(),
		).Launch()
}
