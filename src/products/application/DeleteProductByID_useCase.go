package application

import "mundo/src/products/domain"

type DeleteProduct struct {
	db domain.IProduct
}

func NewDeleteProductByID(db domain.IProduct) *DeleteProduct {
	return &DeleteProduct{db: db}
}

func (uc *DeleteProduct) Run(id int) error {
	return uc.db.DeleteByID(id)
}