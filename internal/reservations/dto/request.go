package dto

type CreateReservationRequest struct {
	ZoneID       uint   `json:"zone_id" validate:"required"`
	LicensePlate string `json:"license_plate" validate:"required,max=15"`
	Status       string `json:"status,omitempty" validate:"omitempty,oneof=active completed cancelled"`
}
