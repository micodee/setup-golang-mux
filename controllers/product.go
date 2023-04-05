package controllers

import (
	"encoding/json"
	"net/http"
	"project/models"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
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

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// 10 base, 64 size
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponError(w, http.StatusBadRequest, err.Error())
		return
	}

	var product []models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponError(w, http.StatusNotFound, "Product not found")
			return
		default:
			ResponError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	ResponJson(w, http.StatusOK, product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		ResponError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := models.DB.Create(&product).Error; err != nil {
		ResponError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponJson(w, http.StatusCreated, product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponError(w, http.StatusBadRequest, err.Error())
		return
	}

	var product models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		ResponError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()

	if models.DB.Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		ResponError(w, http.StatusBadRequest, "Unable to update product")
		return
	}

	product.Id = id

	ResponJson(w, http.StatusOK, product)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	input := map[string]string{"id": ""}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		ResponError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	var product models.Product
	if models.DB.Delete(&product, input["id"]).RowsAffected == 0 {
		ResponError(w, http.StatusBadRequest, "input invalid")
		return
	}

	response := map[string]string{"message": "Product success deleted"}
	ResponJson(w, http.StatusOK, response)
}
