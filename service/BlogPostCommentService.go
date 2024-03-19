package service

import (
	"blogs/model"
	"blogs/repo"
)

type BlogPostCommentService struct {
	BlogPostCommentRepo *repo.BlogPostCommentRepository
}

// func (service *BlogPostCommentService) AddComment(blogPostComment *model.BlogPostComment) (*model.BlogPostComment, error) {
// 	createdBlogPostComment, err2 := service.BlogPostCommentRepo.AddComment(blogPostComment)

// 	if err2 != nil {
// 		return nil, fmt.Errorf("couldn't create")
// 	}

// 	return createdBlogPostComment, nil
// }

func (service *BlogPostCommentService) AddComment(blogPostComment *model.BlogPostComment) error {
	println("id", blogPostComment.BlogPostID)
	err := service.BlogPostCommentRepo.AddComment(blogPostComment)
	if err != nil {
		return err
	}
	return nil
}
