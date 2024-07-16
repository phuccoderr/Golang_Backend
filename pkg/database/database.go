package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect(databaseURL string) *gorm.DB {
	DB, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return DB
}
