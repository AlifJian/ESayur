package Model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string
	Password   string
	Wallet     int
	RoleUserId int
	CartId     int
	RoleUser   RoleUser
	Cart       Cart
}
