package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/models"
)

func ResponJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ResponError(w http.ResponseWriter, code int, message string) {
	ResponJson(w, code, map[string]string{"message": message})
}

func FindProduct(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	if err := models.DB.Find(&products).Error; err != nil {
		ResponError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponJson(w, http.StatusOK, products)
}
func Show(w http.ResponseWriter, r *http.Request) {

}
func Create(w http.ResponseWriter, r *http.Request) {

}
func Update(w http.ResponseWriter, r *http.Request) {

}
func Delete(w http.ResponseWriter, r *http.Request) {

}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
