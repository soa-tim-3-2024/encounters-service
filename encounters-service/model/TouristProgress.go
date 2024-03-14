package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TouristProgress struct {
	ID     uuid.UUID `json:"id"`
	UserId float64   `json:"userId"`
	Xp     int       `json:"xp"`
	Level  int       `json:"level"`
}

func (touristProgression *TouristProgress) BeforeCreate(scope *gorm.DB) error {
	touristProgression.ID = uuid.New()
	return nil
}
