package transaction_history_pg

import (
	"final-project-4/models"
	"final-project-4/pkg/errs"
	"final-project-4/repositories/transaction_history_repository"
	"final-project-4/repositories/user_repository"
	"final-project-4/repositories/product_repository"
	"final-project-4/repositories/category_repository"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type transactionHistoryPG struct {
	db *gorm.DB
	productRepo  product_repository.ProductRepository
	userRepo     user_repository.UserRepository
	categoryRepo category_repository.CategoryRepository

}

func NewTransactionHistoryPG(
	db *gorm.DB,
	productRepo product_repository.ProductRepository,
	userRepo user_repository.UserRepository,
	categoryRepo category_repository.CategoryRepository,
) transaction_history_repository.TransactionHistoryRepository {
	return &transactionHistoryPG{db: db, productRepo: productRepo, userRepo :userRepo, categoryRepo:categoryRepo}
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


func (th *transactionHistoryPG) GetTransactionsByUserID(userID uint) ([]models.TransactionHistory, errs.MessageErr) {
	var transactions []models.TransactionHistory
	result := th.db.Where("user_id = ?", userID).Find(&transactions)

	if err := result.Error; err != nil {
		log.Println(err.Error())
		error := errs.NewInternalServerError("Failed to get transaction histories of user with id")
		return nil, error
	}

	return transactions, nil
}


func (th *transactionHistoryPG) GetUserTransactions() ([]models.TransactionHistory, errs.MessageErr) {
	var transaction []models.TransactionHistory
	err := th.db.Find(&transaction).Error

	if err != nil {
		log.Println("Error : ",err.Error())
		error := errs.NewInternalServerError("Failed to  all user transaction")
		return nil, error
	}

	return transaction, nil
}