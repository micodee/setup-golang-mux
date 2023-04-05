# PREPARE

```bash
go mod init project
```

create file main.go

> main.go

- create port

  ```bash
  func main() {
  router := mux.NewRouter()

  port := "8000"
  fmt.Println("server running on port", port)
  http.ListenAndServe("localhost:"+port, router)
  }
  ```

  and then "CTRL + S"

  ```bash
  go mod tidy
  go get
  go mod download
  ```

## MIGRATION

create folder models

- add file product.go

  > product.go

  ```bash
  type Product struct {
  Id    int64   `gorm:"primaryKey" json:"id"`
  Name  string  `gorm:"type:varchar(300)" json:"name"`
  Stock int32   `gorm:"type:int(5)" json:"stock"`
  Harga float64 `gorm:"type:decimal(14,2)" json:"harga"`
  }
  ```

- add file setup.go

  > setup.go

  ```bash
  var DB *gorm.DB

  func ConnDB() {
  db, err := gorm.Open(mysql.Open("admin:admin@tcp(localhost:3306)/gorilla_crud"))
  if err != nil {
  panic(err)
  }

  db.AutoMigrate(&Product{})
  DB = db

  fmt.Println("Connected to Database")
  }
  ```

write file in main.go

> main.go

```bash
func main() {
router := mux.NewRouter()

    models.ConnDB()

port := "8000"
fmt.Println("server running on port", port)
http.ListenAndServe("localhost:"+port, router)
}
```

create subrouter for documentation
```bash
func main() {
router := mux.NewRouter()

models.ConnDB()

    	subrouter := router.PathPrefix("/api/v1").Subrouter()

port := "8000"
fmt.Println("server running on port", port)
http.ListenAndServe("localhost:"+port, router)
}
```

## CONTROLLERS

create folder controllers

- add file product.go
    >product.go
    ```bash
    func FindProducts(w http.ResponseWriter, r *http.Request) {
    var products []models.Product

	if err := models.DB.Find(&products).Error; err != nil {
		fmt.Println(err)
	}

	response, _ := json.Marshal(products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
    }
    ```

    or create a separate response function

    ```bash

    func ResponJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
    }

    func ResponError(w http.ResponseWriter, code int, message string) {
	ResponJson(w, code, map[string]string{"message": message})
    }
    ```
    and then
    ```bash
    
    func FindProduct(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	if err := models.DB.Find(&products).Error; err != nil {
	ResponError(w, http.StatusInternalServerError, err.Error())
	return
	}

	ResponJson(w, http.StatusOK, products)
    }
    ```

write router in main.go

>main.go
```bash
func main() {
router := mux.NewRouter()

models.ConnDB()

subrouter := router.PathPrefix("/api/v1").Subrouter()

    subrouter.HandleFunc("/products", controllers.FindProduct).Methods("GET")

port := "8000"
fmt.Println("server running on port", port)
http.ListenAndServe("localhost:"+port, router)
}
```

running go run . and check in postman : localhost:8000/api/v1/products