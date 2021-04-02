package model

// Card ...
type Card struct {
	ID   uint64 `gorm:"id"`
	Name string `gorm:"name"`
}

func (Card) TableName() string {
	return "cards"
}
