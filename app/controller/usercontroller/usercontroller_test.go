package usercontroller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dchest/uniuri"
	"github.com/go-chi/chi/v5"
	"github.com/rizface/golang-api-template/app/controller/usercontroller"
	"github.com/rizface/golang-api-template/app/entity/requestentity"
	"github.com/rizface/golang-api-template/app/repository/profilerepository"
	"github.com/rizface/golang-api-template/app/repository/userrepository"
	"github.com/rizface/golang-api-template/app/service/userservice"
	"github.com/rizface/golang-api-template/database"
	"github.com/rizface/golang-api-template/database/postgresql"
	"github.com/rizface/golang-api-template/system/router"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var username string
var userRepository userrepository.UserRepositoryInterface
var profileRepository profilerepository.ProfileRepositoryInterface
var userService userservice.UserServiceInterface
var userController usercontroller.UserControllerInterface
var Router *chi.Mux
var db *gorm.DB

func TestMain(m *testing.M) {
	tables := []interface{}{
		&postgresql.User{},
		&postgresql.Profile{},
	}
	username = uniuri.New()

	dbConfig := database.New()
	dbConfig.Setup()
	db = postgresql.NewConnection(dbConfig)
	db.AutoMigrate(tables...)
	db.Logger = db.Logger.LogMode(0)

	userRepository = userrepository.New()
	profileRepository = profilerepository.New()
	userService = userservice.New(
		userRepository,
		profileRepository,
		db,
	)
	userController = usercontroller.New(userService)
	Router = router.CreateRouter()
	usercontroller.Setup(Router, userController)
	m.Run()
	db.Migrator().DropTable(tables...)
}

func TestRegisterSuccess(t *testing.T) {
	payload := requestentity.User{
		Name:     uniuri.New(),
		Username: username,
		Password: uniuri.New(),
	}
	bytesPayload, _ := json.Marshal(payload)

	request := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(bytesPayload))
	recorder := httptest.NewRecorder()
	Router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestRegisterFailedBecauseDuplicateUsername(t *testing.T) {
	payload := requestentity.User{
		Name:     uniuri.New(),
		Username: username,
		Password: uniuri.New(),
	}
	bytesPayload, _ := json.Marshal(payload)

	request := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(bytesPayload))
	recorder := httptest.NewRecorder()
	Router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestRegisterFailedBadRequest(t *testing.T) {
	payload := requestentity.User{}
	bytesPayload, _ := json.Marshal(payload)

	request := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(bytesPayload))
	recorder := httptest.NewRecorder()
	Router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestFindUserByUsernameSuccess(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s/resources", username), nil)
	recorder := httptest.NewRecorder()
	Router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestFindUserByUsernameNotFound(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s/resources", uniuri.New()), nil)
	recorder := httptest.NewRecorder()
	Router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
}
