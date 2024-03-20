package handler

import (
	"blogs/dto"
	"blogs/model"
	"blogs/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BlogPostHandler struct {
	BlogPostService *service.BlogPostService
}

func MapBlogPostDtoToBlogPost(blogPostDto dto.BlogPostDto) model.BlogPost {
	blogPost := model.BlogPost{
		ID:           blogPostDto.ID,
		AuthorId:     blogPostDto.AuthorId,
		TourId:       1,
		Title:        blogPostDto.Title,
		Description:  blogPostDto.Description,
		CreationDate: blogPostDto.CreationDate,
		ImageURLs:    blogPostDto.ImageURLs,
		Status:       blogPostDto.Status,
	}
	return blogPost
}

// func MapToStatus(status string) model.BlogPostStatus {
// 	switch status {
// 	case "DRAFT":
// 		return model.DRAFT
// 	case "PUBLISHED":
// 		return model.PUBLISHED
// 	case "CLOSED":
// 		return model.CLOSED
// 	case "ACTIVE":
// 		return model.ACTIVE
// 	default:
// 		return model.FAMOUS
// 	}
// }

func (handler *BlogPostHandler) Create(writer http.ResponseWriter, req *http.Request) {

	var blogPostDto dto.BlogPostDto
	err := json.NewDecoder(req.Body).Decode(&blogPostDto)
	blogPost := MapBlogPostDtoToBlogPost(blogPostDto)
	if err != nil {
		println("Error while parsing ", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	createdApp, err := handler.BlogPostService.Create(&blogPost)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(createdApp)
}

func (handler *BlogPostHandler) FindAllBlogPosts(writer http.ResponseWriter, req *http.Request) {
	blogPosts, err := handler.BlogPostService.FindAllBlogPosts()

	println("prvi url: " + blogPosts[0].ImageURLs)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var blogPostsDto []dto.BlogPostDto = handler.MapBlogPostToBlogPostDto(blogPosts)
	println("drugi url: " + blogPostsDto[0].ImageURLs)

	json.NewEncoder(writer).Encode(blogPostsDto)
}

func (handler *BlogPostHandler) MapBlogPostToBlogPostDto(blogPosts []model.BlogPost) []dto.BlogPostDto {

	var blogPostDtos []dto.BlogPostDto
	for _, blogPost := range blogPosts {
		blogPostDto := dto.BlogPostDto{
			ID:           blogPost.ID,
			AuthorId:     blogPost.AuthorId,
			TourId:       blogPost.TourId,
			Title:        blogPost.Title,
			Description:  blogPost.Description,
			CreationDate: blogPost.CreationDate,
			ImageURLs:    blogPost.ImageURLs,
			Status:       blogPost.Status,
		}
		blogPostDtos = append(blogPostDtos, blogPostDto)
	}
	return blogPostDtos
}

// func MapToString(status model.BlogPostStatus) string {
// 	switch status {
// 	case model.DRAFT:
// 		return "DRAFT"
// 	case model.PUBLISHED:
// 		return "PUBLISHED"
// 	case model.CLOSED:
// 		return "CLOSED"
// 	case model.ACTIVE:
// 		return "ACTIVE"
// 	default:
// 		return "FAMOUS"
// 	}
// }

// func (handler *BlogPostHandler) GetById(writer http.ResponseWriter, req *http.Request) {
// 	ID := mux.Vars(req)["ID"]
// 	blogPost, err := handler.BlogPostService.FindBlogPost(ID)
// 	writer.Header().Set("Content-Type", "application/json")
// 	if err != nil {
// 		http.Error(writer, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	blogPostDto := handler.MapOneBlogPostToBlogPostDto(blogPost)
// 	writer.WriteHeader(http.StatusOK)
// 	json.NewEncoder(writer).Encode(blogPostDto)
// }

func (handler *BlogPostHandler) GetById(writer http.ResponseWriter, req *http.Request) {
	println("ncujwcnwicnwincw")
	blogPostID := mux.Vars(req)["id"]

	idBlogPost, err := strconv.Atoi(blogPostID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Invalid tour ID"))
		return
	}

	blogPost, err := handler.BlogPostService.FindBlogPost(idBlogPost)
	if err != nil {
		http.Error(writer, "Failed to find tour", http.StatusInternalServerError)
		return
	}

	blogPostDto := handler.MapOneBlogPostToBlogPostDto(blogPost)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(blogPostDto)
}

func (handler *BlogPostHandler) MapOneBlogPostToBlogPostDto(blogPost *model.BlogPost) dto.BlogPostDto {
	blogPostDto := dto.BlogPostDto{
		ID:           blogPost.ID,
		AuthorId:     blogPost.AuthorId,
		TourId:       blogPost.TourId,
		Title:        blogPost.Title,
		Description:  blogPost.Description,
		CreationDate: blogPost.CreationDate,
		ImageURLs:    blogPost.ImageURLs,
		Status:       blogPost.Status,
	}
	return blogPostDto
}
