package reservations

import (
	"errors"
	"fmt"

	"spot_sync/internal/parkingzone"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)
var (
	ErrZoneNotFound         = errors.New("zone not found")
	ErrNotEnoughSpots        = errors.New("not enough spots available")
	ErrSpotsAlreadyCancelled = errors.New("spots already cancelled")
	ErrForbiddenParkingAccess  = errors.New("you do not own this booking")
)
type Repository interface {
    CreateReservation(reservation *Reservation) error
    GetAllReservations() ([]*Reservation, error)
	CreateWithZoneUpdate(userId uint, ZoneID uint, licensePlate string) (*Reservation, error)
	GetByUserID(userId uint) ([]*Reservation, error)
	GetByID(reservationId uint) (*Reservation, error)
	CancelReservation(userID uint, reservationID uint) error

}

type repository struct {
    db *gorm.DB
}

var (
    ErrReservationNotFound = errors.New("reservation not found")
    ErrUnauthorized        = errors.New("unauthorized")
    ErrAlreadyCancelled    = errors.New("reservation already cancelled")
)

func NewRepository(db *gorm.DB) Repository {
    return &repository{db: db}
}

func (r *repository) CreateReservation(reservation *Reservation) error {
    return r.db.Create(reservation).Error
}

func (r *repository) GetReservationByID(id uint) (*Reservation, error) {
    var reservation Reservation
    if err := r.db.First(&reservation, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrReservationNotFound
        }
        return nil, err
    }
    return &reservation, nil
}

func (r *repository) GetAllReservations() ([]*Reservation, error) {
    var reservations []*Reservation
    if err := r.db.Find(&reservations).Error; err != nil {
        return nil, err
    }
    return reservations, nil
}

func (r *repository) CreateWithZoneUpdate(userId uint, ZoneID uint, licensePlate string) (*Reservation, error) {
	var reservation Reservation
  fmt.Printf("DEBUG REPO userId: %d, ZoneID: %d\n", userId, ZoneID)
	// start transaction
	err := r.db.Transaction(func(tx *gorm.DB) error {

		var zoneData parkingzone.ParkingZone

		err := tx.Clauses(clause.Locking{Strength: "Update"}).First(&zoneData, ZoneID).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrZoneNotFound
			}
			return err
		}
		var activeCount int64
const StatusActive = "active"
		err = tx.Model(&Reservation{}).
			Where("zone_id = ? AND status = ?", ZoneID, StatusActive).
			Count(&activeCount).Error
            
reservation.Status = StatusActive
		if err != nil {
			return err
		}

		if int(activeCount) >= zoneData.TotalCapacity {
			return ErrNotEnoughSpots
		}

		reservation = Reservation{
			UserID:       userId,
			ZoneID:       ZoneID,
			LicensePlate: licensePlate,
			Status:       StatusActive,
		}

		if err := tx.Create(&reservation).Error; err != nil {
			return err
		}


		return nil

	})

	if err != nil {
		return nil, err
	}

	return &reservation, nil

}

func (r *repository) GetByUserID(userId uint) ([]*Reservation, error) {
	var reservations []*Reservation

	err := r.db.Preload("Zone").Where("user_id = ?", userId).Find(&reservations).Error
	if err != nil {
		return nil, err
	}

	return reservations, nil
}

func (r *repository) GetByID(reservationId uint) (*Reservation, error) {
	var reservation Reservation

	err := r.db.First(&reservation, reservationId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrReservationNotFound
		}

		return nil, err
	}

	return &reservation, nil
}
func (r *repository) CancelReservation(userID uint, reservationID uint) error {
    return r.db.Transaction(func(tx *gorm.DB) error {
        var reservation Reservation

        err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
            First(&reservation, reservationID).Error
        if err != nil {
            if errors.Is(err, gorm.ErrRecordNotFound) {
                return ErrReservationNotFound
            }
            return err
        }

        if reservation.UserID != userID {
            return ErrUnauthorized
        }

        if reservation.Status == "cancelled" {
            return ErrAlreadyCancelled
        }

        // update status to cancelled first
        if err := tx.Model(&reservation).Update("status", "cancelled").Error; err != nil {
            return err
        }

        // soft delete (sets deleted_at)
        if err := tx.Delete(&reservation).Error; err != nil {
            return err
        }

        return nil
    })
}