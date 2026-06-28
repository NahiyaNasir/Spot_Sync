package parkingzone

import (
	"time"

	"spot_sync/internal/auth"
	"spot_sync/internal/parkingzone/dto"

	"gorm.io/gorm"
)

type service struct {
	repo Repository
	jwtService auth.JWTService
	 db   *gorm.DB
}

func NewService(repo Repository,jwtService auth.JWTService, db *gorm.DB) *service {
	return &service{repo: repo, jwtService: jwtService, db: db}
}

func (s *service) CreateParkingZone(req *dto.CreateParkingZoneRequest) (*dto.ParkingZoneResponse, error) {
	zone := &ParkingZone{
		Name:          req.Name,
		Type:          req.Type,
		TotalCapacity: req.TotalCapacity,
		PricePerHour:  req.PricePerHour,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := s.repo.CreateParkingZone(zone); err != nil {
		return nil, err
	}

	return &dto.ParkingZoneResponse{
		ID:            zone.ID,
		Name:          zone.Name,
		Type:          zone.Type,
		TotalCapacity: zone.TotalCapacity,
		PricePerHour:  zone.PricePerHour,
		CreatedAt:     zone.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     zone.UpdatedAt.Format(time.RFC3339),
	}, nil
}
  func (s *service) GetParkingZoneByID(id uint) (*dto.ParkingZoneResponse, error) {
	zone, err := s.repo.GetParkingZoneByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.ParkingZoneResponse{
		ID:            zone.ID,
		Name:          zone.Name,
		Type:          zone.Type,
		TotalCapacity: zone.TotalCapacity,
		PricePerHour:  zone.PricePerHour,
		CreatedAt:     zone.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     zone.UpdatedAt.Format(time.RFC3339),
	}, nil
  }
func (s *service) GetAllParkingZones() ([]*dto.ParkingZoneResponse, error) {
	zones, err := s.repo.GetAllParkingZones()
	if err != nil {
		return nil, err
	}
	var responses []*dto.ParkingZoneResponse
	for _, zone := range zones {
		var activeReservations int64

		err := s.db.Table("reservations").
			Where("zone_id = ? AND status = ?", zone.ID, "active").
			Count(&activeReservations).Error

		if err != nil {
			return nil, err
		}
		availableSpots := zone.TotalCapacity - int(activeReservations)
		responses = append(responses, &dto.ParkingZoneResponse{
			ID:            zone.ID,
			Name:          zone.Name,
			Type:          zone.Type,
			TotalCapacity: zone.TotalCapacity,
			AvailableSpots:  availableSpots,
			PricePerHour:  zone.PricePerHour,
			CreatedAt:     zone.CreatedAt.Format(time.RFC3339),
			UpdatedAt:     zone.UpdatedAt.Format(time.RFC3339),
		})
	}
	return responses, nil
	
}