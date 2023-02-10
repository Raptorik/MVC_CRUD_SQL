package controllers

import (
	"html/template"
	"mvc/app/entities"
	"mvc/app/models"
	"net/http"
	"strconv"
)

func Index(response http.ResponseWriter, _ *http.Request) {
	var productModel models.ProductModel
	products, err := productModel.FindAll()
	if err != nil {
		return
	}
	data := map[string]interface{}{
		"products": products,
	}
	tmp, err := template.ParseFiles("views/product/index.html")
	if err != nil {
		return
	}
	if err := tmp.Execute(response, data); err != nil {
		return
	}
}
func Add(response http.ResponseWriter, _ *http.Request) {
	tmp, err := template.ParseFiles("views/product/add.html")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmp.Execute(response, nil)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ProcessAdd(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Error(response, "Failed to parse form data", http.StatusBadRequest)
		return
	}
	var product entities.Product
	product.Name = request.Form.Get("name")
	price, err := strconv.ParseFloat(request.Form.Get("price"), 64)
	if err != nil {
		http.Error(response, "Invalid price value", http.StatusBadRequest)
		return
	}
	product.Price = price
	quantity, err := strconv.ParseInt(request.Form.Get("quantity"), 10, 64)
	if err != nil {
		http.Error(response, "Invalid quantity value", http.StatusBadRequest)
		return
	}
	product.Quantity = quantity
	product.Description = request.Form.Get("description")
	var productModel models.ProductModel
	productModel.Create(&product)
	http.Redirect(response, request, "/product", http.StatusSeeOther)
}
func Delete(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var productModel models.ProductModel
	productModel.Delete(id)
	http.Redirect(response, request, "/product", http.StatusSeeOther)
}

func Edit(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, err := strconv.ParseInt(query.Get("id"), 10, 64)
	if err != nil {
		// handle the error, for example by writing an error response to the client
		http.Error(response, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	var productModel models.ProductModel
	product, err := productModel.Find(id)
	if err != nil {
		// handle the error, for example by writing an error response to the client
		http.Error(response, "Error finding product", http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{
		"product": product,
	}
	tmp, err := template.ParseFiles("views/product/edit.html")
	if err != nil {
		// handle the error, for example by writing an error response to the client
		http.Error(response, "Error parsing template", http.StatusInternalServerError)
		return
	}
	err = tmp.Execute(response, data)
	if err != nil {
		// handle the error, for example by writing an error response to the client
		http.Error(response, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func Update(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Error(response, "Failed to parse form data", http.StatusBadRequest)
		return
	}
	var product entities.Product
	id, err := strconv.ParseInt(request.Form.Get("id"), 10, 64)
	if err != nil {
		// handle the error, for example by writing an error response to the client
		http.Error(response, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	product.Id = id
	product.Name = request.Form.Get("name")
	price, err := strconv.ParseFloat(request.Form.Get("price"), 64)
	if err != nil {
		// handle the error, for example by writing an error response to the client
		http.Error(response, "Invalid price parameter", http.StatusBadRequest)
		return
	}
	product.Price = price
	product.Description = request.Form.Get("description")
	var productModel models.ProductModel
	productModel.Update(product)
	if err != nil {
		http.Error(response, "Error updating product", http.StatusInternalServerError)
		return
	}
	http.Redirect(response, request, "/product", http.StatusSeeOther)
}
