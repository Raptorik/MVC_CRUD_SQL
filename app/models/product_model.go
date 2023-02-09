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
func (*ProductModel) Find(id int64) (entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return entities.Product{}, err
	} else {
		rows, err2 := db.Query("SELECT * from PRODUCT WHERE id = ?", id)
		if err2 == nil {
			return entities.Product{}, err2
		} else {
			var product entities.Product
			for rows.Next() {
				rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description)
			}
			return product, nil
		}
	}
}

func (*ProductModel) Create(product *entities.Product) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("INSERT INTO product (name, price, quantity, description) values(?,?,?,?)", product.Name, product.Price, product.Quantity, product.Description)
		if err2 != nil {
			return false
		} else {
			RowsAffected, _ := result.RowsAffected()
			return RowsAffected > 0
		}
	}
}

func (*ProductModel) Update(product entities.Product) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE product SET name = ?, price = ?, quantity = ?, description = ? WHERE id = ?", product.Name, product.Price, product.Quantity, product.Description, product.Id)
		if err2 != nil {
			return false
		} else {
			RowsAffected, _ := result.RowsAffected()
			return RowsAffected > 0
		}
	}
}
func (*ProductModel) Delete(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("DELETE from product WHERE id = ?", id)
		if err2 != nil {
			return false
		} else {
			RowsAffected, _ := result.RowsAffected()
			return RowsAffected > 0
		}
	}
}
