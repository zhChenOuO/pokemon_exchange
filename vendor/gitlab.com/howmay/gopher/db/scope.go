package db

import "gorm.io/gorm"

// SelectDelete 查詢已軟刪
func SelectDelete(db *gorm.DB) *gorm.DB {
	return db.Unscoped()
}

// SelectForUpdate select for update
func SelectForUpdate(db *gorm.DB) *gorm.DB {
	return db.Set("gorm:query_option", "FOR UPDATE")
}

