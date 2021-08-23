package dto

import (
	"github.com/google/uuid"
)

type ProductDto struct {
	Id           *uuid.UUID `json:"id,omitempty" validate:"omitempty,uuid"`
	Sku          string     `json:"sku" validate:"required"`
	Name         string     `json:"name" validate:"required"`
	Type         string     `json:"type" validate:"required"`
	PriceInCents int32      `json:"priceInCents" validate:"required,gte=0"`
	LandingUrl   string     `json:"landingUrl,omitempty" validate:"omitempty,url"`
}
