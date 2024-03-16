package migration

import (
	"blogs/model"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.BlogPost{}, &model.BlogPostComment{}, &model.Image{}); err != nil {
		return err
	}
	return nil
}
