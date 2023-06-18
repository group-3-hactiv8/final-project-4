package product_repository

import (
	"final-project-4/models"
	"final-project-4/pkg/errs"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) (*models.Product, errs.MessageErr)
	GetAllProducts() (*[]models.Product, uint, errs.MessageErr)
	GetProductByID(id uint) (*models.Product, errs.MessageErr)
	GetProductByIdUpdate(id uint) (*models.Product, errs.MessageErr)
	UpdateStock(product *models.Product) errs.MessageErr
	UpdateProducts(product *models.Product, productUpdate *models.Product) (*models.Product, errs.MessageErr)
	DeleteProducts(product *models.Product) errs.MessageErr
}
