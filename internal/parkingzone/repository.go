package parkingzone

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	CreateParkingZone(zone *ParkingZone) error
	GetParkingZoneByID(id uint) (*ParkingZone, error)
		GetAllParkingZones() ([]*ParkingZone, error)
}

type repository struct {
	db *gorm.DB
}
var ErrParkingZoneNotFound=errors.New("parking zone not found")

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateParkingZone(zone *ParkingZone) error {
	return r.db.Create(zone).Error
}
func (r *repository) GetParkingZoneByID(id uint) (*ParkingZone, error) {
	var zone ParkingZone	
	if err := r.db.First(&zone, id).Error; err != nil {
		return nil, err
	}
	return &zone, nil
}
func (r *repository) GetAllParkingZones() ([]*ParkingZone, error) {
	var zones []*ParkingZone	
	result := r.db.Find(&zones)
	if result.Error != nil {
		return nil, result.Error
	}
	return zones, nil}
