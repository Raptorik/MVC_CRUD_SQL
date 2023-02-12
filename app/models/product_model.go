package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"mvc/app/entities"
)

type ProductModel struct {
}

func GetDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbName := "my-mvc"
	dbUser := "user"
	dbPass := "password"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	return
}

func (*ProductModel) FindAll() ([]entities.Product, error) {
	db, err := GetDB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection: %w", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	rows, err := db.Query("SELECT * from `my-mvc`.product_list")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve products from database: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)

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
	db, err := GetDB()
	if err != nil {
		return entities.Product{}, fmt.Errorf("failed to get database connection: %w", err)
	}
	if err2 := db.Close(); err2 != nil {
		err = fmt.Errorf("failed to get database connection")
	}

	row := db.QueryRow("SELECT * from `my-mvc`.product_list WHERE Id = ?", id)
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
	db, err := GetDB()
	if err != nil {
		return 0, fmt.Errorf("failed to get database connection: %w", err)
	}
	if err2 := db.Close(); err2 != nil {
		err = fmt.Errorf("failed to get database connection")
	}

	result, err := db.Exec("INSERT INTO `my-mvc`.product_list (name, price, quantity, description) values(?,?,?,?)",
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
	db, err := GetDB()
	if err != nil {
		return nil
	}
	if err2 := db.Close(); err2 != nil {
		err = fmt.Errorf("failed to get database connection")
	}

	result, _ := db.Exec("UPDATE `my-mvc`.product_list SET name = ?, price = ?, quantity = ?, description = ? WHERE id = ?",
		product.Name, product.Price, product.Quantity, product.Description, product.Id)
	return result
}
func (m *ProductModel) Delete(id int64) (deleted bool, err error) {
	db, err := GetDB()
	if err != nil {
		return false, err
	}

	result, err := db.Exec("DELETE from `my-mvc`.product_list WHERE id = ?", id)
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
