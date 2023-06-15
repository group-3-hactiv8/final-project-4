package transaction_history_pg

import (
	"final-project-4/models"
	"final-project-4/pkg/errs"
	"final-project-4/repositories/category_repository"
	"final-project-4/repositories/product_repository"
	"final-project-4/repositories/transaction_history_repository"
	"final-project-4/repositories/user_repository"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type transactionHistoryPG struct {
	db 			*gorm.DB
	productRepo	product_repository.ProductRepository
	userRepo 	user_repository.UserRepository
	categoryRepo category_repository.CategoryRepository
}

func NewTransactionHistoryPG(db *gorm.DB, productRepo	product_repository.ProductRepository,
	userRepo 	user_repository.UserRepository, categoryRepo category_repository.CategoryRepository,
	) transaction_history_repository.TransactionHistoryRepository {
	return &transactionHistoryPG{db, productRepo, userRepo, categoryRepo}
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


func (t *transactionHistoryPG) GetTransactionsByUserID(userID uint) ([]models.TransactionHistory, errs.MessageErr) {
	var transactions []models.TransactionHistory
	if err := t.db.Find(&transactions, "user_id = ?", userID).Error; err != nil {
		return nil, errs.NewInternalServerError(
			fmt.Sprintf(
				"Failed to get transaction histories of user with id %d",
				userID,
			),
		)
	}

	return transactions, nil
}

func (t *transactionHistoryPG) GetUserTransactions() ([]models.TransactionHistory, errs.MessageErr) {
	var transactions []models.TransactionHistory
	if err := t.db.Find(&transactions).Error; err != nil {
		return nil, errs.NewInternalServerError("Failed to get all transactions")
	}

	return transactions, nil
}
