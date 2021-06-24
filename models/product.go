package models

import (
	"github.com/google/uuid"
)

type Product struct {
	Id           uuid.UUID
	Sku          string
	Name         string
	Type         string //TODO maybe not string
	PriceInCents int32
}
