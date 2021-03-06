package persistence

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"goworkshop/model"
)

const (
	UNIQUE_BOOK_TITLE_CONSTRAINT = "title_unique"
)

var Connection *gorm.DB

func InitDB() (*gorm.DB, error) {

	DBInstance, err := gorm.Open("postgres", "host=localhost port=5432 user=dbadmin " +
		"password=dbadmin dbname=workshop_db sslmode=disable")

	if err != nil {
		fmt.Printf("Error while aquiring db connection: %s", err)
		return nil, err
	}
	DBInstance.DB().SetMaxOpenConns(20)

	DBInstance.LogMode(true)

	// Tables should be singular
	DBInstance.SingularTable(true)

	// Migrating the schema
	// This call will only create the new table if it does not exist - or add new columns , it will not modify
	// any data present in the table or remove or modify any columns
	DBInstance.AutoMigrate(model.Author{})
	DBInstance.AutoMigrate(model.Book{})

	// Adding foreign key constraints on the book table
	DBInstance.Table("book").AddForeignKey("author_uuid", "author(uuid)", "RESTRICT", "RESTRICT")

	//add uniqueness on book.title column
	DBInstance.Table("book").AddUniqueIndex(UNIQUE_BOOK_TITLE_CONSTRAINT, "title")

	Connection = DBInstance

	return DBInstance, nil
}
