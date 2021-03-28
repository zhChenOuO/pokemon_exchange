package model

import "gorm.io/gorm"

// Card ...
type Card struct {
	gorm.Model
	Name string
}
