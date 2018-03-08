package web

import (
	"net/http"
	"github.com/gorilla/mux"
	"goworkshop/model"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"goworkshop/persistence"
)

//Demonstrates the basic functionality of private and public modifiers in GO
func Index(w http.ResponseWriter, r *http.Request) {
	helloWorkshop := struct {
		Message        string `json:"message"`
		privateMessage string `json:"privateMessage"`
		NoTagField     string `json:"-"`
	}{
		Message:        "Hello workshop!",
		privateMessage: "Message that does not appear in response :).",
		NoTagField:     "This message won't appear either",
	}
	WriteJson(w, helloWorkshop)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	WriteJson(w, model.Books)
}

func GetBookByUUID(w http.ResponseWriter, r *http.Request) {
	//TODO Check this function by API
	var bookUUID = mux.Vars(r)["uuid"]

	var book model.Book
	err :=  persistence.Connection.Where("uuid = ", bookUUID).Find(&book, 1)

	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	} else {
		WriteJson(w, book)
	}
}

func DeleteBookByUUID(w http.ResponseWriter, r *http.Request) {
	var bookUUID = mux.Vars(r)["uuid"]
	err := persistence.Connection.Delete(&model.Book{ UUID : bookUUID}).Error
	if err != nil {
		fmt.Fprintf(w, "Failed to delete book: %s", err)
	} else {
		WriteJson(w, model.Books)
	}
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &book)
	if err != nil {
		fmt.Fprintf(w, "Failed to create book: %s", err)
	} else {
		model.Books.Add(book)
		persistence.Connection.Create(&book)
		WriteJson(w, book)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &book)
	if err != nil {
		fmt.Fprintf(w, "Failed to unmarshal book: %s", err)
		return
	}

	err = persistence.Connection.Save(&book).Error
	if err != nil {
		fmt.Fprintf(w, "Failed to unmarshal book: %s", err)
		return
	}

	WriteJson(w, book)
}