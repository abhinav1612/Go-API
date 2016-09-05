package main

import {
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  _ "github.com/lib/pq"
  "log"
  "strconv"
}

type ProductForm struct {
  Name string `json:"name" binding:"required"`
  Price float64 `json:"price" binding:"required"`
}

func getAllProducts(c *gin.Context) {
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

func checkErr(err error) {
    if err != nil {
        log.Fatalln(err)
    }
}
