package main

import (
	"fmt"
	"goworkshop/importer"
)


func main() {

	fmt.Println(importer.ImportBooks("importer/authors.json"))

	fmt.Println(importer.ImportAuthors("importer/books.json"))

	
}
