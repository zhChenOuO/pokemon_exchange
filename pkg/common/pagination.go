package common

import (
	"gorm.io/gorm"
)

const globalDefaultPerPage = 20

// SetTotalCountAndPage 用來計算總數和分頁
func (p *Pagination) SetTotalCountAndPage(total int64) {
	p.CheckOrSetDefault()
	p.TotalCount = total

	quotient := p.TotalCount / p.PerPage
	remainder := p.TotalCount % p.PerPage
	if remainder > 0 {
		quotient++
	}
	p.TotalPage = quotient
}

// CheckOrSetDefault 檢查Page值若未設置則設置預設值
func (p *Pagination) CheckOrSetDefault(params ...int64) *Pagination {
	var defaultPerPage int64
	if len(params) >= 1 {
		defaultPerPage = params[0]
	}

	if defaultPerPage <= 0 {
		defaultPerPage = globalDefaultPerPage
	}

	if p.Page == 0 {
		p.Page = 1
	}
	if p.PerPage == 0 {
		p.PerPage = defaultPerPage
	}
	return p
}

// LimitAndOffset return limit and offset
func (p *Pagination) LimitAndOffset(db *gorm.DB) *gorm.DB {
	if p.PerPage != 0 || p.Offset() != 0 {
		db = db.Limit(int(p.PerPage)).Offset(int(p.Offset()))
	}
	return db
}

// Offset 計算 offset 的值
func (p *Pagination) Offset() int64 {
	if p.Page <= 0 {
		return 0
	}
	return (p.Page - 1) * p.PerPage
}
