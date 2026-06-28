package reservations

import (
	"time"

	"spot_sync/internal/parkingzone"
	"spot_sync/internal/reservations/dto"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	UserID       uint   `gorm:"not null" json:"user_id"`
	ZoneID       uint   `gorm:"not null" json:"zone_id"`
	LicensePlate string `gorm:"type:varchar(15);not null" json:"license_plate"`
	Status       string `gorm:"type:varchar(20);not null;default:'active'" json:"status"`
    Zone parkingzone.ParkingZone `gorm:"foreignKey:ZoneID;references:ID" json:"zone"`
}
func (r *Reservation) ToResponse() *dto.MyReservationResponse {
    
    return &dto.MyReservationResponse{
        ID:           r.ID,
         UserID:       r.UserID,  
        ZoneID:       r.ZoneID,
        LicensePlate: r.LicensePlate,
        Status:       r.Status,
        Zone: dto.ZoneResponse{
            ID:   r.Zone.ID,
            Name: r.Zone.Name,
            Type: r.Zone.Type,
        },
        CreatedAt: r.CreatedAt.Format(time.RFC3339),
    }
}

