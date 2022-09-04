package postgresql_test

import (
	"testing"

	"github.com/rizface/golang-api-template/database"
	"github.com/rizface/golang-api-template/database/postgresql"
	"github.com/stretchr/testify/assert"
)

func TestConnec(t *testing.T) {
	dbConfig := database.New()
	dbConfig.Setup()
	db := postgresql.NewConnection(dbConfig)
	assert.NotNil(t, db)
}
