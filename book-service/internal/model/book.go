package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title           string    `gorm:"not null"`
	Author          string
	ISBN            string `gorm:"uniqueIndex;not null"`
	Description     string
	TotalCopies     int32 `gorm:"not null"`
	AvailableCopies int32 `gorm:"not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}
