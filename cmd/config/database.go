package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	// dsn := "host=localhost user=postgres password= dbname=go_grpc port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	dsn := "postgres://postgres:@localhost:5432/go_grpc?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database" + err.Error())
	}

	return db
}
