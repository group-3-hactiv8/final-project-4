package http_handlers

import (
	"final-project-4/dto"
	"final-project-4/pkg/errs"
	"final-project-4/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{userService: userService}
}

// RegisterUser godoc
//
//	@Summary		Register a user
//	@Description	Register a user by json
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.NewUserRequest	true	"Create user request body"
//	@Success		201		{object}	dto.NewUserResponse
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/users/register [post]
func (u *userHandler) RegisterUser(ctx *gin.Context) {
	var requestBody dto.NewUserRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	err2 := requestBody.ValidateStruct()
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	createdUser, err3 := u.userService.RegisterUser(&requestBody)
	if err3 != nil {
		ctx.JSON(err3.StatusCode(), err3)
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

// LoginUser godoc
//
//	@Summary		Login
//	@Description	Login by json
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.LoginUserRequest	true	"Login user request body"
//	@Success		200		{object}	dto.LoginUserResponse
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		400		{object}	errs.MessageErrData
//	@Router			/users/login [post]
func (u *userHandler) LoginUser(ctx *gin.Context) {
	var requestBody dto.LoginUserRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	err2 := requestBody.ValidateStruct()
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	token, err := u.userService.LoginUser(&requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, token)
}

// TopupBalance godoc
//
//	@Summary		Add more balance of a user
//	@Description	Add more balance of a user by json
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.TopupBalanceRequest	true	"Add more balance of a user request body"
//	@Success		200		{object}	dto.TopupBalanceResponse
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		400		{object}	errs.MessageErrData
//	@Router			/users/topup [patch]
func (u *userHandler) TopupBalance(ctx *gin.Context) {
	var requestBody dto.TopupBalanceRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	err2 := requestBody.ValidateStruct()
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	// mustget = ambil data dari middleware authentication.
	// Tp hasil returnnya hanya empty interface, jadi harus
	// di cast dulu ke jwt.MapClaims.
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	updatedUserResponse, err2 := u.userService.TopupBalance(userId, &requestBody)
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusOK, updatedUserResponse)
}
