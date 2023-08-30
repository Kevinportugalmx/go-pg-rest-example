package main

import (
	"net/http"

	"abc.com/db"
	"abc.com/handlers"
	"abc.com/models"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
	// database connection
	db.DBConnection()
	// db.DB.Migrator().DropTable(models.User{})
	db.DB.AutoMigrate(models.User{})

	if err := http.ListenAndServe(":5000", handlers.Router()); err != nil {
		panic(err)
	}

}
