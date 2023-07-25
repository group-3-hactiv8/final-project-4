package http_handlers

import (
	"final-project-4/dto"
	"final-project-4/models"
	"final-project-4/pkg/errs"
	"final-project-4/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type transactionHistoryHandler struct {
	transactionHistoryService services.TransactionHistoryService
}

func NewTransactionHistoryHandler(transactionHistoryService services.TransactionHistoryService) *transactionHistoryHandler {
	return &transactionHistoryHandler{transactionHistoryService: transactionHistoryService}
}

// CreateTransaction godoc
//
//	@Summary		Create a Transaction
//	@Description	Create a Transaction  by json
//	@Tags			transactions
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.NewTransactionRequest	true	"Create Transaction request body"
//	@Success		201		{object}	dto.NewTransactionResponse
//
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
//
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/transactions [post]
func (th *transactionHistoryHandler) CreateTransaction(ctx *gin.Context) {
	var requestBody dto.NewTransactionRequest

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

	createdTransactionHistory, err3 := th.transactionHistoryService.CreateTransaction(&requestBody, userId)
	if err3 != nil {
		ctx.JSON(err3.StatusCode(), err3)
		return
	}

	ctx.JSON(http.StatusCreated, createdTransactionHistory)
}

// GetTransactionsByUserID godoc
//
//	@Summary		Get user transaction
//	@Description	Get user transaction by json
//	@Tags			transactions
//	@Produce		json
//	@Success		200		{object}	dto.GetTransactionsByUserIDResponse
//
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
//
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/transactions/my-transactions [get]
func (th *transactionHistoryHandler) GetTransactionsByUserID(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*models.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	transactions, err := th.transactionHistoryService.GetTransactionsByUserID(userData.ID)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

// GetUserTransactions godoc
//
//	@Summary		Get user transaction
//	@Description	Get user transaction by json
//	@Tags			transactions
//	@Produce		json
//	@Success		200		{object}	dto.GetUserTransactions
//
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
//
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/transactions/user-transactions [get]
func (th *transactionHistoryHandler) GetUserTransactions(ctx *gin.Context) {
	transactions, err := th.transactionHistoryService.GetUserTransactions()
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}
	ctx.JSON(http.StatusOK, transactions)
}
