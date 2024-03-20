package dto

import "time"

type BlogPostDto struct {
	ID              int       `json:"ID"`
	AuthorId        int       `json:"authorId"`
	TourId          int       `json:"tourId"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	CreationDate    time.Time `json:"creationDate"`
	ImageURLs       string    `json:"imageURLs,omitempty"`
	BlogPostComment []string  `json:"blogPostComments"`
	Status          string    `json:"status"`
}
