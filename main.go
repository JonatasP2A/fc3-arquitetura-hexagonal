package main

import (
	"database/sql"

	db2 "github.com/JonatasP2A/go-hexagonal/adapters/db"
	"github.com/JonatasP2A/go-hexagonal/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	defer db.Close()
	
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product Example", 30)

	productService.Enable(product)
}