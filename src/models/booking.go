package models

import (
	"time"

	"github.com/google/uuid"
)

const CreatedAt = "CreatedAt"
const ModifiedAt = "ModifiedAt"
const DeletedAt = "DeletedAt"

type BookingTransportStatus struct {
	ID    string `gorm:"unique" json:"id,omitempty"`
	Title string `gorm:"not null" json:"title,omitempty"`
}

type BookingPaymentStatus struct {
	ID    string `gorm:"unique" json:"id,omitempty"`
	Title string `gorm:"not null" json:"title,omitempty"`
}

type Location struct {
	ID              uint          `gorm:"primary_key" json:"id,omitempty"`
	Name            string        `gorm:"not null" json:"name,omitempty"`
	Country         string        `gorm:"not null" json:"country,omitempty"`
	Code            string        `gorm:"not null" json:"code,omitempty"`
	IsTransitPoint  bool          `gorm:"not null" json:"isTransitPoint"`
	TransitMethodId uint          `gorm:"null;" json:"transitMethodId"`
	Parent          uint          `gorm:"not null" json:"parent,omitempty"`
	Order           uint          `gorm:"not null" json:"order,omitempty"`
	LocationGroupID uint          `gorm:"not null" json:"locationGroupID,omitempty"`
	LocationGroup   LocationGroup `gorm:"foreignKey:LocationGroupID;reference:ID"`
}

type LocationGroup struct {
	ID   uint   `gorm:"primary_key" json:"id,omitempty"`
	Name string `gorm:"not null" json:"name,omitempty"`
}

type TransportationType struct {
	ID   string `gorm:"unique" json:"id,omitempty"`
	Name string `gorm:"not null" json:"name,omitempty"`
}

type Booking struct {
	ID                     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	RoundTrip              string    `gorm:"null;default:no" json:"roundTrip,omitempty"`
	From                   int       `gorm:"null" json:"from,omitempty"`
	To                     int       `gorm:"null" json:"to,omitempty"`
	DateFrom               time.Time `gorm:"null" json:"dateFrom,omitempty"`
	DateTo                 time.Time `gorm:"null" json:"dateTo,omitempty"`
	BookingPaymentStatus   int       `gorm:"null;default:0" json:"bookingPaymentStatus,omitempty"`
	BookingTransportStatus int       `gorm:"null;default:0" json:"bookingTransportStatus,omitempty"`
	CreatedAt              time.Time `gorm:"null" json:"-"`
	UpdatedAt              time.Time `gorm:"null" json:"-"`
	DeletedAt              time.Time `gorm:"null" json:"-"`
}
