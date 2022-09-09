package postgresql

import (
	"fmt"
	"time"

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
	sqlDb, _ := db.DB()
	sqlDb.SetConnMaxIdleTime(10 * time.Minute)
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(2 * time.Hour)

	return db
}
