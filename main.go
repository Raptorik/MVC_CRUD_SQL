package main

import (
	"mvc/app/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/product", controllers.Index)
	http.HandleFunc("/product/index", controllers.Index)
	http.HandleFunc("/product/add", controllers.Add)
	http.HandleFunc("/product/processadd", controllers.ProcessAdd)
	http.HandleFunc("/product/delete", controllers.Delete)
	http.HandleFunc("/product/edit", controllers.Edit)
	http.HandleFunc("/product/update", controllers.Update)

	http.ListenAndServe(":3000", nil)
}
