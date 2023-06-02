package http_handlers

import (
	"final-project-4/dto"
	"final-project-4/pkg/errs"
	"final-project-4/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *productHandler {
	return &productHandler{productService: productService}
}

// CreateProduct godoc
//
//	@Summary		Create a product
//	@Description	Create a product by json
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.NewProductRequest	true	"Create product request body"
//	@Success		201		{object}	dto.NewProductResponse
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/products [post]
func (p *productHandler) CreateProduct(ctx *gin.Context) {
	var requestBody dto.NewProductRequest

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

	createdProduct, err3 := p.productService.CreateProduct(&requestBody)
	if err3 != nil {
		ctx.JSON(err3.StatusCode(), err3)
		return
	}

	ctx.JSON(http.StatusCreated, createdProduct)
}

// GetAllProducts godoc
//
//	@Summary		Get all products
//	@Description	Get all products by json
//	@Tags			products
//	@Produce		json
//	@Success		200		{object}	dto.AllProductsResponse
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/products [get]
func (p *productHandler) GetAllProducts(ctx *gin.Context) {
	allProductsResponse, err := p.productService.GetAllProducts()

	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, allProductsResponse)
}
