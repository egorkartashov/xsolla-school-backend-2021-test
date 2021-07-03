package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductDto struct {
	gorm.Model
	Id           uuid.UUID
	Sku          string
	Name         string
	Type         string //TODO maybe not string
	PriceInCents int32
}
