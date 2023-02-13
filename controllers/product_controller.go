package controllers

import (
	"fmt"
	"html/template"
	"mvc/entities"
	"mvc/models"
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
	tmp, err := template.ParseFiles("index.html")
	if err != nil {
		return
	}
	if err := tmp.Execute(response, data); err != nil {
		return
	}
}
func Add(response http.ResponseWriter, _ *http.Request) {
	tmp, _ := template.ParseFiles("add.html")
	err := tmp.Execute(response, nil)
	if err != nil {
		return
	}
}

func ProcessAdd(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		return
	}
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
	fmt.Println("id", id)
	var productModel models.ProductModel
	productModel.Delete(id)
	http.Redirect(response, request, "/product", http.StatusSeeOther)
}

func Edit(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var productModel models.ProductModel
	product, _ := productModel.Find(id)
	data := map[string]interface{}{
		"product": product,
	}
	tmp, _ := template.ParseFiles("edit.html")
	err := tmp.Execute(response, data)
	if err != nil {
		return
	}
}

func Update(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		return
	}
	var product entities.Product
	product.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
	product.Name = request.Form.Get("name")
	product.Price, _ = strconv.ParseFloat(request.Form.Get("price"), 64)
	product.Quantity, _ = strconv.ParseInt(request.Form.Get("quantity"), 10, 64)
	product.Description = request.Form.Get("description")
	var productModel models.ProductModel
	productModel.Update(product)
	http.Redirect(response, request, "/product", http.StatusSeeOther)
}
