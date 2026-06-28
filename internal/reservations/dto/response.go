package dto


type ZoneResponse struct {
    ID   uint   `json:"id"`
    Name string `json:"name"`
    Type string `json:"type"`
}
type MyReservationResponse struct {
	ID           uint   `json:"id"`
	UserID       uint   `json:"user_id"`
	ZoneID       uint   `json:"zone_id"`
	LicensePlate string `json:"license_plate"`
	Status       string `json:"status"`
	Zone           ZoneResponse `json:"zone"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}


type CreateReservationResponse struct {
	ID           uint   `json:"id"`
	UserID       uint   `json:"user_id"`
	ZoneID       uint   `json:"zone_id"`
	LicensePlate string `json:"license_plate"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
