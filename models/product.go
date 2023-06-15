package models

import (
	"final-project-4/pkg/errs"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title      string `gorm:"not null"`
	Price      uint   `gorm:"not null"`
	Stock      uint   `gorm:"not null"`
	CategoryId uint   `gorm:"not null"`
	Category   Category
	TransactionHistories []TransactionHistory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go
const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return errs.NewUnprocessableEntity(err.Error())
	}

	isMoreThanFour := govalidator.InRangeInt(p.Stock, 5, MaxInt) // kalau 5 tetep true (lower bound)

	if !isMoreThanFour {
		return errs.NewUnprocessableEntity("Stock must have value more than four")
	}

	isLTE50Mil := govalidator.InRangeInt(p.Price, 0, 50000000) // kalau 5 tetep true (lower bound)

	if !isLTE50Mil {
		return errs.NewUnprocessableEntity("Price must have non-negative value and less than or equal to Rp 50.000.000")
	}

	return nil
}
