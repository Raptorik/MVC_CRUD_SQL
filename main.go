package main

import (
	"mvc/app/controllers/productcontroller"
	"net/http"
)

func main() {
	http.HandleFunc("/", productcontroller.Index)
	http.HandleFunc("/product", productcontroller.Index)
	http.HandleFunc("/product/index", productcontroller.Index)
	http.HandleFunc("/product/add", productcontroller.Add)
	http.HandleFunc("/product/processadd", productcontroller.ProcessAdd)
	http.ListenAndServe(":8080", nil)
}
