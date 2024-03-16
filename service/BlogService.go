package service

import (
	"blogs/repo"
)

type BlogService struct {
	BlogRepo *repo.BlogRepository
}
