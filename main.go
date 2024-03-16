package main

import (
	"blogs/handler"
	"blogs/model"
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

	database.Exec("DROP TABLE IF EXISTS blog_posts CASCADE")
	database.Exec("DROP TABLE IF EXISTS blog_post_comments CASCADE")
	database.Exec("DROP TABLE IF EXISTS images CASCADE")
	database.AutoMigrate(&model.BlogPost{})
	database.AutoMigrate(&model.BlogPostComment{})
	database.AutoMigrate(&model.Image{})
	database.Exec("INSERT INTO blog_posts (author_id, tour_id, title, description, creation_date, status) VALUES (1, 1, 'Blog1', 'Lep blog 1', '2024-03-02 00:00:00+01', 1)")
	database.Exec("INSERT INTO blog_posts (author_id, tour_id, title, description, creation_date, status) VALUES (2, 2, 'Blog2', 'Lep blog 2', '2024-03-02 00:00:00+01', 1)")
	database.Exec("INSERT INTO blog_posts (author_id, tour_id, title, description, creation_date, status) VALUES (3, 3, 'Blog3', 'Lep blog 3', '2024-03-02 00:00:00+01', 1)")
	database.Exec("INSERT INTO blog_post_comments (text, user_id, creation_time, last_updated_time, blog_post_id) VALUES ('Odlican', 1, '2024-03-02 00:00:00+01', '2024-03-02 00:00:00+01', 3)")
	database.Exec("INSERT INTO blog_post_comments (text, user_id, creation_time, last_updated_time, blog_post_id) VALUES ('Super', 2, '2024-03-02 00:00:00+01', '2024-03-02 00:00:00+01', 3)")
	database.Exec("INSERT INTO blog_post_comments (text, user_id, creation_time, last_updated_time, blog_post_id) VALUES ('Lose', 3, '03-02-2024', '04-02-2024', 3)")
	database.Exec("INSERT INTO images (blog_post_id, image_urls) VALUES (1, 'https://www.forbes.com/advisor/wp-content/uploads/2021/03/traveling-based-on-fare-deals.jpg')")
	database.Exec("INSERT INTO images (blog_post_id, image_urls) VALUES (2, 'https://www.imtilakgroup.com/cdn-cgi/image/format=auto,fit=contain/https://imt-assets.fra1.digitaloceanspaces.com/Safaraq/posts/f6e1ff32786faea5463a21081af936d85Io_66367v.jpg')")
	database.Exec("INSERT INTO images (blog_post_id, image_urls) VALUES (3, 'https://www.bsr.org/images/heroes/bsr-travel-hero..jpg')")

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

	router := routing.SetupRoutes(blogPostHandler)

	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8082", router))
}
