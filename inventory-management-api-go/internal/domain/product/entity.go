package product

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:150;not null"`
	Description string    `gorm:"size:255"`
	SKU         string    `gorm:"size:80;not null;uniqueIndex"`
	Price       float64   `gorm:"type:numeric(12,2);not null;check:price >= 0"`
	Stock       int64     `gorm:"not null;default:0;check:stock >= 0"`
	CategoryID  uint      `gorm:"not null;index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
