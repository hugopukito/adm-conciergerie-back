package entity

import (
	"time"

	"github.com/google/uuid"
)

type PropertyType string
type PropertyCondition string

const (
	House     PropertyType      = "house"
	Apartment PropertyType      = "apartment"
	Renovate  PropertyCondition = "renovate"
	Decorate  PropertyCondition = "decorate"
	Ready     PropertyCondition = "ready"
)

type Form struct {
	ID                uuid.UUID         `json:"ID,omitempty"`
	Date              time.Time         `json:"date,omitempty"`
	PropertyType      PropertyType      `json:"propertyType,omitempty"`
	Surface           int               `json:"surface,omitempty"`
	PropertyCondition PropertyCondition `json:"propertyCondition,omitempty"`
	Mail              string            `json:"mail,omitempty"`
	Phone             string            `json:"phone,omitempty"`
	Notes             string            `json:"notes,omitempty"`
}
