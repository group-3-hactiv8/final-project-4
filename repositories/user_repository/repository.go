package user_repository

import (
	"final-project-4/models"
	"final-project-4/pkg/errs"
)

type UserRepository interface {
	SeedingAdmin()
	RegisterUser(user *models.User) (*models.User, errs.MessageErr)
	GetUserByID(id uint) (*models.User, errs.MessageErr)
	GetUserByEmail(user *models.User) errs.MessageErr
	UpdateUser(user *models.User) (*models.User, errs.MessageErr)
}
