package goft

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

type GormAdapter struct {
	*gorm.DB
}

func (this *GormAdapter) Name() string {
	return "GormAdapter"
}
func NewGormAdapter() *GormAdapter {
	db, err := gorm.Open("mysql",
		"plat:mTAerlrufO@tcp(192.168.1.205:3600)/plat_gaea?charset=utf8&parseTime=true&loc=Local&multiStatements=true")
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(5)                   //最大空闲数
	db.DB().SetMaxOpenConns(10)                  //最大打开连接数
	db.DB().SetConnMaxLifetime(time.Second * 30) //空闲连接生命周期
	return &GormAdapter{DB: db}
}
