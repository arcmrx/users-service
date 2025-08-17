package database

import (
	"log"

	"github.com/arcmrx/users-service/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=arcuser password=secret dbname=users port=5432 sslmode=disable"
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
		return nil, err
	}

	if err := DB.AutoMigrate(&user.User{}); err != nil {
		log.Fatalf("Could not migrate database: %v", err)
		return nil, err
	}

	return DB, nil
}
