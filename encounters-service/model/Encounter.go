package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// enum u go-u
type EncounterStatus int

type EncounterType int

const (
	Active EncounterStatus = iota
	Draft
	Archived
)

const (
	Social EncounterType = iota
	Hidden
	Misc
	KeyPoint
)

type StringArray []string

type Encounter struct {
	ID          uuid.UUID       `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Picture     string          `json:"picture"`
	Longitude   float64         `json:"longitude"`
	Latitude    float64         `json:"latitude"`
	XpReward    int             `json:"xpReward"`
	Status      EncounterStatus `json:"status"`
	Type        EncounterType   `json:"type"`
}

func (encounter *Encounter) BeforeCreate(scope *gorm.DB) error {
	encounter.ID = uuid.New()
	return nil
}

func (a StringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = nil
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, a)
}
