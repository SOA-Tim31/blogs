package migration

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(); err != nil {
		return err
	}
	return nil
}
