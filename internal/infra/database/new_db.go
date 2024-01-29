package database

import (
	campaing "github.com/henrique998/email-N/internal/domain/campaign"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	dsn := "host=localhost user=henrique password=123456 dbname=postgres port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect to database")
	}

	db.AutoMigrate(&campaing.Campaing{}, &campaing.Contact{})

	return db
}
