package main

import (
	"blogs/handler"
	"blogs/repo"
	"blogs/routing"
	"blogs/service"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connection_url := "user=postgres password=super dbname=SOA port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(connection_url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return database
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	blogRepo := &repo.BlogRepository{DatabaseConnection: database}
	blogService := &service.BlogService{BlogRepo: blogRepo}
	blogHandler := &handler.BlogHandler{BlogService: blogService}

	router := routing.SetupRoutes(blogHandler)

	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8081", router))
}
