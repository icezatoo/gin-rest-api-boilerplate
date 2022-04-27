package db

import (
	"github.com/icezatoo/gin-rest-api-boilerplate/config"
	"github.com/icezatoo/gin-rest-api-boilerplate/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection(config *config.Config) *gorm.DB {

	db, err := gorm.Open(postgres.Open(config.DbSource), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connection to database failed")
		logrus.Fatalf("Failed to connect to database: %v", err)
	}

	if config.Environment != "production" {
		logrus.Info("Connection to database is successful")
	}

	err = db.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
