package main

import (
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
