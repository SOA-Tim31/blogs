package main

import (
	"blogs/handler"
	"blogs/migration"
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
		log.Fatalf("Failed to connect to database: %v", err)
		return nil
	}

	if err := database.Exec("CREATE SCHEMA IF NOT EXISTS blog_posts").Error; err != nil {
		log.Fatalf("Failed to create schema: %v", err)
		return nil
	}

	if err := migration.AutoMigrate(database); err != nil {
		log.Fatalf("Failed to perform auto migration: %v", err)
		return nil
	}
	return database
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	blogPostRepo := &repo.BlogPostRepository{DatabaseConnection: database}
	blogPostService := &service.BlogPostService{BlogPostRepo: blogPostRepo}
	blogPostHandler := &handler.BlogPostHandler{BlogPostService: blogPostService}

	blogPostCommentRepo := &repo.BlogPostCommentRepository{DatabaseConnection: database}
	blogPostCommentService := &service.BlogPostCommentService{BlogPostCommentRepo: blogPostCommentRepo}
	blogPostCommentHandler := &handler.BlogPostCommentHandler{BlogPostCommentService: blogPostCommentService}

	router := routing.SetupRoutes(blogPostHandler, blogPostCommentHandler)

	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8082", router))
}
