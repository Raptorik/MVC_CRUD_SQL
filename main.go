package main

import (
	"log"
	"mvc/controllers"
	"net/http"
)

func main() {
	log.Println("server started on: http://localhost:3000")
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/product", controllers.Index)
	http.HandleFunc("/product/index", controllers.Index)
	http.HandleFunc("/product/add", controllers.Add)
	http.HandleFunc("/product/processadd", controllers.ProcessAdd)
	http.HandleFunc("/product/delete", controllers.Delete)
	http.HandleFunc("/product/edit", controllers.Edit)
	http.HandleFunc("/product/update", controllers.Update)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return
	}
}
