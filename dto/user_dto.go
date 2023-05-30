package dto

import (
	"final-project-4/models"
	"final-project-4/pkg/errs"
	"time"

	"github.com/asaskevich/govalidator"
)

type NewUserRequest struct {
	FullName string `json:"full_name" valid:"required~Your Full Name is required"`
	Email    string `json:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password string `json:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
}

func (u *NewUserRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

func (u *NewUserRequest) UserRequestToModel() *models.User {
	return &models.User{
		FullName: u.FullName,
		Email:    u.Email,
		Password: u.Password,
	}
}

type NewUserResponse struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Balance   uint      `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginUserRequest struct {
	Email    string `json:"email" valid:"required~Your Email is required"`
	Password string `json:"password" valid:"required~Your password is required"`
}

func (u *LoginUserRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

func (u *LoginUserRequest) LoginUserRequestToModel() *models.User {
	return &models.User{
		Email:    u.Email,
		Password: u.Password,
	}
}

type LoginUserResponse struct {
	Token string `json:"token"`
}

type TopupBalanceRequest struct {
	Balance uint `json:"balance"`
}

func (u *TopupBalanceRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	isUnderOrEqualToMillion := govalidator.InRangeInt(u.Balance, 0, 100000000) // kalau 0 tetep true (lower bound)

	if !isUnderOrEqualToMillion {
		return errs.NewUnprocessableEntity("Balance must have non-negative value and less than-or-equal to 100.000.000")
	}

	return nil
}

func (u *TopupBalanceRequest) TopupBalanceRequestToModel() *models.User {
	return &models.User{
		Balance: u.Balance,
	}
}

type TopupBalanceResponse struct {
	Message string `json:"message"`
}
