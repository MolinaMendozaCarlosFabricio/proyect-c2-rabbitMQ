package domain

type IService interface{
	SendRequestToVerifyMethod(id_request int, id_product int, quiantity int)error
}