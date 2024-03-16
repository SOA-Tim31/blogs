package dto

import "time"

type BlogPostDto struct {
	Id           int       `gorm:"column:id;type:integer" json:"id"`
	AuthorId     int       `json:"authorId"`
	TourId       int       `json:"tourId"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreationDate time.Time `json:"creationDate"`
	Status       string    `json:"status"`
}
