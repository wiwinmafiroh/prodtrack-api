package product_repository

import (
	"prodtrack-api/entity"
	"prodtrack-api/pkg/errs"
)

type ProductRepository interface {
	CreateProduct(productEntity entity.Product) (*entity.Product, errs.ErrorResponse)
	GetAllProducts() ([]*entity.Product, errs.ErrorResponse)
	GetUserProducts(userId uint) ([]*entity.Product, errs.ErrorResponse)
	GetProductById(productId uint) (*entity.Product, errs.ErrorResponse)
	UpdateProductById(productEntity entity.Product) (*entity.Product, errs.ErrorResponse)
	DeleteProductById(productId uint) errs.ErrorResponse
}
