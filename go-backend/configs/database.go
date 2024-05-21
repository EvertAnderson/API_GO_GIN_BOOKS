package configs

import (
	"log"

	"gorm.io/driver/postgres"
	db_con "gorm.io/gorm"
)

var DB *db_con.DB

func ConnectToDB() {
	var err error
	dsn := "host=localhost user=postgres password=S4postgresql7@ dbname=db_go_books port=5432 sslmode=disable"
	DB, err = db_con.Open(postgres.Open(dsn), &db_con.Config{})

	if err != nil {
		log.Fatal("Failed to connect DB")
	}
}
