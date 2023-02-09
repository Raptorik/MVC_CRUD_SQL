package models

import (
	"mvc/app/config"
	"mvc/app/entities"
)

type ProductModel struct {
}

func (*ProductModel) FindAll() ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT * from PRODUCT")
		if err2 == nil {
			return nil, err2
		} else {
			var products []entities.Product
			for rows.Next() {
				var product entities.Product
				rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description)
				products = append(products, product)
			}
			return products, nil
		}
	}
}

func (*ProductModel) Create(product *entities.Product) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("INSERT INTO product(name, price, quantity, description) values(?,?,?,?)", product.Name, product.Price, product.Quantity, product.Description)
		if err2 != nil {
			return false
		} else {
			RowsAffected, _ := result.RowsAffected()
			return RowsAffected > 0
		}
	}
}
