package main

import (
	"fmt"
	"os/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestDbConnection(t *testing.T) {
	dsn := "host=localhost user=rizal password=3748 dbname=db_startup_bwa port=5432 sslmode=disable TimeZone=UTC"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	assert.Nil(t, err)
}

func TestRetrieveUserTableFromDB(t *testing.T) {

	// get database connection
	dsn := "host=localhost user=rizal password=3748 dbname=db_startup_bwa port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// pastikan error == nil
	assert.Nil(t, err)

	// buat object array dari entity struct
	var users []user.User

	// pastikan object kosong sebelum query ke db
	assert.Equal(t, 0, len(users))

	// query ke db
	if assert.NotNil(t, db) {
		db.Find(&users)
	}

	// pastikan object user tidak nil lagi
	assert.NotEqual(t, 0, &users)

	for _, user := range users {
		fmt.Println(user.Name)
	}

}
