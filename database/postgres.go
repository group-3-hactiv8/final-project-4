package database

import (
	"final-project-4/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	// password = "santoso" // password database local rakha
	dbPort = 5432
	// dbPort = 5433 // port db local rakha
	dbname = "final-project-4"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbname, dbPort,
	)
	dsn := config
	// inject variable db
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	err = db.Debug().AutoMigrate(models.Product{}, models.Category{}, models.User{}, models.TransactionHistory{})

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	log.Println("Successfully connected to database")

}

func GetPostgresInstance() *gorm.DB {
	return db
}
