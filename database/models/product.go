package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id           uuid.UUID
	Sku          string
	Name         string
	Type         string //TODO maybe not string
	PriceInCents int32
}
