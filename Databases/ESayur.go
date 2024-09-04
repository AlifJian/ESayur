package Databases

import (
	model "ESayur/Model"
	util "ESayur/Util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (db *gorm.DB, err error) {
	stringConnect := "root:@tcp(127.0.0.1:3306)/esayur?charset=utf8mb4&parseTime=True&loc=Local"

	defer util.Catch()

	db, err = gorm.Open(mysql.Open(stringConnect), &gorm.Config{})
	if err != nil {
		panic("Databases Error:" + err.Error())
	}

	db.AutoMigrate(
		&model.User{},
		&model.Cart{},
		&model.Product{},
		&model.RoleUser{},
	)
	return
}
