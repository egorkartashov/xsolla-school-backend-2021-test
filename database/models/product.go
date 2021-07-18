package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID           uuid.UUID
	Sku          string `gorm:"uniqueIndex"`
	Name         string
	Type         string //TODO maybe not string
	PriceInCents int32
}
