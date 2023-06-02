package services

import (
	"final-project-4/dto"
	"final-project-4/pkg/errs"
	"final-project-4/repositories/product_repository"
)

type ProductService interface {
	CreateProduct(payload *dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr)
	GetAllProducts() (*dto.AllProductsResponse, errs.MessageErr)
}

type productService struct {
	productRepo product_repository.ProductRepository
}

func NewProductService(productRepo product_repository.ProductRepository) ProductService {
	return &productService{productRepo: productRepo}
}

func (p *productService) CreateProduct(payload *dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr) {
	newProduct := payload.NewProductRequestToModel()

	// check if category exist or not too in repo
	createdProduct, err := p.productRepo.CreateProduct(newProduct)
	if err != nil {
		return nil, err
	}

	response := &dto.NewProductResponse{
		ID:         createdProduct.ID,
		Title:      createdProduct.Title,
		Stock:      createdProduct.Stock,
		Price:      createdProduct.Price,
		CategoryId: createdProduct.CategoryId,
		CreatedAt:  createdProduct.CreatedAt,
	}

	return response, nil
}

func (p *productService) GetAllProducts() (*dto.AllProductsResponse, errs.MessageErr) {
	allProducts, totalCount, err := p.productRepo.GetAllProducts()
	if err != nil {
		return nil, err
	}

	var productsListResponse []dto.NewProductResponse = make([]dto.NewProductResponse, 0, totalCount)

	var productsResponse dto.NewProductResponse

	for _, product := range *allProducts {
		productsResponse = dto.NewProductResponse{
			ID:         product.ID,
			Title:      product.Title,
			Stock:      product.Stock,
			Price:      product.Price,
			CategoryId: product.CategoryId,
			CreatedAt:  product.CreatedAt,
		}
		productsListResponse = append(productsListResponse, productsResponse)
	}

	response := &dto.AllProductsResponse{
		Products: productsListResponse,
	}

	return response, nil
}
