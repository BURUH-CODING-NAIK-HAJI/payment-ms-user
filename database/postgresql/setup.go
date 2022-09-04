package postgresql

import (
	"fmt"

	databaseconfig "github.com/rizface/golang-api-template/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(dbConfig *databaseconfig.Database) *gorm.DB {
	dsn := fmt.Sprintf(
		"user=%s password=%s port=%s host=%s sslmode=disable dbname=%s TimeZone=Asia/Jakarta",
		dbConfig.Username, dbConfig.Password, dbConfig.Port, dbConfig.Host, dbConfig.Name,
	)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
