package transaction_history_pg

import (
	"final-project-4/models"
	"final-project-4/pkg/errs"
	"final-project-4/repositories/transaction_history_repository"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type transactionHistoryPG struct {
	db *gorm.DB
}

func NewTransactionHistoryPG(db *gorm.DB) transaction_history_repository.TransactionHistoryRepository {
	return &transactionHistoryPG{db: db}
}

func (th *transactionHistoryPG) CreateTransaction(newTx *models.TransactionHistory) (*models.TransactionHistory, errs.MessageErr) {
	if err := th.db.Create(newTx).Error; err != nil {
		log.Println(err.Error())
		message := fmt.Sprintf("Failed to register a new transaction with User ID %d, Product ID %d, and quantity %d", newTx.UserId, newTx.ProductId, newTx.Quantity)
		error := errs.NewInternalServerError(message)
		return nil, error
	}

	return newTx, nil
}
