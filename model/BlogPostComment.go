package model

import "time"

type BlogPostComment struct {
	ID              int       `json:"id" gorm:"primaryKey;autoIncrement;column:Id"`
	Text            string    `json:"text"`
	UserId          int       `json:"userId"`
	BlogPostID      int       `json:"blogPostID" gorm:"type:int;column:BlogPostID"`
	CreationTime    time.Time `json:"creationDate"`
	LastUpdatedTime time.Time `json:"lastUpdatedTime"`
}

func (BlogPostComment) TableName() string {
	return "blog_posts.blog_post_comments"
}
