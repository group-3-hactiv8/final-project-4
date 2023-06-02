package dto

import (
	"final-project-4/models"
	"final-project-4/pkg/errs"

	"github.com/asaskevich/govalidator"
)

type NewTransactionRequest struct {
	ProductId uint `json:"product_id" valid:"required~Product ID is required"`
	Quantity  uint `json:"quantity" valid:"required~Quantity is required"`
}

func (t *NewTransactionRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(t)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

func (t *NewTransactionRequest) NewTransactionRequestToModel() *models.TransactionHistory {
	return &models.TransactionHistory{
		ProductId: t.ProductId,
		Quantity:  t.Quantity,
	}
}

type TransactionBillResponse struct {
	TotalPrice   uint   `json:"total_price"`
	Quantity     uint   `json:"quantity"`
	ProductTitle string `json:"product_title"`
}

type NewTransactionResponse struct {
	Message         string                  `json:"message"`
	TransactionBill TransactionBillResponse `json:"transaction_bill"`
}
