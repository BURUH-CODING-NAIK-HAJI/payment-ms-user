package usercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rizface/golang-api-template/app/entity/requestentity"
	"github.com/rizface/golang-api-template/app/entity/responseentity"
	"github.com/rizface/golang-api-template/app/errorgroup"
	"github.com/rizface/golang-api-template/app/service/userservice"
)

type UserControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	FindOneByUsername(w http.ResponseWriter, r *http.Request)
}

type UserController struct {
	userservice userservice.UserServiceInterface
}

func New(userservice userservice.UserServiceInterface) UserControllerInterface {
	return &UserController{
		userservice: userservice,
	}
}

func Setup(router *chi.Mux, controller UserControllerInterface) {
	router.Post("/", controller.Create)
	router.Get("/{username}/resources", controller.FindOneByUsername)
}

func (controller *UserController) Create(w http.ResponseWriter, r *http.Request) {
	payload := new(requestentity.User)
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		panic(errorgroup.REGISTER_FAILED)
	}

	err = payload.Validate()
	if err != nil {
		panic(err)
	}

	result := controller.userservice.Create(payload)
	json.NewEncoder(w).Encode(responseentity.Response{
		Code:    http.StatusOK,
		Message: "SUCCESS",
		Data:    result,
	})
}

func (controller *UserController) FindOneByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	result := controller.userservice.FindOneByUsername(username)
	json.NewEncoder(w).Encode(responseentity.Response{
		Code:    http.StatusOK,
		Message: "SUCCESS",
		Data:    result,
	})
}
