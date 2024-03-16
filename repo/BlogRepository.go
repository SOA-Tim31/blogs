package repo

import "gorm.io/gorm"

type BlogRepository struct {
	DatabaseConnection *gorm.DB
}
