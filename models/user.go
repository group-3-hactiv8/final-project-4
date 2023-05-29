package models

import (
	"final-project-4/helpers"
	"final-project-4/pkg/errs"
	"fmt"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `gorm:"not null" json:"full_name"`
	Email    string `gorm:"not null;uniqueIndex" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Role     string `gorm:"not null" json:"role"`
	Balance  uint   `gorm:"not null" json:"balance"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return errs.NewUnprocessableEntity(err.Error())
	}

	isUnderOrEqualToMillion := govalidator.InRangeInt(u.Balance, 0, 100000000) // kalau 0 tetep true (lower bound)

	if !isUnderOrEqualToMillion {
		return errs.NewUnprocessableEntity("Balance must have non-negative value and under or equal to 100.000.000")
	}

	u.Password = helpers.HashPass(u.Password)

	if u.Role == "admin" || u.Role == "customer" {
		return nil
	} else {
		message := fmt.Sprintf("Invalid Role value: %s", u.Role)
		return errs.NewUnprocessableEntity(message)
	}
}
