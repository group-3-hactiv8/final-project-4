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
	GetTransactionsByUserID(userID uint) ([]dto.GetTransactionsByUserIDResponse, errs.MessageErr)
	GetUserTransactions() ([]dto.GetUserTransactionsResponse, errs.MessageErr)
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
	product.ID = uint(payload.ProductId) // Mengubah nilai ProductId menjadi uint

	product, err := th.productRepo.GetProductByID(uint(product.ID))
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

	user, err = th.userRepo.GetUserByID(uint(user.ID))
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

func (th *transactionHistoryService) GetTransactionsByUserID(user_id uint) ([]dto.GetTransactionsByUserIDResponse, errs.MessageErr) {
	transactions, err := th.transactionHistoryRepo.GetTransactionsByUserID(user_id)
	if err != nil {
		return nil, err
	}

	response := []dto.GetTransactionsByUserIDResponse{}
	for _, transaction := range transactions {
		product, err := th.productRepo.GetProductByID(transaction.ProductId)
		if err != nil {
			return nil, err
		}

		response = append(response, dto.GetTransactionsByUserIDResponse{
			ID:         transaction.ID,
			ProductID:  transaction.ProductId,
			UserID:     transaction.UserId,
			Quantity:   transaction.Quantity,
			TotalPrice: transaction.TotalPrice,
			Product: dto.ProductDataWithCategoryIDAndIntegerPrice{
				ID:         product.ID,
				Title:      product.Title,
				Price:      product.Price,
				Stock:      product.Stock,
				CategoryID: product.CategoryId,
				CreatedAt:  product.CreatedAt,
				UpdatedAt:  product.UpdatedAt,
			},
		})
	}

	return response, nil
}

func (th *transactionHistoryService) GetUserTransactions() ([]dto.GetUserTransactionsResponse, errs.MessageErr) {
	transactions, err := th.transactionHistoryRepo.GetUserTransactions()
	if err != nil {
		return nil, err
	}

	response := []dto.GetUserTransactionsResponse{}
	for _, transaction := range transactions {
		product, err := th.productRepo.GetProductByID(transaction.ProductId)
		if err != nil {
			return nil, err
		}

		user, errGetUser := th.userRepo.GetUserByID(transaction.UserId)
		if errGetUser != nil {
			return nil, errGetUser
		}

		response = append(response, dto.GetUserTransactionsResponse{
			ID:         transaction.ID,
			ProductID:  transaction.ProductId,
			UserID:     transaction.UserId,
			Quantity:   transaction.Quantity,
			TotalPrice: transaction.TotalPrice,
			Product: dto.ProductDataWithCategoryIDAndIntegerPrice{
				ID:         product.ID,
				Title:      product.Title,
				Price:      product.Price,
				Stock:      product.Stock,
				CategoryID: product.CategoryId,
				CreatedAt:  product.CreatedAt,
				UpdatedAt:  product.UpdatedAt,
			},
			User: dto.UserData{
				ID:        user.ID,
				Email:     user.Email,
				FullName:  user.FullName,
				Balance:   user.Balance,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
		})
	}

	return response, nil
}
