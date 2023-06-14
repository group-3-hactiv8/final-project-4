package services

import (
	"final-project-4/dto"
	"final-project-4/pkg/errs"
	"final-project-4/repositories/category_repository"
)

type CategoryService interface {
	CreateCategory(payload *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.MessageErr)
	GetAllCategory() ([]dto.GetAllCategoryResponse, errs.MessageErr)
	UpdateCategory(id uint, payload *dto.NewCategoryRequest) (*dto.UpdateCategoryResponse, errs.MessageErr)
	DeleteCategory(id uint) (*dto.DeleteCategoryResponse, errs.MessageErr)
}

type categoryService struct {
	categoryRepo category_repository.CategoryRepository
}

func NewCategoryService(categoryRepo category_repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
} 

func (c *categoryService) CreateCategory(payload *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.MessageErr) {
	newCty := payload.NewCategoryRequestToModel()

	createCategory, err := c.categoryRepo.CreateCategory(newCty)
	if err != nil {
		return nil, err
	}

	response := &dto.NewCategoryResponse {
		ID: createCategory.ID,
		Type: createCategory.Type,
		CreatedAt: createCategory.CreatedAt,
	}
	return response, nil
}


func (c *categoryService) GetAllCategory() ([]dto.GetAllCategoryResponse, errs.MessageErr) {
	allCategories, err := c.categoryRepo.GetAllCategory()

	if err != nil {
		return nil, err
	}

	response := []dto.GetAllCategoryResponse{}
	for _, cty := range allCategories {
		products, err := c.categoryRepo.GetProductsByCategoryID(cty.ID)
		if err != nil {
			return nil, err
		}

		productsResponses := []dto.ProductForGetAllCategoryResponse{}
		for _, product := range products {
			productsResponses = append(productsResponses, dto.ProductForGetAllCategoryResponse{
				ID:          product.ID,
				Title:       product.Title,
				Price: 		 product.Price,
				Stock:      product.Stock,
				CreatedAt:  product.CreatedAt,
				UpdatedAt:   product.UpdatedAt,
			})
		}

		response = append(response, dto.GetAllCategoryResponse{
			ID:        cty.ID,
			Type:      cty.Type,
			SoldProductAmount: cty.SoldProductAmount,
			CreatedAt: cty.CreatedAt,
			UpdatedAt: cty.UpdatedAt,
			Tasks:     productsResponses,
		})
	}

	return response, nil
}

func (c *categoryService) UpdateCategory(id uint, payload *dto.NewCategoryRequest) (*dto.UpdateCategoryResponse, errs.MessageErr)  {
	category, err := c.categoryRepo.GetCategoryById(id)

	if id < 1 {
		idError := errs.NewBadRequest("ID value must be positive")
		return nil, idError
	}

	if err != nil {
		return nil, err
	}

	updateNewCategory := payload.NewCategoryRequestToModel()

	updateCategory, err2 := c.categoryRepo.UpdateCategory(category, updateNewCategory)

	if err2 != nil {
		return nil, err2
	}

	response := &dto.UpdateCategoryResponse {
		ID: updateCategory.ID,
		Type: updateCategory.Type,
		UpdatedAt: category.UpdatedAt,
	}
	return response, nil
}

func (c *categoryService) DeleteCategory(id uint) (*dto.DeleteCategoryResponse, errs.MessageErr) {
	category, err := c.categoryRepo.GetCategoryById(id)

	if err != nil {
		return nil, err
	}

	if err := c.categoryRepo.DeleteCategory(category); err != nil {
		return nil, err
	}

	deleteResponse := &dto.DeleteCategoryResponse{
		Message: "Category has been successfully deleted",
	}
	return deleteResponse, nil
}