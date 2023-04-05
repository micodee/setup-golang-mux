package main

import (
	"fmt"
	"net/http"
	"project/controllers"
	"project/models"

	"github.com/gorilla/mux"
)

func main() {
	// framework gorilla mux
	router := mux.NewRouter()
	
	models.ConnDB()

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	subrouter.HandleFunc("/products", controllers.FindProduct).Methods("GET")
	subrouter.HandleFunc("/product/{id}", controllers.GetProduct).Methods("GET")
	subrouter.HandleFunc("/product", controllers.CreateProduct).Methods("POST")
	subrouter.HandleFunc("/product/{id}", controllers.Update).Methods("PUT")
	subrouter.HandleFunc("/product", controllers.Delete).Methods("DELETE")

	// create server port
	port := "8000"
	fmt.Println("server running on port", port)
	http.ListenAndServe("localhost:"+port, router)
}
