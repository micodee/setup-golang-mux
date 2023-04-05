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

	subrouter.HandleFunc("/", controllers.Home).Methods("GET")
	subrouter.HandleFunc("/products", controllers.Index).Methods("GET")
	subrouter.HandleFunc("/product/{id}", controllers.Show).Methods("GET")
	subrouter.HandleFunc("/product", controllers.Create).Methods("POST")
	subrouter.HandleFunc("/product/{id}", controllers.Update).Methods("PUT")
	subrouter.HandleFunc("/product", controllers.Delete).Methods("DELETE")

	// create server port
	port := "8000"
	fmt.Println("server running on port", port)
	http.ListenAndServe("localhost:"+port, router)
}
