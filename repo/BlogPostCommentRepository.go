package repo

import (
	"blogs/model"

	"gorm.io/gorm"
)

type BlogPostCommentRepository struct {
	DatabaseConnection *gorm.DB
}

// func (repo *BlogPostCommentRepository) AddComment(BlogPostComment *model.BlogPostComment) (*model.BlogPostComment, error) {
// 	dbResult := repo.DatabaseConnection.Create(BlogPostComment)
// 	if dbResult.Error != nil {
// 		return nil, dbResult.Error
// 	}
// 	return BlogPostComment, nil
// }

func (repo *BlogPostCommentRepository) AddComment(BlogPostComment *model.BlogPostComment) error {
	dbResult := repo.DatabaseConnection.Create(BlogPostComment)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
