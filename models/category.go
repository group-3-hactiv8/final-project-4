package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Type              string `gorm:"not null" json:"type"`
	SoldProductAmount uint   `json:"sold_product_amount"`
}
