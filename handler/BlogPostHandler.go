package handler

import (
	"blogs/dto"
	"blogs/model"
	"blogs/service"
	"encoding/json"
	"net/http"
)

type BlogPostHandler struct {
	BlogPostService *service.BlogPostService
}

func MapBlogPostDtoToBlogPost(blogPostDto dto.BlogPostDto) model.BlogPost {
	blogPost := model.BlogPost{
		Id:           blogPostDto.Id,
		AuthorId:     blogPostDto.AuthorId,
		TourId:       blogPostDto.TourId,
		Title:        blogPostDto.Title,
		Description:  blogPostDto.Description,
		CreationDate: blogPostDto.CreationDate,
		Status:       MapToStatus(blogPostDto.Status),
	}
	return blogPost
}

func MapToStatus(status string) model.BlogPostStatus {
	switch status {
	case "DRAFT":
		return model.DRAFT
	case "PUBLISHED":
		return model.PUBLISHED
	case "CLOSED":
		return model.CLOSED
	case "ACTIVE":
		return model.ACTIVE
	default:
		return model.FAMOUS
	}
}

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
	println("HHBHEEINCENCENCENCO")
	blogPosts, err := handler.BlogPostService.FindAllBlogPosts()

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var blogPostsDto []dto.BlogPostDto = handler.MapBlogPostToBlogPostDto(blogPosts)

	json.NewEncoder(writer).Encode(blogPostsDto)
}

func (handler *BlogPostHandler) MapBlogPostToBlogPostDto(blogPosts []model.BlogPost) []dto.BlogPostDto {

	var blogPostDtos []dto.BlogPostDto
	for _, blogPost := range blogPosts {
		blogPostDto := dto.BlogPostDto{
			Id:           blogPost.Id,
			AuthorId:     blogPost.AuthorId,
			TourId:       blogPost.TourId,
			Title:        blogPost.Title,
			Description:  blogPost.Description,
			CreationDate: blogPost.CreationDate,
			Status:       MapToString(blogPost.Status),
		}
		blogPostDtos = append(blogPostDtos, blogPostDto)
	}
	return blogPostDtos
}

func MapToString(status model.BlogPostStatus) string {
	switch status {
	case model.DRAFT:
		return "DRAFT"
	case model.PUBLISHED:
		return "PUBLISHED"
	case model.CLOSED:
		return "CLOSED"
	case model.ACTIVE:
		return "ACTIVE"
	default:
		return "FAMOUS"
	}
}
