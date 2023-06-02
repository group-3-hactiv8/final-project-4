package http_handlers

import (
	"final-project-4/dto"
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
