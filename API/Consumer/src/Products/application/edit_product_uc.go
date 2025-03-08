package application

import (
	"fmt"

	"products_api.com/p/src/Products/domain"
)

type EditProductUC struct {
	db domain.IProduct
}

func NewEditProductUC(db domain.IProduct)*EditProductUC{
	return&EditProductUC{db: db}
}

func(uc *EditProductUC)Execute(id int, name_product string, category string, price float32, stock int)error{
	product := &domain.Product{ID: id, Name_product: name_product, Category: category, Price: price, Stock: stock}
	fmt.Println("Ejecutando caso de uso")
	return uc.db.EditProductMethod(*product)
}