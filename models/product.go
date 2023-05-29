package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title      string `gorm:"not null" json:"title"`
	Price      uint   `gorm:"not null" json:"price"`
	Stock      uint   `gorm:"not null" json:"stock"`
	CategoryId uint   `json:"category_id"`
	Category   Category
}

// func (sc *SocialMedia) BeforeCreate(tx *gorm.DB) error {
// 	_, err := govalidator.ValidateStruct(sc)

// 	if err != nil {
// 		return errs.NewUnprocessableEntity(err.Error())
// 	}
// 	return nil
// }
