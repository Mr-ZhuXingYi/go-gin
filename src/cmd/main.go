package main

import (
	"go-gin/src/classes"
	"go-gin/src/goft"
)

func main() {
	goft.Start().Mount(
		"v1",
		classes.NewIndexClass(),
		classes.NewUserClass(),
	).Launch()
}
