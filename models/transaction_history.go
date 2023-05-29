package models

import "gorm.io/gorm"

type TransactionHistory struct {
	gorm.Model
	ProductId  uint `json:"product_id"`
	Product    Product
	UserId     uint `json:"user_id"`
	Quantity   uint `gorm:"not null" json:"quantity"`
	TotalPrice uint `gorm:"not null" json:"total_price"`
}
