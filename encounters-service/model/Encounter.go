package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// enum u go-u
type EncounterStatus int

type EncounterType int

type StringArray []string

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

type Encounter struct {
	ID          int             `json:"id"`
	Title       string          `json:"name"`
	Description string          `json:"description"`
	Picture     string          `json:"picture"`
	Longitude   float64         `json:"longitude"`
	Latitude    float64         `json:"latitude"`
	XpReward    int             `json:"xpReward"`
	Status      EncounterStatus `json:"status"`
	Type        EncounterType   `json:"type"`
}

// konvertuje tip podatka iz go-a u tip podatka u bazi (jer gorm ne moze sam da rukuje sa nizom stringova kao atributom)
func (a StringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Konvertuje iz tipa podatka iz baze u tip podatka u go-u
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
