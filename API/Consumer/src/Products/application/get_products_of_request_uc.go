package application

import (
	"fmt"

	"products_api.com/p/src/Products/domain"
)

type GetProductsOfARequestUC struct {
	db domain.IProduct
}

func NewGetProductsOfARequestUC(db domain.IProduct)*GetProductsOfARequestUC{
	return&GetProductsOfARequestUC{db: db}
}

func(uc *GetProductsOfARequestUC)Execute(id int)([]domain.Product, error){
	fmt.Println("Ejecutando caso de uso")
	return uc.db.GetProductsOfARequestMethod(id)
}