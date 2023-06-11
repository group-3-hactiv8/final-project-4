package dto

import (
	"time"
	"final-project-4/models"
	"final-project-4/pkg/errs"

	"github.com/asaskevich/govalidator"
)

type NewCategoryRequest struct {
	Type string `json:"type" valid:"required~Your Type is required"`
}

func (newCty *NewCategoryRequest) NewCategoryRequestToModel() *models.Category {
	return &models.Category{
		Type: newCty.Type,
	}
}

func (newCty *NewCategoryRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(newCty)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}


type NewCategoryResponse struct {
	ID   				uint 	 	`json:"id"`
	Type 				string 		`json:"type"`
	SoldProductAmount  	uint  		`json:"sold_product_amount"`
	CreatedAt 			time.Time 	`json:"created_at"`
}

type UpdateCategoryRequest struct {
	Type string `json:"type" valid:"required~Your Type is required"`
}

func (newCty *UpdateCategoryRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(newCty)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

type UpdateCategoryResponse struct {
	ID   				uint    	`json:"id"`
	Type 				string 		`json:"type"`
	SoldProductAmount  	uint  		`json:"sold_product_amount"`
	UpdatedAt 			time.Time 	`json:"updated_at"`
}

type GetAllCategoryResponse struct {
	ID        uint       `json:"id"`
	Type      string     `json:"type"`
	SoldProductAmount  	uint  		`json:"sold_product_amount"`	
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Tasks     []ProductForGetAllCategoryResponse `json:"Product"`
}

type ProductForGetAllCategoryResponse struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      uint      `json:"price"`
	Stock      uint      `json:"stock"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt 			time.Time 	`json:"updated_at"`
}

type DeleteCategoryResponse struct {
	Message string `json:"message"`
}
