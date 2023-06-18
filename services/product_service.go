package services

import (
	"final-project-4/dto"
	"final-project-4/pkg/errs"
	"final-project-4/repositories/product_repository"
)

type ProductService interface {
	CreateProduct(payload *dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr)
	GetAllProducts() (*dto.AllProductsResponse, errs.MessageErr)
	UpdateProducts(id uint, payload *dto.NewProductRequest) (*dto.UpdateProductResponse, errs.MessageErr)
	DeleteProduct(id uint) (*dto.DeleteProductResponse, errs.MessageErr)
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


func (p *productService) UpdateProducts(id uint, payload *dto.NewProductRequest) (*dto.UpdateProductResponse, errs.MessageErr)  {
	product, err := p.productRepo.GetProductByIdUpdate(id)

	if id < 1 {
		idError := errs.NewBadRequest("ID value must be positive")
		return nil, idError
	}

	if err != nil {
		return nil, err
	}

	updateNewProduct := payload.NewProductRequestToModel()

	updateProduct, err2 := p.productRepo.UpdateProducts(product, updateNewProduct)

	if err2 != nil {
		return nil, err2
	}

	response := &dto.UpdateProductResponse {
		ID: updateProduct.ID,
		Title:      updateProduct.Title,
		Price:      updateProduct.Price,
		Stock:      updateProduct.Stock,
		CategoryId: updateProduct.CategoryId,
		CreatedAt: updateProduct.CreatedAt,
		UpdatedAt: updateProduct.UpdatedAt,
	}
	return response, nil
}

func (p *productService) DeleteProduct(id uint) (*dto.DeleteProductResponse, errs.MessageErr) {
	product, err := p.productRepo.GetProductByIdUpdate(id)

	if err != nil {
		return nil, err
	}

	if err := p.productRepo.DeleteProducts(product); err != nil {
		return nil, err
	}

	deleteResponse := &dto.DeleteProductResponse{
		Message: "Product has been successfully deleted",
	}
	return deleteResponse, nil
}