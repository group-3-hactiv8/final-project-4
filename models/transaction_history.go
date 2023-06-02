package models

import (
	"final-project-4/pkg/errs"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type TransactionHistory struct {
	gorm.Model
	ProductId  uint `gorm:"not null" json:"product_id"`
	Product    Product
	UserId     uint `gorm:"not null" json:"user_id"`
	User       User
	Quantity   uint `gorm:"not null" json:"quantity"`
	TotalPrice uint `gorm:"not null" json:"total_price"`
}

func (th *TransactionHistory) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(th)

	if err != nil {
		return errs.NewUnprocessableEntity(err.Error())
	}
	return nil
}
