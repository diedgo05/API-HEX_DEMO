// responsabilidad que tiene este archivo es crear
package application

import (
	"mundo/src/products/domain"
)

type CreateProduct struct {
	db domain.IProduct
}

func NewCreateProduct(db domain.IProduct) *CreateProduct {
	return &CreateProduct{db: db}
}

func (uc *CreateProduct) Run(product domain.Product) error {
	return uc.db.Save(product)
} 

