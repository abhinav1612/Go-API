package main

import (
  "github.com/gin-gonic/gin"
  //"github.com/gin-gonic/gin/binding"
  //"database/sql"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  _ "github.com/lib/pq"
  "fmt"
  "log"
  "math/rand"
  "time"
  "strconv"
  //"encoding/json"
)

const (
    DB_USER     = "furlenco"
    DB_PASSWORD = ""
    DB_NAME     = "furlenco_development"
)

type Product struct {
    gorm.Model
    Name string `gorm:"type:varchar(100);not null"`
    Searchkey string `gorm:"type:varchar(100);unique"`
    MonthlyRentalPrice float64
}

type ProductForm struct {
  Name string `json:"name" binding:"required"`
  Price float64 `json:"price" binding:"required"`
}

func main() {

  // Initializing APP
  app := gin.Default()

  // Initializing DB
  //db, err := gorm.Open("postgres", "host=localhost user=furlenco dbname=furlenco_development sslmode=disable")
  //db.CreateTable(&User{})
  //initDB()
  app.GET("/test",test)
  app.GET("/products",index)
  app.POST("/products",create)
  app.GET("/products/:id",productById)
  app.PUT("/products/:id",update)

  app.Run(":8080")
}

func index(c *gin.Context) {
  db, err := gorm.Open("postgres", "host=localhost user=furlenco dbname=furlenco_development sslmode=disable")
  checkErr(err)
  db.LogMode(true)
  var products [] Product
  db.Find(&products)
  content := gin.H{}
  for k,v := range products {
    content[strconv.Itoa(k)] = v
  }
  c.JSON(200,content)
}

func create(c *gin.Context){
  var json ProductForm
  c.BindJSON(&json)
  db, err := gorm.Open("postgres", "host=localhost user=furlenco dbname=furlenco_development sslmode=disable")
  checkErr(err)
  db.LogMode(true)
  newProduct := createProduct(json.Name,json.Price)
  fmt.Printf("Product to store: %v\n", newProduct)
  db.Create(&newProduct)
  defer db.Close()

  content := gin.H{
           "result": "Success",
           "id": newProduct.ID,
           "name": newProduct.Name,
           "price": newProduct.MonthlyRentalPrice,
       }
  c.JSON(201, content)
}

func productById(c *gin.Context){
  db, err := gorm.Open("postgres", "host=localhost user=furlenco dbname=furlenco_development sslmode=disable")
  checkErr(err)
  db.LogMode(true)
  var product Product
  product_id := c.Params.ByName("id")
  db.First(&product,product_id)

  content := gin.H{
    "name" : product.Name,
    "price" : product.MonthlyRentalPrice,
  }

  c.JSON(200,content)
}

func update(c *gin.Context){
  var json ProductForm
  c.BindJSON(&json)
  db, err := gorm.Open("postgres", "host=localhost user=furlenco dbname=furlenco_development sslmode=disable")
  checkErr(err)
  db.LogMode(true)
  var product Product
  product_id := c.Params.ByName("id")
  db.First(&product,product_id)

  product.Name = json.Name
  product.MonthlyRentalPrice = json.Price

  db.Save(&product)

  content := gin.H{
    "name" : product.Name,
    "price" : product.MonthlyRentalPrice,
  }

  c.JSON(200,content)
}

func test(c *gin.Context) {
  content := gin.H{"msg":"Hello World"}
  c.JSON(200,content)
}

func createProduct(name string, price float64) Product {

  product := Product{
      //CreatedAt:    time.Now().UnixNano(),
      //UpdatedAt:    time.Now().UnixNano(),
      Name:         name,
      Searchkey:    RandomString(14),
      MonthlyRentalPrice: price ,
  }
  return product
}

func initDB() {

  db, err := gorm.Open("postgres", "host=localhost user=furlenco dbname=furlenco_development sslmode=disable")
  checkErr(err)
  db.LogMode(true)
  db.Debug().AutoMigrate(&Product{})
  // defer db.Close()
  defer db.Close()
}

// RandomString Generator
func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func checkErr(err error) {
    if err != nil {
        log.Fatalln(err)
    }
}
