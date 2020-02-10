package database

import (
	"github.com/jinzhu/gorm"
)

// Migrate the whole data model
func Migrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Customer{})

	// onUpdate := "RESTRICT"
	// onDelete := "CASCADE"
	// db.Model(&Course{}).AddForeignKey("music_school_id", "music_schools(id)", onDelete, onUpdate)
	return db
}
