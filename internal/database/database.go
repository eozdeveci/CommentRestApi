package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/labstack/gommon/log"
)

func NewDatabase() (*gorm.DB, error) {
	log.Info("Setting up database connection")

	dbUserName := "postgres"
	dbPassword := "12345"
	dbHost := "localhost"
	dbTable := "FirstRestApi"
	dbPort := "5432"

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUserName, dbTable, dbPassword)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return db, err
	}

	if err := db.DB().Ping(); err != nil {
		return db, err
	}

	return db, nil
}
