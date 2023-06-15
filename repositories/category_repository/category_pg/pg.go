package category_pg

import (
	"final-project-4/models"
	"final-project-4/pkg/errs"
	"final-project-4/repositories/category_repository"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type categoryPG struct {
	db *gorm.DB
}

func NewCategoryPG(db *gorm.DB) category_repository.CategoryRepository {
	return &categoryPG{db: db}
}


func (c *categoryPG) CreateCategory (newCategory *models.Category) (*models.Category, errs.MessageErr) {

	if err := c.db.Create(newCategory).Error; err != nil {
		log.Println(err.Error())
		message := fmt.Sprintf("Failed To Create Category with Type : %s", newCategory.Type)
		error := errs.NewInternalServerError(message)
		return nil, error
	}

	return newCategory, nil
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

func (c *categoryPG) GetCategoryById(id uint) (*models.Category, errs.MessageErr){
	var category models.Category
	result := c.db.First(&category, id)

	if err := result.Error; err != nil {
		log.Println("Error : ",err.Error())
		error := errs.NewNotFound(fmt.Sprintf("failed to get category by id :", category.ID))
		return nil, error
	}
	return &category, nil
}

func (r *categoryPG) GetAllCategory() ([]models.Category, errs.MessageErr) {
	var categories []models.Category
	err := r.db.Find(&categories).Error

	if err != nil {
		log.Println("Error : ",err.Error())
		error := errs.NewInternalServerError("Failed to fetching all Category")
		return nil, error
	}

	return categories, nil
}

func (r *categoryPG) GetProductsByCategoryID(categoryId uint) ([]models.Product, errs.MessageErr) {
	var tasks []models.Product
	err := r.db.Where("category_id = ?", categoryId).Find(&tasks).Error

	if err != nil {
		log.Println("Error : ",err.Error())
		error := errs.NewNotFound(fmt.Sprintf("failed to get Product by id :", categoryId))
		return nil, error
	}

	return tasks, nil
}

func (c *categoryPG) UpdateCategory(category *models.Category, ctyUpdate *models.Category) (*models.Category, errs.MessageErr) {
	err := c.db.Model(category).Updates(ctyUpdate).Error
	if err != nil {
		message := fmt.Sprintf("Failed to Update Category with Id : %v", category.ID)
		err2 := errs.NewNotFound(message)
		return nil, err2
	}
	return category, nil
}

func (c *categoryPG) DeleteCategory(category *models.Category) errs.MessageErr {
	result := c.db.Delete(category)

	if err := result.Error; err != nil {
		log.Println("Error : ",err.Error())
		error := errs.NewInternalServerError(fmt.Sprintf("Failed to delete Category by id : %v", category.ID))
		return error
	}
	return  nil
}
