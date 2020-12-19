package classes

import (
	"go-gin/src/goft"
)

type IndexClass struct {
	*goft.GormAdapter
}

func NewIndexClass() *IndexClass {
	return &IndexClass{}
}

func (this *IndexClass) Index() string {
	return "index:hello world"
}

func (this *IndexClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/index", this.Index)
}
