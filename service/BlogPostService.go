package service

import (
	"blogs/model"
	"blogs/repo"
	"fmt"
)

type BlogPostService struct {
	BlogPostRepo *repo.BlogPostRepository
}

func (service *BlogPostService) Create(appRating *model.BlogPost) (*model.BlogPost, error) {
	createdRating, err2 := service.BlogPostRepo.Create(appRating)

	if err2 != nil {
		return nil, fmt.Errorf("couldn't create")
	}

	return createdRating, nil
}

func (service *BlogPostService) FindAllBlogPosts() ([]model.BlogPost, error) {
	return service.BlogPostRepo.FindAll()
}
