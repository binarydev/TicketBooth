package datastore

import (
	"gorm.io/gorm"
)

type RemoteIndex struct {
	gorm.Model
	Name   string
	URL    string `gorm:"column:url"`
	APIKey string `gorm:"column:api_key"`
}
