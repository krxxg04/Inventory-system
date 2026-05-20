package category

import "time"

type Category struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:120;not null"`
	Description string    `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
