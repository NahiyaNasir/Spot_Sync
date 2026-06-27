package user

import (
	"errors"

	"gorm.io/gorm"
)
 var ErrorAlreadyExists= errors.New("user already exists")
type Repository interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
	GetAllUsers() ([]User, error)
	Update(user *User) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

 func(r *repository) CreateUser(user *User) error{
result :=r.db.Create(&user) // pass pointer of data to Create
  if result.Error != nil {
	if errors.Is(result.Error,gorm.ErrDuplicatedKey){
		return  ErrorAlreadyExists
	}
	 return result.Error
  }
  return nil
 }
func(r *repository) GetUserByEmail(email string)(*User,error){
 var user User
 	result := r.db.Where(&User{Email: email}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
 }
 func (r *repository) GetAllUsers() ([]User, error) {
	var users []User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *repository) Update(u *User) error {
	return r.db.Save(u).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&User{}, id).Error
}
