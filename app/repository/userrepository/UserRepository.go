package userrepository

import (
	"github.com/google/uuid"
	"github.com/rizface/golang-api-template/app/entity/requestentity"
	"github.com/rizface/golang-api-template/app/entity/responseentity"
	"github.com/rizface/golang-api-template/database/postgresql"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindOneByUsername(username string, db *gorm.DB) (*responseentity.User, error)
	Create(user *requestentity.User, db *gorm.DB) (*responseentity.User, error)
}

type UserRepository struct{}

func New() UserRepositoryInterface {
	return &UserRepository{}
}

func (userrepository *UserRepository) FindOneByUsername(username string, db *gorm.DB) (*responseentity.User, error) {
	user := &postgresql.User{}
	result := db.Joins("Profile").Where(&postgresql.User{Username: username}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user.ToDomain(user.Password), nil
}

func (userrepository *UserRepository) Create(user *requestentity.User, db *gorm.DB) (*responseentity.User, error) {
	payload := &postgresql.User{
		ID:       uuid.NewString(),
		Username: user.Username,
		Password: user.Password,
	}
	err := db.Create(payload).Error
	if err != nil {
		return nil, err
	}
	return payload.ToDomain(nil), nil
}
