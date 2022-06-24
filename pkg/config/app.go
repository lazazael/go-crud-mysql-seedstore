package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "laza:Laza/rest?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
