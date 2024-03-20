package service

import (
	"blogs/model"
	"blogs/repo"
	"fmt"
)

type BlogPostService struct {
	BlogPostRepo *repo.BlogPostRepository
}

func (service *BlogPostService) Create(blogPost *model.BlogPost) (*model.BlogPost, error) {
	createBlogPost, err2 := service.BlogPostRepo.Create(blogPost)

	if err2 != nil {
		return nil, fmt.Errorf("couldn't create")
	}

	return createBlogPost, nil
}

func (service *BlogPostService) FindAllBlogPosts() ([]model.BlogPost, error) {
	return service.BlogPostRepo.FindAll()
}

// func (service *BlogPostService) FindBlogPost(ID string) (*model.BlogPost, error) {
// 	blogPost, err := service.BlogPostRepo.FindById(ID)
// 	if err != nil {
// 		return nil, fmt.Errorf(fmt.Sprintf("menu item with ID %s not found", ID))
// 	}
// 	return &blogPost, nil
// }

func (service *BlogPostService) FindBlogPost(id int) (*model.BlogPost, error) {
	blogPost, err := service.BlogPostRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &blogPost, nil

}
