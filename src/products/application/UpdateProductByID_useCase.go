package application

import "mundo/src/products/domain"

type UpdateProduct struct {
	db domain.IProduct
}

func NewUpdateProductByID(db domain.IProduct) *UpdateProduct {
	return &UpdateProduct{db: db}
}

func (uc *UpdateProduct) Run(id int, product domain.Product) error {
	return uc.db.UpdateByID(id, product)
}