package config

import (
	"github.com/saw/gorm"
	_ "github.com/saw/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "saw:shreyanshsaw@gmail.com/simplerest?charset=utf&&parseTime=True&loc=local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db

}
