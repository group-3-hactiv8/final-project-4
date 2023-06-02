package product_pg

import (
	"final-project-4/models"
	"final-project-4/pkg/errs"
	"final-project-4/repositories/product_repository"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type productPG struct {
	db *gorm.DB
}

func NewProductPG(db *gorm.DB) product_repository.ProductRepository {
	return &productPG{db: db}
}

func (p *productPG) CreateProduct(product *models.Product) (*models.Product, errs.MessageErr) {
	err := p.db.Where("id = ?", product.CategoryId).Take(&models.Category{}).Error

	if err != nil {
		message := fmt.Sprintf("Category with ID %v not found", product.CategoryId)
		err2 := errs.NewNotFound(message)
		return nil, err2
	}

	err = p.db.Create(product).Error

	if err != nil {
		log.Println(err.Error())
		message := fmt.Sprintf("Failed to create a Product with title %s", product.Title)
		errFailedToCreate := errs.NewInternalServerError(message)
		return nil, errFailedToCreate
	}

	return product, nil
}

func (p *productPG) GetAllProducts() (*[]models.Product, uint, errs.MessageErr) {
	var allProducts *[]models.Product
	result := p.db.Find(&allProducts)

	if err := result.Error; err != nil {
		log.Println(err.Error())
		error := errs.NewInternalServerError("Can't fetch all Products data")
		return nil, 0, error
	}

	totalCount := result.RowsAffected

	return allProducts, uint(totalCount), nil
}

func (p *productPG) GetProductByID(product *models.Product) errs.MessageErr {
	err := p.db.Where("id = ?", product.ID).Take(&product).Error
	// Karna di Take, objek product akan terupdate.

	if err != nil {
		message := fmt.Sprintf("product with ID %v not found", product.ID)
		err2 := errs.NewNotFound(message)
		return err2
	}

	return nil
}

func (p *productPG) UpdateStock(product *models.Product) errs.MessageErr {
	err := p.db.Model(&models.Product{}).Where("id = ?", product.ID).Update("stock", product.Stock).Error

	if err != nil {
		message := fmt.Sprintf("product with ID %d not found", product.ID)
		err2 := errs.NewNotFound(message)
		return err2
	}

	return nil
}
