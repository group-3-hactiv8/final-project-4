package dto

import (
	"final-project-4/models"
	"final-project-4/pkg/errs"
	"time"

	"github.com/asaskevich/govalidator"
)

type NewProductRequest struct {
	Title      string `json:"title" valid:"required~title is required"`
	Price      int    `json:"price" valid:"required~price is required"`
	Stock      int    `json:"stock" valid:"required~stock is required"`
	CategoryId int    `json:"category_Id" valid:"required~Category ID is required"`
}

// stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go
const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func (p *NewProductRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return errs.NewBadRequest(err.Error())
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

func (p *NewProductRequest) NewProductRequestToModel() *models.Product {
	return &models.Product{
		Title:      p.Title,
		Price:      uint(p.Price),
		Stock:      uint(p.Stock),
		CategoryId: uint(p.CategoryId),
	}
}

type NewProductResponse struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      uint      `json:"price"`
	Stock      uint      `json:"stock"`
	CategoryId uint      `json:"category_Id"`
	CreatedAt  time.Time `json:"created_at"`
}

type AllProductsResponse struct {
	Products []NewProductResponse `json:"products"`
}
