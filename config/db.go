package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDBConnection(source string) *gorm.DB {
    db, err := gorm.Open(postgres.Open(source), &gorm.Config{})

}
