package application

import (
	"fmt"

	"products_api.com/p/src/Products/domain"
)

type DeleteProductUC struct {
	db domain.IProduct
}

func NewDeleteProductUC(db domain.IProduct)*DeleteProductUC{
	return&DeleteProductUC{db: db}
}

func(uc *DeleteProductUC)Execute(id int)error{
	fmt.Println("Ejecutando caso de uso")
	return uc.db.DeleteProductMethod(id)
}