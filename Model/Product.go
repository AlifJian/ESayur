package Model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	NamaProduct  string
	HargaProduct int
	CartId       int
}
