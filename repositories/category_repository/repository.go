package category_repository

import (
	"final-project-4/models"
	"final-project-4/pkg/errs"
)

type CategoryRepository interface {
	UpdateSoldProductAmount(*models.Category) errs.MessageErr
	GetCategoryByID(*models.Category) errs.MessageErr
}
