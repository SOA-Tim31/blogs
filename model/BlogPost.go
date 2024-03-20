package model

import "time"

// type BlogPostStatus int

// const (
// 	DRAFT BlogPostStatus = iota
// 	PUBLISHED
// 	CLOSED
// 	ACTIVE
// 	FAMOUS
// )

type BlogPost struct {
	ID              int               `json:"id" gorm:"primaryKey;autoIncrement;column:Id"`
	AuthorId        int               `json:"authorId"`
	TourId          int               `json:"tourId"`
	Title           string            `json:"title"`
	Description     string            `json:"description"`
	CreationDate    time.Time         `json:"creationDate"`
	ImageURLs       string            `json:"imageURLs,omitempty"`
	BlogPostComment []BlogPostComment `json:"blogPostComment,omitempty"`
	Status          string            `json:"status"`
}

func (BlogPost) TableName() string {
	return "blog_posts.blog_posts"
}
