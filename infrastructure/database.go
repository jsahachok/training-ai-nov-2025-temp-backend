package infrastructure

import (
	"log"
	infrastructureModels "workshop4-backend/infrastructure/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() *Database {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&infrastructureModels.UserModel{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database connected successfully")

	return &Database{
		DB: db,
	}
}

func (d *Database) GetDB() *gorm.DB {
	return d.DB
}
