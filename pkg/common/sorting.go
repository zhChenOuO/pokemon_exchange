package common

import (
	"fmt"

	"gorm.io/gorm"
)

// Sort 依單一欄位單一方向排序
func (s *Sorting) Sort(db *gorm.DB) *gorm.DB {
	if len(s.SortField) != 0 && len(s.SortOrder) != 0 {
		db = db.Order(fmt.Sprintf("%s %s", s.SortField, s.SortOrder))
	}
	return db
}
