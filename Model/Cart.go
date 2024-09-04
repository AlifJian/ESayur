package Model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Products []Product
	Total    int
}
