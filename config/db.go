package config

import (
	"test-build-env/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDBConnection(source string)  {
  var err error  
  DB, err = gorm.Open(postgres.Open(source), &gorm.Config{})
   
  if err != nil {
      panic(err)
  }
  
  DB.AutoMigrate(&model.Book{})

}
