package postgresql

import (
	"time"

	"github.com/rizface/golang-api-template/app/entity/responseentity"
)

type Profile struct {
	ID        string `gorm:"primaryKey"`
	UserId    string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (profile *Profile) ToDomain() *responseentity.Profile {
	return &responseentity.Profile{
		Id:        profile.ID,
		UserId:    profile.UserId,
		Name:      profile.Name,
		CreatedAt: profile.CreatedAt,
		UpdatedAt: profile.UpdatedAt,
	}
}
