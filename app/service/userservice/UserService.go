package userservice

import (
	"errors"

	"github.com/rizface/golang-api-template/app/entity/requestentity"
	"github.com/rizface/golang-api-template/app/entity/responseentity"
	"github.com/rizface/golang-api-template/app/errorgroup"
	"github.com/rizface/golang-api-template/app/repository/profilerepository"
	"github.com/rizface/golang-api-template/app/repository/userrepository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	Create(user *requestentity.User) map[string]interface{}
	FindOneByUsername(username string) *responseentity.User
}

type UserService struct {
	userrepository    userrepository.UserRepositoryInterface
	profilerepository profilerepository.ProfileRepositoryInterface
	db                *gorm.DB
}

func New(
	userrepository userrepository.UserRepositoryInterface,
	profilerepository profilerepository.ProfileRepositoryInterface,
	db *gorm.DB,
) UserServiceInterface {
	return &UserService{
		userrepository:    userrepository,
		profilerepository: profilerepository,
		db:                db,
	}
}

func (userservice *UserService) Create(user *requestentity.User) map[string]interface{} {
	var result map[string]interface{}
	existingUser, _ := userservice.userrepository.FindOneByUsername(user.Username, userservice.db)
	if existingUser != nil {
		panic(errorgroup.USERNAME_ALREADY_TAKEN)
	}

	err := userservice.db.Transaction(func(tx *gorm.DB) error {
		byteEncryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user.Password = string(byteEncryptedPassword)
		createdUser, err := userservice.userrepository.Create(user, tx)
		if err != nil {
			return err
		}

		createdProfile, err := userservice.profilerepository.Create(createdUser.Id, tx, user)
		if err != nil {
			return err
		}

		result = map[string]interface{}{
			"id":        createdUser.Id,
			"username":  createdUser.Username,
			"createdAt": createdUser.CreatedAt,
			"updatedAt": createdUser.UpdatedAt,
			"profile":   createdProfile,
		}
		return nil
	})

	if err != nil {
		// TODO LOG ORIGINAL ERROR
		panic(errorgroup.REGISTER_FAILED)
	}

	return result
}

func (userservice *UserService) FindOneByUsername(username string) *responseentity.User {
	result, err := userservice.userrepository.FindOneByUsername(username, userservice.db)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		panic(errorgroup.USER_NOT_FOUND)
	}

	return result
}
