package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/rizface/golang-api-template/app/controller/usercontroller"
	"github.com/rizface/golang-api-template/app/repository/profilerepository"
	"github.com/rizface/golang-api-template/app/repository/userrepository"
	"github.com/rizface/golang-api-template/app/service/userservice"
	"gorm.io/gorm"
)

func SetupController(router *chi.Mux, postgres *gorm.DB) {
	userrepository := userrepository.New()
	profilerepository := profilerepository.New()
	userservice := userservice.New(userrepository, profilerepository, postgres)
	userController := usercontroller.New(userservice)
	usercontroller.Setup(router, userController)
}

func CreateHttpServer(router http.Handler) *http.Server {
	var appPort string

	if len(appPort) == 0 {
		appPort = "9000"
	} else {
		appPort = os.Getenv("APP_PORT")
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%v", appPort),
		Handler: router,
	}

	return httpServer
}
