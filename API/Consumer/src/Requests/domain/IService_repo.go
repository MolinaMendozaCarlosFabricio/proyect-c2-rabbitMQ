package domain

type IService interface{
	ConfirmValidationRequestMethod(id_request int, id_status int)error
}