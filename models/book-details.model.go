package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Isbn           int    `gorm:"unique;not null" json:"isbn"`
	Title          string `gorm:"not null" json:"title"`
	Publisher      string `gorm:"not null" json:"publisher"`
	Published_year int    `gorm:"not null" json:"published_year"`
	Synopsis       string `gorm:"nullable" json:"synopsis"`
}
