package application

import "mundo/src/products/domain"

type GetAllProduct struct {
	db domain.IProduct
}

func NewGetAllProduct(db domain.IProduct) *GetAllProduct {
	return &GetAllProduct{db: db}
}

func (uc *GetAllProduct) Run() ([]domain.Product, error) {
	return uc.db.GetAll()
}