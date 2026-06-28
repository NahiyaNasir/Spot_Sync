package parkingzone

import (
	"time"

	"spot_sync/internal/parkingzone/dto"
	"gorm.io/gorm"
)

type ParkingZone struct {
	gorm.Model
	Name          string    `gorm:"type:varchar(100);not null" json:"name"`
	Type          string    `gorm:"type:varchar(30);not null" json:"type"`
	TotalCapacity int       `gorm:"not null" json:"total_capacity"`
	PricePerHour  float64   `gorm:"type:decimal(10,2);not null" json:"price_per_hour"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (p *ParkingZone) ToResponse() *dto.ParkingZoneResponse {
	return &dto.ParkingZoneResponse{
		ID:            p.ID,
		Name:          p.Name,
		Type:          p.Type,
		TotalCapacity: p.TotalCapacity,
		PricePerHour:  p.PricePerHour,
		CreatedAt:     p.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     p.UpdatedAt.Format(time.RFC3339),
	}
}
