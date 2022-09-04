package profilerepository

import (
	"github.com/google/uuid"
	"github.com/rizface/golang-api-template/app/entity/requestentity"
	"github.com/rizface/golang-api-template/app/entity/responseentity"
	"github.com/rizface/golang-api-template/database/postgresql"
	"gorm.io/gorm"
)

type ProfileRepositoryInterface interface {
	Create(userId string, db *gorm.DB, profile *requestentity.User) (*responseentity.Profile, error)
}

type ProfileRepository struct {
}

func New() ProfileRepositoryInterface {
	return &ProfileRepository{}
}

func (profilerepository *ProfileRepository) Create(userId string, db *gorm.DB, user *requestentity.User) (*responseentity.Profile, error) {
	payload := &postgresql.Profile{
		ID:     uuid.New().String(),
		UserId: userId,
		Name:   user.Name,
	}
	err := db.Create(payload).Error
	if err != nil {
		return nil, err
	}
	return payload.ToDomain(), nil
}
