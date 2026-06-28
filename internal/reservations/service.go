package reservations

import (
	"fmt"
	"time"

	"spot_sync/internal/reservations/dto"
)

type service struct {
    repo Repository
}

func NewService(repo Repository) *service {
    return &service{repo: repo}
}


func (s *service) CreateReservation(userID uint, req *dto.CreateReservationRequest) (*dto.CreateReservationResponse, error) {
	reservation, err := s.repo.CreateWithZoneUpdate(userID, req.ZoneID, req.LicensePlate)
	if err != nil {
		return nil, err
	}
 fmt.Printf("DEBUG SERVICE userID: %d, ZoneID: %d\n", userID, req.ZoneID)
	return &dto.CreateReservationResponse{
		ID:           reservation.ID,
		UserID:       reservation.UserID,
		ZoneID:       reservation.ZoneID,
		LicensePlate: reservation.LicensePlate,
		Status:       reservation.Status,
		CreatedAt:    reservation.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    reservation.UpdatedAt.Format(time.RFC3339),
	}, nil


}


func (s *service) GetAllReservations() ([]*dto.CreateReservationResponse, error) {
    reservations, err := s.repo.GetAllReservations()
    if err != nil {
        return nil, err
    }

    var responses []*dto.CreateReservationResponse
    for _, reservation := range reservations {
        responses = append(responses, &dto.CreateReservationResponse{
            ID:           reservation.ID,
            UserID:       reservation.UserID,
            ZoneID:       reservation.ZoneID,
            LicensePlate: reservation.LicensePlate,
            Status:       reservation.Status,
            CreatedAt:    reservation.CreatedAt.Format(time.RFC3339),
            UpdatedAt:    reservation.UpdatedAt.Format(time.RFC3339),
        })
    }
    return responses, nil
	
}

func (s *service) GetMyReservations(userId uint) ([]*dto.MyReservationResponse, error) {
	reservations, err := s.repo.GetByUserID(userId)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.MyReservationResponse, len(reservations)) // Initialize the slice with the correct length

	for i, r := range reservations {
		responses[i] = r.ToResponse()
	}

	return responses, nil
}
func (s *service) CancelReservation(userID uint, reservationID uint) error {
    return s.repo.CancelReservation(userID, reservationID)
}