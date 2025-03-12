package domain

import "products_api.com/p/src/Products/domain"

type IRequest interface {
	ValidateRequestMethod(id int) (bool, error)
	UpdateRequestsStatusMethod(id_status int, id_request int) error
	GetAllMyRequestsMethod(id int) ([]Request, error)
	GetOneOfMyRequestsMethod(id_request int)([]Request, error)
	ReduceStockMethod(id int) error
	ReduceStockOfAProductMethod(acquire domain.Acquires)(int, error)
}