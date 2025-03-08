package application

import (
	"fmt"

	"products_api.com/p/src/Products/domain"
)

type CreateProductUC struct {
	db domain.IProduct
}

func NewCreateProductUC(db domain.IProduct)*CreateProductUC{
	return&CreateProductUC{db: db}
}

func(uc *CreateProductUC)Execute(name_product string, category string, price float32, stock int)error{
	product := &domain.Product{
		ID: 0, Name_product: name_product, Category: category, Price: price, Stock: stock,
	}
	fmt.Println("Ejecutando caso de uso")
	return uc.db.SaveProductMethod(*product)
}