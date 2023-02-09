package controllers

import (
	"html/template"
	"mvc/app/entities"
	"mvc/app/models"
	"net/http"
	"strconv"
)

func Index(response http.ResponseWriter, request *http.Request) {
	var productModel models.ProductModel
	products, _ := productModel.FindAll()
	data := map[string]interface{}{
		"products": products,
	}
	tmp, _ := template.ParseFiles("views/product/index.html")
	tmp.Execute(response, data)
}
func Add(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("views/product/add.html")
	tmp.Execute(response, nil)
}

func ProcessAdd(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var product entities.Product
	product.Name = request.Form.Get("name")
	product.Price, _ = strconv.ParseFloat(request.Form.Get("price"), 64)
	product.Quantity, _ = strconv.ParseInt(request.Form.Get("quantity"), 10, 64)
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
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var productModel models.ProductModel
	product, _ := productModel.Find(id)
	productModel.Find(id)
	data := map[string]interface{}{
		"product": product,
	}
	tmp, _ := template.ParseFiles("views/product/edit.html")
	tmp.Execute(response, data)
}

func Update(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var product entities.Product
	product.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
	product.Name = request.Form.Get("name")
	product.Price, _ = strconv.ParseFloat(request.Form.Get("price"), 64)
	product.Description = request.Form.Get("description")
	var productModel models.ProductModel
	productModel.Update(product)
	http.Redirect(response, request, "/product", http.StatusSeeOther)
}
