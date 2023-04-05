# CRUD Backend Golang Gorilla Mux

## Technology
- Golang
- Gorm
- Mysql

## Install
install all package using GO
```bash
go mod tidy
go mod download
go get
```

## Usage
Local
```bash
go run main.go
```

## Postman Collection
- [Link Collection](https://api.postman.com/collections/24967780-fb14cb40-6e45-4cc8-9af1-d567d0fed858?access_key=PMAT-01GX84PSZ7KWNA9YJMG4KG9XMX)

# Create Project


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
create db gorilla_crud in mysql

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

## CRUD
>controllers/product.go
- get product by id

    ```bash
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
    ```
- create product

    ```bash
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
    ```

    create in postman method POST>body>raw>json
    ```bash
    {
        "name": "Topi",
        "stock": 10,
        "harga": 10000
    }
    ```

- update product
    ```bash
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
    ```
    update in postman method PUT>body>raw>json
    - endpoint error : localhost:8000/api/v1/product/100
        ```bash
        {
            "name": "Topi",
            "stock": 10,
            "harga": 10000
        }
        ```
- delete product

    ```bash
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
    ```
    delete in postman method DELETE>body>raw>json
    - endpoint error : localhost:8000/api/v1/product
        ```bash
        {
            "id": "1000"
        }
        ```

add route endpoint crud file in main.go
>main.go

```bash
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
```