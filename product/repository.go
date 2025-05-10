package product

import (
	"context"
	"database/sql"
	"errors"
)

const (
	getProducts = `SELECT id, product_name, price 
			FROM 
			product`

	getProductById = `SELECT id, product_name, price 
			FROM product 
			WHERE id = $1`

	createProdut = `INSERT INTO 
			product (product_name, price, description, stock, category)
			VALUES 
			($1, $2, $3, $4, $5) 
			RETURNING id`

	deleteProduct = `DELETE 
			FROM product 
			WHERE id = $1 
			RETURNING id, product_name, price;`

	updateProduct = `UPDATE product 
			SET product_name = $1, price = $2 WHERE id = $3 
			RETURNING id, product_name, price;`
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type repository struct {
	sqlDb *sql.DB
}

type Repository interface {
	GetProduct(ctx context.Context) (Product, error)
	NewProduct(ctx context.Context, product *Product) (int, error)
}

func (r *repository) GetProducts(ctx context.Context) (products *Product, err error) {
	products = &Product{} // Initialize the products pointer
	row := r.sqlDb.QueryRowContext(ctx, getProducts)
	sqlErr := row.Scan(
		&products.ID,
		&products.Name,
		&products.Price,
		&products.Description,
		&products.Stock,
		&products.Category)

	if sqlErr != sql.ErrNoRows {
		err = ErrUserNotFound
	}

	return
}

func (r *repository) NewProduct(ctx context.Context, product *Product) (int, error) {
	product = &Product{}
	row := r.sqlDb.QueryRowContext(ctx, createProdut,
		product.Name,
		product.Price,
		product.Description,
		product.Stock,
		product.Category)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
