package models

import (
	"database/sql"
	"fmt"

	"mvc/app/config"
	"mvc/app/entities"
)

type ProductModel struct {
}

func (*ProductModel) FindAll() ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection: %w", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * from PRODUCT")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve products from database: %w", err)
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product from database: %w", err)
		}
		products = append(products, product)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("failed to iterate over products: %w", err)
	}
	return products, nil
}

func (*ProductModel) Find(id int64) (entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return entities.Product{}, fmt.Errorf("failed to get database connection: %w", err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT * from PRODUCT WHERE id = ?", id)
	var product entities.Product
	err = row.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.Product{}, fmt.Errorf("product with id %d not found", id)
		}
		return entities.Product{}, fmt.Errorf("failed to retrieve product from database: %w", err)
	}
	return product, nil
}

func (*ProductModel) Create(product *entities.Product) (int64, error) {
	db, err := config.GetDB()
	if err != nil {
		return 0, fmt.Errorf("failed to get database connection: %w", err)
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO product (name, price, quantity, description) values(?,?,?,?)",
		product.Name, product.Price, product.Quantity, product.Description)
	if err != nil {
		return 0, fmt.Errorf("failed to insert product into database: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to insert product into database: %w", err)
	}
	return id, nil
}

func (*ProductModel) Update(product entities.Product) sql.Result {
	db, err := config.GetDB()
	if err != nil {
		return nil
	}
	defer db.Close()

	result, _ := db.Exec("UPDATE product SET name = ?, price = ?, quantity = ?, description = ? WHERE id = ?",
		product.Name, product.Price, product.Quantity, product.Description, product.Id)
	return result
}
func (m *ProductModel) Delete(id int64) (deleted bool, err error) {
	db, err := config.GetDB()
	if err != nil {
		return false, err
	}

	result, err := db.Exec("DELETE from product WHERE id = ?", id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
