package models

import (
	"time"

	"gorm.io/gorm"
)

type MediaRequest struct {
	gorm.Model
	Name                  string
	MediaType             string
	ReleaseDate           *time.Time
	RemoteIndexIdentifier int64
	RemoteIndexID         uint
	RemoteIndex           RemoteIndex
}
