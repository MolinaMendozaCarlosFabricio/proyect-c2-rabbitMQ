package infrastructure

import (
	"log"

	"products_api.com/p/src/Products/domain"
	"products_api.com/p/src/core"
)

type ProductRepoMySQL struct {
	Connection core.ConectionMySQL
}

func NewProductRepoMySQL()*ProductRepoMySQL{
	conn := core.MySQLConection()
	if conn.Err != "" {
		log.Fatal("Error al configurar la pool de conexiones ", conn.Err)
	}
	return&ProductRepoMySQL{Connection: *conn}
}

func(r *ProductRepoMySQL)SaveProductMethod(product domain.Product)error{
	query := "INSERT INTO products (name_product, category, price, stock) VALUES (?, ?, ?, ?)"
	_, err := r.Connection.ExecPreparedQuerys(query, product.Name_product, product.Category, product.Price, product.Stock)
	if err != nil {
        log.Fatalf("Error al registrar producto:", err)
    }
    return err
}

func (r *ProductRepoMySQL)GetAllProductsMethod()([]domain.Product, error){
	query := "SELECT id, name_product, category, price, stock FROM products"
	rows, err := r.Connection.FetchRows(query)
	var products []domain.Product
	if err != nil {
        log.Fatalf("Error al obtener productos:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var id int
		var name_product string
		var category string
		var price float32
		var stock int

		if err := rows.Scan(&id, &name_product, &category, &price, &stock); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		product := domain.Product{ID: id, Name_product: name_product, Category: category, Price: price, Stock: stock}
		products = append(products, product)
	}
	return products, err
}

func(r *ProductRepoMySQL)GetProductsOfARequestMethod(id int)([]domain.Product, error){
	query := "SELECT products.id, products.name_product, products.category, products.price, products.stock FROM products INNER JOIN acquires ON acquires.id_product = products.id  WHERE acquires.id_request = ?"
	rows, err := r.Connection.FetchRows(query, id)
	var products []domain.Product
	if err != nil {
        log.Fatalf("Error al obtener productos:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var id int
		var name_product string
		var category string
		var price float32
		var stock int

		if err := rows.Scan(&id, &name_product, &category, &price, &stock); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		product := domain.Product{ID: id, Name_product: name_product, Category: category, Price: price, Stock: stock}
		products = append(products, product)
	}
	return products, err
}

func(r *ProductRepoMySQL)EditProductMethod(product domain.Product)error{
	query := "UPDATE products SET name_product = ?, category = ?, price = ?, stocl = ? WHERE id = ?"
	_, err := r.Connection.ExecPreparedQuerys(
		query, product.Name_product, product.Category, product.Price, product.Stock, product.ID,
	)
	if err != nil {
        log.Fatalf("Error al editar producto: ", err)
    }
    return err
}

func(r *ProductRepoMySQL)DeleteProductMethod(id int)error{
	query := "DELETE FROM products WHERE id = ?"
	_, err := r.Connection.ExecPreparedQuerys(query, id)
	if err != nil {
		log.Fatalf("Error al eliminar producto: ", err)
	}
	return err
}