package common

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

var zeroTime = time.Time{}

func (where *BaseWhere) Where(db *gorm.DB) *gorm.DB {
	if where.IDs != nil && len(where.IDs) != 0 {
		db = db.Where("id IN (?)", where.IDs)
	}
	if where.SearchIn != "" && where.Keyword != "" {
		fields := strings.Split(where.SearchIn, ",")
		for _, field := range fields {
			field := field
			db = db.Where(fmt.Sprintf("%s like ?", field), "%"+where.Keyword+"%")
		}
	}
	if where.CreatedAtLt != nil && *where.CreatedAtLt != zeroTime {
		sec := (*where.CreatedAtLt).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("created_at < ?", start)
	}
	if where.CreatedAtLte != nil && *where.CreatedAtLte != zeroTime {
		sec := (*where.CreatedAtLte).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("created_at <= ?", start)
	}
	if where.CreatedAtGt != nil && *where.CreatedAtGt != zeroTime {
		sec := (*where.CreatedAtGt).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("created_at > ?", start)
	}
	if where.CreatedAtGte != nil && *where.CreatedAtGte != zeroTime {
		sec := (*where.CreatedAtGte).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("created_at >= ?", start)
	}
	if where.UpdatedAtLt != nil && *where.UpdatedAtLt != zeroTime {
		sec := (*where.UpdatedAtLt).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("updated_at < ?", start)
	}
	if where.UpdatedAtLte != nil && *where.UpdatedAtLte != zeroTime {
		sec := (*where.UpdatedAtLte).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("updated_at <= ?", start)
	}
	if where.UpdatedAtGt != nil && *where.UpdatedAtGt != zeroTime {
		sec := (*where.UpdatedAtGt).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("updated_at > ?", start)
	}
	if where.UpdatedAtGte != nil && *where.UpdatedAtGte != zeroTime {
		sec := (*where.UpdatedAtGte).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("updated_at >= ?", start)
	}
	if where.CreatorID > 0 {
		db = db.Where("creator_id = ?", where.CreatorID)
	}
	if where.CreatorName != "" {
		db = db.Where("creator_name = ?", where.CreatorName)
	}
	if where.UpdaterID > 0 {
		db = db.Where("updater_id = ?", where.UpdaterID)
	}
	if where.UpdaterName != "" {
		db = db.Where("updater_name = ?", where.UpdaterName)
	}
	if where.DeletedAtLt != nil && *where.DeletedAtLt != zeroTime {
		sec := (*where.DeletedAtLt).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("deleted_at < ?", start)
	}
	if where.DeletedAtLte != nil && *where.DeletedAtLte != zeroTime {
		sec := (*where.DeletedAtLte).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("deleted_at <= ?", start)
	}
	if where.DeletedAtGt != nil && *where.DeletedAtGt != zeroTime {
		sec := (*where.DeletedAtGt).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("deleted_at > ?", start)
	}
	if where.DeletedAtGte != nil && *where.DeletedAtGte != zeroTime {
		sec := (*where.DeletedAtGte).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("deleted_at >= ?", start)
	}
	if where.ExpiredAtGt != nil && *where.ExpiredAtGt != zeroTime {
		sec := (*where.ExpiredAtGt).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("expired_at > ?", start)
	}
	if where.ExpiredAtGte != nil && *where.ExpiredAtGte != zeroTime {
		sec := (*where.ExpiredAtGte).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("expired_at >= ?", start)
	}
	if where.ExpiredAtLt != nil && *where.ExpiredAtLt != zeroTime {
		sec := (*where.ExpiredAtLt).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("expired_at < ?", start)
	}
	if where.ExpiredAtLte != nil && *where.ExpiredAtLte != zeroTime {
		sec := (*where.ExpiredAtLte).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("expired_at <= ?", start)
	}
	if where.CompletedAtGt != nil && *where.CompletedAtGt != zeroTime {
		sec := (*where.CompletedAtGt).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("completed_at > ?", start)
	}
	if where.CompletedAtGte != nil && *where.CompletedAtGte != zeroTime {
		sec := (*where.CompletedAtGte).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("completed_at >= ?", start)
	}
	if where.CompletedAtLt != nil && *where.CompletedAtLt != zeroTime {
		sec := (*where.CompletedAtLt).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("completed_at < ?", start)
	}
	if where.CompletedAtLte != nil && *where.CompletedAtLte != zeroTime {
		sec := (*where.CompletedAtLte).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("completed_at <= ?", start)
	}
	if where.TradeTimeGt != nil && *where.TradeTimeGt != zeroTime {
		sec := (*where.TradeTimeGt).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("trade_time > ?", start)
	}
	if where.TradeTimeGte != nil && *where.TradeTimeGte != zeroTime {
		sec := (*where.TradeTimeGte).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("trade_time >= ?", start)
	}
	if where.TradeTimeLt != nil && *where.TradeTimeLt != zeroTime {
		sec := (*where.TradeTimeLt).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("trade_time < ?", start)
	}
	if where.TradeTimeLte != nil && *where.TradeTimeLte != zeroTime {
		sec := (*where.TradeTimeLte).Unix()
		start := time.Unix(int64(sec), 0).UTC()
		db = db.Where("trade_time <= ?", start)
	}
	return db
}
