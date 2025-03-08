package domain

type IRequest interface{
	CreateRequestMethod(request Request)(int, error)
	AddProductToRequestMethod(id_request int, id_product int, quantity int)error
}