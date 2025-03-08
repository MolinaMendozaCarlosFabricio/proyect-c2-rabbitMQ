package domain

type IRequest interface{
	CreateRequestMethod(request Request)error
	AddProductToRequestMethod(id_request int, id_product int, quantity int)error
}