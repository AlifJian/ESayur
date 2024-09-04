package Model

import "gorm.io/gorm"

type RoleUser struct {
	gorm.Model
	NamaRole string
}
