package web

import (
	"net/http"
	"goworkshop/model"
	"github.com/gorilla/mux"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"goworkshop/persistence"
)

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	WriteJson(w, model.Authors)
}

func GetAuthorByUUID(w http.ResponseWriter, r *http.Request) {
	authorUUID := mux.Vars(r)["uuid"]
	author, err := persistence.Connection.First(authorUUID)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	} else {
		WriteJson(w, author)
	}
}

func DeleteAuthorByUUID(w http.ResponseWriter, r *http.Request) {
	var authorUUID = mux.Vars(r)["uuid"]
	err := persistence.Connection.Delete(&model.Author{UUID : authorUUID})
	if err != nil {
		fmt.Fprintf(w, "Failed to delete author: %s", err)
	}

	WriteJson(w, model.Authors)

}

func AddAuthor(w http.ResponseWriter, r *http.Request) {
	var author model.Author
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &author)
	if err != nil {
		fmt.Fprintf(w, "Failed to create author: %s", err)
	}

	err = persistence.Connection.Create(&author).Error

	if err != nil {
		fmt.Fprintf(w, "Failed to update author: %s", err)
		return
	}

	WriteJson(w, author)

}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var author model.Author
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &author)
	if err != nil {
		fmt.Fprintf(w, "Failed to unmarshal the body author: %s", err)
		return
	}

	err = persistence.Connection.Save(&author).Error

	if err != nil {
		fmt.Fprintf(w, "Failed to update author: %s", err)
		return
	}

	WriteJson(w, author)
}
