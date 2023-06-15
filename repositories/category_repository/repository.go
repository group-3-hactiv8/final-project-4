package category_repository

import (
	"final-project-4/models"
	"final-project-4/pkg/errs"
)

type CategoryRepository interface {
	CreateCategory(category *models.Category) (*models.Category, errs.MessageErr)
	GetAllCategory() ([]models.Category, errs.MessageErr)
	GetProductsByCategoryID(categoryId uint) ([]models.Product, errs.MessageErr)
	GetCategoryById(id uint) (*models.Category, errs.MessageErr)
	UpdateCategory(category *models.Category, ctyUpdate *models.Category) (*models.Category, errs.MessageErr)
	DeleteCategory(category *models.Category) errs.MessageErr
	UpdateSoldProductAmount(*models.Category) errs.MessageErr
	GetCategoryByID(*models.Category) errs.MessageErr
}
