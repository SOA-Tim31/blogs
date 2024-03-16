package model

import "time"

type BlogPostStatus int

const (
	DRAFT BlogPostStatus = iota
	PUBLISHED
	CLOSED
	ACTIVE
	FAMOUS
)

type BlogPost struct {
	Id           int            `gorm:"column:id;type:integer" json:"id"`
	AuthorId     int            `json:"authorId"`
	TourId       int            `json:"tourId"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	CreationDate time.Time      `json:"creationDate"`
	Status       BlogPostStatus `json:"status"`
}
