package application

import (
	"fmt"

	"products_api.com/p/src/Requests/domain"
	domain_product "products_api.com/p/src/Products/domain"
)

type ReduceStockOfAProductUC struct {
	db domain.IRequest
}

func NewReduceStockOfAProductUC(db domain.IRequest)*ReduceStockOfAProductUC{
	return&ReduceStockOfAProductUC{db: db}
}

func(uc *ReduceStockOfAProductUC)Execute(id_request int, id_product int, quantity int)error{
	fmt.Println("Ejecutando caso de uso")
	acquire := &domain_product.Acquires{Id_request: id_request, Id_product: id_product, Quantity: quantity}
	return uc.db.ReduceStockOfAProductMethod(*acquire)
}