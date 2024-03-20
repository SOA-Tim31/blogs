package handler

import (
	"blogs/model"
	"blogs/service"
	"encoding/json"
	"net/http"
)

type BlogPostCommentHandler struct {
	BlogPostCommentService *service.BlogPostCommentService
}

// func (handler *BlogPostCommentHandler) AddComment(writer http.ResponseWriter, req *http.Request) {
// 	// Izvlačenje blogPostid parametra iz rute
// 	vars := mux.Vars(req)
// 	blogPostIDStr, ok := vars["blogPostid"]
// 	if !ok || blogPostIDStr == "" {
// 		println("Missing or empty blogPostid parameter")
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	blogPostID, err := strconv.Atoi(blogPostIDStr)
// 	if err != nil {
// 		println("Error while parsing blogPostID ", err.Error())
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// Dekodiranje JSON tijela zahtjeva u blogPostComment strukturu
// 	var blogPostComment model.BlogPostComment
// 	err = json.NewDecoder(req.Body).Decode(&blogPostComment)
// 	if err != nil {
// 		println("Error while parsing blogPostCommentDto ", err.Error())
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// Postavljanje BlogPostId u blogPostComment objektu
// 	blogPostComment.BlogPostId = blogPostID

// 	// Dodavanje komentara kroz BlogPostCommentService
// 	createdBlogPostComment, err := handler.BlogPostCommentService.AddComment(&blogPostComment)
// 	if err != nil {
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// Slanje odgovora s kreiranim komentarom
// 	writer.WriteHeader(http.StatusOK)
// 	json.NewEncoder(writer).Encode(createdBlogPostComment)
// }

func (h *BlogPostCommentHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	var blogPostComment model.BlogPostComment

	// Dekodiranje JSON tijela zahtjeva u blogPostComment objekt
	err := json.NewDecoder(r.Body).Decode(&blogPostComment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	blogPostComment.BlogPostID = 1

	// Provjera da li blogPostComment objekt sadrži BlogPostID
	if blogPostComment.BlogPostID == 0 {
		http.Error(w, "Missing BlogPostID in request body", http.StatusBadRequest)
		return
	}

	// Ostatak koda
	err = h.BlogPostCommentService.AddComment(&blogPostComment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
