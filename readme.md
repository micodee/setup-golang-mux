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
