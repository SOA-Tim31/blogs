package routing

import (
	"blogs/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(handler *handler.BlogPostHandler, handler1 *handler.BlogPostCommentHandler) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/blog/create", handler.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/blog/all", handler.FindAllBlogPosts).Methods("GET", "OPTIONS")
	router.HandleFunc("/blog/get/{id}", handler.GetById).Methods("GET", "OPTIONS")
	router.HandleFunc("/blogComment/create", handler1.AddComment).Methods("POST", "OPTIONS")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	return router
}
