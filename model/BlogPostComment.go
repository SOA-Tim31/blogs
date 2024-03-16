package model

import "time"

type BlogPostComment struct {
	Id              int       `gorm:"column:id;type:integer" json:"id"`
	Text            string    `json:"text"`
	UserId          int       `json:"userId"`
	BlogPostId      int       `json:"blogPostId"`
	CreationTime    time.Time `json:"creationDate"`
	LastUpdatedTime time.Time `json:"lastUpdatedTime"`
}
