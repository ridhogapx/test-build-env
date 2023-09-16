package model

import "gorm.io/gorm"

type Book struct {
  gorm.Model
  Title string `gorm:"type:varchar(20)"`
  Author string `gorm:"type:varchar(20)"`
}

type SaveBookRequest struct {
  Title string `json:"title"`
  Author string `json:"author"`
}
