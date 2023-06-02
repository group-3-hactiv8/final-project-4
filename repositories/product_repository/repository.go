package product_repository

import (
	"final-project-4/models"
	"final-project-4/pkg/errs"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) (*models.Product, errs.MessageErr)
	GetAllProducts() (*[]models.Product, uint, errs.MessageErr)
	GetProductByID(product *models.Product) errs.MessageErr
	UpdateStock(product *models.Product) errs.MessageErr
}
