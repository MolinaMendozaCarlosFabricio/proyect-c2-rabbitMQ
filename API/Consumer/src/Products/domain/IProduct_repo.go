package domain

type IProduct interface{
	SaveProductMethod(product Product)error
	GetAllProductsMethod()([]Product, error)
	GetProductsOfARequestMethod(id int)([]Product, error)
	EditProductMethod(product Product)error
	DeleteProductMethod(id int)error
}