package Model

import "fmt"

type UserModel struct {
	Name string
	Age  int `uri:"id" binding:"required,gt=0"`
}

func (this *UserModel) String() string {
	return fmt.Sprintf("name:%s,age:%d", this.Name, this.Age)
}
