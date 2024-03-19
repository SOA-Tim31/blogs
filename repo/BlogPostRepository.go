package repo

import (
	"blogs/model"

	"gorm.io/gorm"
)

type BlogPostRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *BlogPostRepository) FindAll() ([]model.BlogPost, error) {
	var blogPosts []model.BlogPost
	dbResult := repo.DatabaseConnection.Find(&blogPosts)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return blogPosts, nil
}

func (repo *BlogPostRepository) Create(blogPost *model.BlogPost) (*model.BlogPost, error) {
	dbResult := repo.DatabaseConnection.Create(blogPost)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return blogPost, nil
}

// func (repo *BlogPostRepository) FindById(id string) (model.BlogPost, error) {
// 	blogPost := model.BlogPost{}
// 	dbResult := repo.DatabaseConnection.First(&blogPost, "ID = ?", id)
// 	if dbResult != nil {
// 		return blogPost, dbResult.Error
// 	}
// 	return blogPost, nil
// }

func (repo *BlogPostRepository) FindById(id int) (model.BlogPost, error) {
	blogPost := model.BlogPost{}

	dbResult := repo.DatabaseConnection.Preload("BlogPostComment").First(&blogPost, id)
	if dbResult != nil {
		return blogPost, dbResult.Error
	}
	return blogPost, nil
}
