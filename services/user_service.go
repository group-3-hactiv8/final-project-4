package services

import (
	"final-project-4/dto"
	"final-project-4/helpers"
	"final-project-4/models"
	"final-project-4/pkg/errs"
	"final-project-4/repositories/user_repository"
	"fmt"
)

type UserService interface {
	RegisterUser(payload *dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
	LoginUser(payload *dto.LoginUserRequest) (*dto.LoginUserResponse, errs.MessageErr)
	TopupBalance(id uint, payload *dto.TopupBalanceRequest) (*dto.TopupBalanceResponse, errs.MessageErr)
}

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) RegisterUser(payload *dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
	newUser := payload.UserRequestToModel()
	newUser.Role = "customer"
	newUser.Balance = 0

	createdUser, err := u.userRepo.RegisterUser(newUser)
	if err != nil {
		return nil, err
	}

	response := &dto.NewUserResponse{
		ID:        createdUser.ID,
		FullName:  createdUser.FullName,
		Email:     createdUser.Email,
		Password:  createdUser.Password,
		Balance:   createdUser.Balance,
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

func (u *userService) TopupBalance(id uint, payload *dto.TopupBalanceRequest) (*dto.TopupBalanceResponse, errs.MessageErr) {
	initialUser := &models.User{}
	initialUser.ID = id

	_, err := u.userRepo.GetUserByID(initialUser.ID)
	if err != nil {
		return nil, err
	}

	initialUser.Balance += payload.Balance

	updatedUser, err := u.userRepo.UpdateUser(initialUser)
	if err != nil {
		return nil, err
	}

	message := fmt.Sprintf("Your balance has been successfully updated to Rp %d", updatedUser.Balance)
	response := &dto.TopupBalanceResponse{
		Message: message,
	}

	return response, nil
}