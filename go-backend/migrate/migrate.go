package main

import (
	"github.com/EvertAnderson/configs"
	"github.com/EvertAnderson/models"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Book{})
}
