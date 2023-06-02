package category_pg

import (
	"final-project-4/models"
	"final-project-4/pkg/errs"
	"final-project-4/repositories/category_repository"
	"fmt"

	"gorm.io/gorm"
)

type categoryPG struct {
	db *gorm.DB
}

func NewCategoryPG(db *gorm.DB) category_repository.CategoryRepository {
	return &categoryPG{db: db}
}

func (c *categoryPG) UpdateSoldProductAmount(category *models.Category) errs.MessageErr {
	err := c.db.Model(&models.Category{}).Where("id = ?", category.ID).Update("sold_product_amount", category.SoldProductAmount).Error

	if err != nil {
		message := fmt.Sprintf("category with ID %d not found", category.ID)
		err2 := errs.NewNotFound(message)
		return err2
	}

	return nil
}

func (c *categoryPG) GetCategoryByID(category *models.Category) errs.MessageErr {
	err := c.db.Where("id = ?", category.ID).Take(&category).Error
	// Karna di Take, objek category akan terupdate.

	if err != nil {
		message := fmt.Sprintf("category with ID %v not found", category.ID)
		err2 := errs.NewNotFound(message)
		return err2
	}

	return nil
}
