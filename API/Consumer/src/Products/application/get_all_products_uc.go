package application

import (
	"fmt"

	"products_api.com/p/src/Products/domain"
)

type GetAllProductsUC struct {
	db domain.IProduct
}

func NewGetAllProductsUC(db domain.IProduct)*GetAllProductsUC{
	return&GetAllProductsUC{db: db}
}

func(uc *GetAllProductsUC)Execute()([]domain.Product, error){
	fmt.Println("Ejecutando caso de uso")
	return uc.db.GetAllProductsMethod()
}