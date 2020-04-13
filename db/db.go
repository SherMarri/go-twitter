package db

import (
	"log"

	"./migrations"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// CreateDatabaseConnection Utility for creating and returning a db connection
func CreateDatabaseConnection() *gorm.DB {
	log.Println("Creating connection...")
	db, err := gorm.Open("mysql", "vd:asd@/go_tweets?parseTime=true")
	migrations.RunMigrations(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected successfully.")
	return db
}
