package postgresql

import (
	"time"

	"github.com/rizface/golang-api-template/app/entity/responseentity"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	Username  string `gorm:"unique;index"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Profile   Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (user *User) ToDomain(password interface{}) *responseentity.User {
	profile := user.Profile.ToDomain()

	return &responseentity.User{
		Id:        user.ID,
		Username:  user.Username,
		Password:  password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Profile:   *profile,
	}
}
