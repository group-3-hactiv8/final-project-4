package transaction_history_repository

import (
	"final-project-4/models"
	"final-project-4/pkg/errs"
)

type TransactionHistoryRepository interface {
	CreateTransaction(*models.TransactionHistory) (*models.TransactionHistory, errs.MessageErr)
	GetTransactionsByUserID(userID uint) ([]models.TransactionHistory, errs.MessageErr)

	GetUserTransactions() ([]models.TransactionHistory, errs.MessageErr)

}
