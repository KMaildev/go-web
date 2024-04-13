package main

import (
	"go-web/config"
	"go-web/controllers/categorycontroller"
	"go-web/controllers/homecontroller"
	"go-web/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/", homecontroller.Welcome)

	http.HandleFunc("/category", categorycontroller.Index)
	http.HandleFunc("/category/create", categorycontroller.Create)
	http.HandleFunc("/category/edit", categorycontroller.Edit)
	http.HandleFunc("/category/delete", categorycontroller.Delete)

	// 3. Products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server running on prot 8080")
	http.ListenAndServe(":8080", nil)
}
