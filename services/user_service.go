package services

import (
	"final-project-4/dto"
	"final-project-4/helpers"
	"final-project-4/pkg/errs"
	"final-project-4/repositories/user_repository"
)

type UserService interface {
	RegisterUser(payload *dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
	LoginUser(payload *dto.LoginUserRequest) (*dto.LoginUserResponse, errs.MessageErr)
	UpdateBalance(id uint, payload *dto.UpdateBalanceRequest) (*dto.UpdateBalanceResponse, errs.MessageErr)
}

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) RegisterUser(payload *dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
	newUser := payload.UserRequestToModel()
	newUser.Role = "member"

	createdUser, err := u.userRepo.RegisterUser(newUser)
	if err != nil {
		return nil, err
	}

	response := &dto.NewUserResponse{
		FullName:  createdUser.FullName,
		Email:     createdUser.Email,
		ID:        createdUser.ID,
		CreatedAt: createdUser.CreatedAt,
	}

	return response, nil
}

func (u *userService) LoginUser(payload *dto.LoginUserRequest) (*dto.LoginUserResponse, errs.MessageErr) {
	user := payload.LoginUserRequestToModel()
	passwordFromRequest := user.Password

	err := u.userRepo.GetUserByEmail(user)

	if err != nil {
		return nil, err
	}

	isTheSame := helpers.ComparePass([]byte(user.Password), []byte(passwordFromRequest))
	// harus pake method comparePass ini instead of pake statement Where buat nyari di DB.
	// karena passwordnya disimpan setelah di hash pada function BeforeCreate.

	if !isTheSame {
		err := errs.NewBadRequest("Wrong email/password")
		return nil, err
	}

	token := helpers.GenerateToken(user.ID, user.Email)

	response := &dto.LoginUserResponse{
		Token: token,
	}

	return response, nil
}

func (u *userService) UpdateBalance(id uint, payload *dto.UpdateBalanceRequest) (*dto.UpdateBalanceResponse, errs.MessageErr) {
	userUpdateRequest := payload.UpdateBalanceRequestToModel()

	userUpdateRequest.ID = id

	updatedUser, err := u.userRepo.UpdateBalance(userUpdateRequest)
	if err != nil {
		return nil, err
	}

	response := &dto.UpdateBalanceResponse{
		FullName:  updatedUser.FullName,
		Email:     updatedUser.Email,
		ID:        updatedUser.ID,
		UpdatedAt: updatedUser.UpdatedAt,
	}

	return response, nil
}
