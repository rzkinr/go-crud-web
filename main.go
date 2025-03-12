package main

import (
	"go-web-native/config"
	"go-web-native/controllers/categoriescontroller"
	"go-web-native/controllers/logincontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectedDB()

	http.HandleFunc("/login", logincontroller.HandlerLogin)

	// 2. Categories
	http.HandleFunc("/categories", categoriescontroller.Index)
	http.HandleFunc("/categories/add", categoriescontroller.Add)
	http.HandleFunc("/categories/detil/", categoriescontroller.Detil)
	http.HandleFunc("/categories/edit", categoriescontroller.Edit)
	http.HandleFunc("/categories/delete", categoriescontroller.Delete)

	log.Println("Starting server on :3306")
	http.ListenAndServe(":3306", nil)
}
