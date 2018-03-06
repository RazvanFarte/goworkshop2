package main

import (
	"fmt"
	"goworkshop/importer"
	"goworkshop/model"
	"goworkshop/web"
)

func main() {
	model.Authors = importer.ImportAuthors("importer/authors.json")
	fmt.Printf("Imported authors are: %s\n", model.Authors)
	model.Books = importer.ImportBooks("importer/books.json")
	fmt.Printf("Imported books are: %s\n", model.Books)
	web.StartServer()
}

	fmt.Println("hello world!")
	fmt.Println("My name is Claudia")
	fmt.Println("Razvan Farte was here")
	fmt.Println("my name is Ioan")
	fmt.Println("Hello my name is Alex")
	fmt.Println("And I want to merge my changes")
	fmt.Println("hello my name is alin")
	fmt.Println("Hello my nombre es Bogdan")
	fmt.Println("my name is Ioan")
	fmt.Println("salut lumeee!!!... de la Radu Dragan")
	fmt.Println("Hello my name is Florin!")
	fmt.Println("salut!")
	fmt.Println("Hello, my name is Bogdan!")
	fmt.Println("Hello my name is Eduard")
	fmt.Println("Hello, my name is Andrei")
}
