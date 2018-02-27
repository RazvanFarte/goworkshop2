package main

import (
	"fmt"
	"goworkshop/importer"
)


func main() {

	fmt.Println(importer.ImportBooks("importer/authors.json"))

}
