package importer

import (
	"io/ioutil"
	"fmt"
	"goworkshop/models"
	"encoding/json"
)

func ImportBooks(filename string) []models.BookDto{
	values, err := ioutil.ReadFile(filename)

	if(err != nil){
		fmt.Println("Error from opening file")
		panic(err)
	}

	var books []models.BookDto
	_ = json.Unmarshal(values, &books)

	return books
}


func ImportAuthors(filename string) []models.AuthorDto{
	values, err := ioutil.ReadFile(filename)

	if(err != nil){
		fmt.Println("Error from opening file")
		panic(err)
	}

	_ = json.Unmarshal(values, &models.Authors)

	return models.Authors
}