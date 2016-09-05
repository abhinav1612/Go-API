package dbConfig

import (
  "database/sql"
  _ "github.com/lib/pq"
  "log"
  "fmt"
)

const (
    DB_USER     = "furlenco"
    DB_PASSWORD = ""
    DB_NAME     = "furlenco_development"
)

func InitDB(){
  //dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable",
  //   DB_USER, DB_PASSWORD, DB_NAME)
  dbinfo := fmt.Sprintf("user=furlenco dbname=furlenco_development sslmode=disable")
  db, err := sql.Open("postgres", dbinfo)
  checkErr(err)
  defer db.Close()

  fmt.Println("# Querying")
  rows, err := db.Query("SELECT * FROM torus_products")
  checkErr(err)

  for rows.Next() {
        var uid int
        var monthly_rental_price float32
        var name string
        var search_key string
        err = rows.Scan(&uid, &monthly_rental_price, &name, &search_key)
        checkErr(err)
        fmt.Println("uid | monthly_rental_price | name | search_key ")
        fmt.Printf("%3v | %8v | %6v | %6v\n", uid, monthly_rental_price, name, search_key)
    }
}

func checkErr(err error) {
    if err != nil {
        log.Fatalln(err)
    }
}
