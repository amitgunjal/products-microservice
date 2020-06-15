package utils

import (
	"log"
	"products/constants"
	"products/pkg/models"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func StartUp() (*gorm.DB, error) {
	return CreateDBConnection()

}

func CreateDBConnection() (*gorm.DB, error) {
	// Please define your username and password for MySQL.
	dbConnectionURI := LoadEnvConfigs()
	log.Printf("Making DB Connection")
	db, err = gorm.Open(constants.MysqlKey, dbConnectionURI)
	// NOTE: See we’re using = to assign the global var
	// instead of := which would assign it only in this function
	db.SingularTable(true)
	db.AutoMigrate(&models.Product{})
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connected to MySQL at ", dbConnectionURI)
	}
	log.Printf("Connected to database at- {%s} ", dbConnectionURI)
	return db, nil
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
