package model

type Image struct {
	Id         int    `gorm:"column:id;type:integer" json:"id"`
	BlogPostId int    `json:"blogPostId"`
	ImageURLs  string `json:"imageURLs,omitempty"`
}
