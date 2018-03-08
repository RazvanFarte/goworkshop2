package main

import (
	"goworkshop/model"
	"goworkshop/web"
	"goworkshop/persistence"
	"fmt"
)


func main() {
	_, err := persistence.InitDB()
	if err != nil{
		panic(err)
	}

	//model.Authors = importer.ImportAuthors("importer/authors.json")
	fmt.Printf("Imported authors are: %s\n", model.Authors)
	//model.Books = importer.ImportBooks("importer/books.json")
	//fmt.Printf("Imported books are: %s\n", model.Books)
	web.StartServer()
}
