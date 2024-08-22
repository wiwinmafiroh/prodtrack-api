package handler

import (
	"prodtrack-api/dto"
	"prodtrack-api/entity"
	"prodtrack-api/pkg/errs"
	"prodtrack-api/pkg/helpers"
	"prodtrack-api/service"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) productHandler {
	return productHandler{
		productService: productService,
	}
}

// CreateProduct godoc
// @ID create-product
// @Summary Create a new product
// @Description Create a new product by providing the necessary details. Authentication is required using a bearer token. Ensure all required fields are included in the request body.
// @Tags Products
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token for authentication. Obtain this token by logging in." default(Bearer <token>)
// @Param Product body dto.ProductRequest true "Product creation data. Ensure all required fields are included."
// @Success 201 {object} dto.ProductCreatedResponse
// @Router /products [post]
func (p *productHandler) CreateProduct(ctx *gin.Context) {
	err := helpers.CheckContentType(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode(), err)
		return
	}

	var productRequest dto.ProductRequest

	if err := ctx.ShouldBindJSON(&productRequest); err != nil {
		errBindJSON := errs.NewUnprocessableEntityError("Invalid request body")

		ctx.AbortWithStatusJSON(errBindJSON.StatusCode(), errBindJSON)
		return
	}

	userID := ctx.MustGet("userData").(entity.User).Id

	result, err := p.productService.CreateProduct(userID, productRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

// GetProducts godoc
// @ID get-products
// @Summary Retrieve a list of products
// @Description Retrieve a list of products associated with the authenticated user. Authentication is required using a bearer token. Admins will receive a list of all products, while regular users will only receive the products they have added.
// @Tags Products
// @Produce json
// @Param Authorization header string true "Bearer token for authentication. Obtain this token by logging in." default(Bearer <token>)
// @Success 200 {object} dto.ProductsRetrievedResponse
// @Router /products [get]
func (p *productHandler) GetProducts(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(entity.User)

	result, err := p.productService.GetProducts(userData.Id, string(userData.Role))
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

// GetProductById godoc
// @ID get-product-by-id
// @Summary Retrieve a product by ID
// @Description Retrieve the details of a specific product by its ID. Admins can retrieve any product, while regular users can only retrieve the products they have added. Authentication is required using a bearer token.
// @Tags Products
// @Produce json
// @Param Authorization header string true "Bearer token for authentication. Obtain this token by logging in." default(Bearer <token>)
// @Param productId path int true "ID of the product to be retrieved"
// @Success 200 {object} dto.ProductRetrievedResponse
// @Router /products/{productId} [get]
func (p *productHandler) GetProductById(ctx *gin.Context) {
	productId, err := helpers.GetParamId(ctx, "productId")
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode(), err)
		return
	}

	result, err := p.productService.GetProductById(productId)
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

// UpdateProductById godoc
// @ID update-product-by-id
// @Summary Update a product by ID
// @Description Update the details of a specific product by its ID. Authentication is required using a bearer token. Admins can update any product, while regular users can only update the products they have added.
// @Tags Products
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token for authentication. Obtain this token by logging in." default(Bearer <token>)
// @Param productId path int true "ID of the product to be updated"
// @Param Product body dto.ProductUpdateRequest true "Product update data. Ensure the request body contains all required fields for the update."
// @Success 200 {object} dto.ProductUpdatedResponse
// @Router /products/{productId} [put]
func (p *productHandler) UpdateProductById(ctx *gin.Context) {
	productId, err := helpers.GetParamId(ctx, "productId")
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode(), err)
		return
	}

	err = helpers.CheckContentType(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode(), err)
		return
	}

	var productRequest dto.ProductUpdateRequest

	if err := ctx.ShouldBindJSON(&productRequest); err != nil {
		errBindJSON := errs.NewUnprocessableEntityError("Invalid request body")

		ctx.AbortWithStatusJSON(errBindJSON.StatusCode(), errBindJSON)
		return
	}

	result, err := p.productService.UpdateProductById(productId, productRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

// DeleteProductById godoc
// @ID delete-product
// @Summary Delete a product by ID
// @Description Delete a specific product by its ID. Authentication is required using a bearer token. Admins can delete any product, while regular users can only delete the products they have added.
// @Tags Products
// @Produce json
// @Param Authorization header string true "Bearer token for authentication. Obtain this token by logging in." default(Bearer <token>)
// @Param productId path int true "ID of the product to be deleted"
// @Success 200 {object} dto.ProductDeletedResponse
// @Router /products/{productId} [delete]
func (p *productHandler) DeleteProductById(ctx *gin.Context) {
	productId, err := helpers.GetParamId(ctx, "productId")
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode(), err)
		return
	}

	result, err := p.productService.DeleteProductById(productId)
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}
