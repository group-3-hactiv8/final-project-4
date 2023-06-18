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

func (p *productPG) GetProductByID(productID uint) (*models.Product, errs.MessageErr) {
	product := &models.Product{}
	err := p.db.First(product, productID).Error

	if err != nil {
		message := fmt.Sprintf("Product with ID %v not found", productID)
		err := errs.NewNotFound(message)
		return nil, err
	}

	return product, nil
}

func (c *productPG) GetProductByIdUpdate(id uint) (*models.Product, errs.MessageErr){
	var product models.Product
	result := c.db.First(&product, id)

	if err := result.Error; err != nil {
		log.Println("Error : ",err.Error())
		error := errs.NewNotFound(fmt.Sprintf("failed to get Product by id :", product.ID))
		return nil, error
	}
	return &product, nil
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

func (c *productPG) UpdateProducts(product *models.Product, productUpdate *models.Product) (*models.Product, errs.MessageErr) {
	err := c.db.Model(product).Updates(productUpdate).Error
	if err != nil {
		message := fmt.Sprintf("Failed to Update Product with Id : %v", product.ID)
		err2 := errs.NewNotFound(message)
		return nil, err2
	}
	return product, nil
}

func (c *productPG) DeleteProducts(product *models.Product) errs.MessageErr {
	result := c.db.Delete(product)

	if err := result.Error; err != nil {
		log.Println("Error : ",err.Error())
		error := errs.NewInternalServerError(fmt.Sprintf("Failed to delete Product by id : %v", product.ID))
		return error
	}
	return  nil
}