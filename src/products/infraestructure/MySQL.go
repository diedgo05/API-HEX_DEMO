// Tecnologia y vamos a implementar los metodos de la interfaz
package infraestructure

import (
	"database/sql"
	"fmt"
	"mundo/src/products/domain"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (mysql *MySQL) Save(product domain.Product) error {
	query:= "INSERT INTO products (name, price) VALUES (?,?)"
	_, err := mysql.db.Exec(query, product.GetName(), product.GetPrice())

	if err != nil {
		return err
	}

	fmt.Println("Producto Creado")
	return nil
}

func (mysql *MySQL) GetAll() ([]domain.Product, error) {
	query := "SELECT * FROM products"
	rows, err := mysql.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var products []domain.Product

	for rows.Next() {
		var prod domain.Product
		var idprod int
		var name string
		var price float32
		if err := rows.Scan(&idprod, &name, &price); err != nil {
			fmt.Println("error al mostrar la fila: %w", err)
		}
		prod = *domain.NewProduct(name, price)
		prod.SetID(int32(idprod))
		products = append(products, prod)
		fmt.Printf("ID: %d, Name: %s, Price: %f\n", idprod,name,price)
	}

	fmt.Println("Lista de productos")
	return products, nil
}

func (mysql *MySQL) DeleteByID(id int) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := mysql.db.Exec(query, id)
	fmt.Println("PRODUCTO ELIMINADO")
	return err 
}

func (mysql *MySQL) UpdateByID(id int, product domain.Product) error {
	query := "UPDATE products SET name = ?, price = ? WHERE id = ?"
	_, err := mysql.db.Exec(query, product.GetName(), product.GetPrice(), id)
	fmt.Println("PRODUCTO ACTUALIZADO")
	return err
}