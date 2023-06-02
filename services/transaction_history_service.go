package services

import (
	"final-project-4/dto"
	"final-project-4/models"
	"final-project-4/pkg/errs"
	"final-project-4/repositories/category_repository"
	"final-project-4/repositories/product_repository"
	"final-project-4/repositories/transaction_history_repository"
	"final-project-4/repositories/user_repository"
)

type TransactionHistoryService interface {
	CreateTransaction(payload *dto.NewTransactionRequest, user_id uint) (*dto.NewTransactionResponse, errs.MessageErr)
}

type transactionHistoryService struct {
	transactionHistoryRepo transaction_history_repository.TransactionHistoryRepository
	productRepo            product_repository.ProductRepository
	userRepo               user_repository.UserRepository
	categoryRepo           category_repository.CategoryRepository
}

func NewTransactionHistoryService(
	transactionHistoryRepo transaction_history_repository.TransactionHistoryRepository,
	productRepo product_repository.ProductRepository,
	userRepo user_repository.UserRepository,
	categoryRepo category_repository.CategoryRepository,
) TransactionHistoryService {
	return &transactionHistoryService{
		transactionHistoryRepo: transactionHistoryRepo,
		productRepo:            productRepo,
		userRepo:               userRepo,
		categoryRepo:           categoryRepo,
	}
}

func (th *transactionHistoryService) CreateTransaction(payload *dto.NewTransactionRequest, user_id uint) (*dto.NewTransactionResponse, errs.MessageErr) {
	newTransaction := payload.NewTransactionRequestToModel()
	newTransaction.UserId = user_id

	// check if product exist or not
	product := &models.Product{}
	product.ID = payload.ProductId

	err := th.productRepo.GetProductByID(product)
	if err != nil {
		return nil, err
	}

	// check if product's stock is more than the quantity
	if product.Stock < newTransaction.Quantity {
		err := errs.NewBadRequest("Not enough product's stock")
		return nil, err
	}

	// calculate total price
	totalPrice := newTransaction.Quantity * product.Price
	newTransaction.TotalPrice = totalPrice

	// check if user's balance is more than the total price
	user := &models.User{}
	user.ID = user_id
	err = th.userRepo.GetUserByID(user)
	if err != nil {
		return nil, err
	}
	// filling user's attributes from query to get user's balance data

	if user.Balance < totalPrice {
		err := errs.NewBadRequest("Not enough balance, please top up")
		return nil, err
	}

	createdTransactionHistory, err := th.transactionHistoryRepo.CreateTransaction(newTransaction)
	if err != nil {
		return nil, err
	}

	// update product's stock (subtract with quantity)
	product.Stock -= createdTransactionHistory.Quantity
	if err := th.productRepo.UpdateStock(product); err != nil {
		return nil, err
	}

	// update user's balance (substract with total price)
	user.Balance -= totalPrice
	if _, err := th.userRepo.UpdateUser(user); err != nil {
		return nil, err
	}

	// update category's sold_product_amount (add with quantity)
	category := &models.Category{}
	category.ID = product.CategoryId
	if err := th.categoryRepo.GetCategoryByID(category); err != nil {
		return nil, err // category not found (gamungkin sih sebenarnya)
	}
	// all category's attributes are filled.
	category.SoldProductAmount += createdTransactionHistory.Quantity
	if err := th.categoryRepo.UpdateSoldProductAmount(category); err != nil {
		return nil, err
	}

	txBillResp := dto.TransactionBillResponse{
		TotalPrice:   createdTransactionHistory.TotalPrice,
		Quantity:     createdTransactionHistory.Quantity,
		ProductTitle: product.Title,
	}

	response := &dto.NewTransactionResponse{
		Message:         "You have successfully purchased the product",
		TransactionBill: txBillResp,
	}

	return response, nil
}
