package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"math"

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

func (encounter *Encounter) CanActivate(userLongitute float64, userLatitude float64) (bool, error) {
	if encounter.Status != EncounterStatus(Active) {
		return false, errors.New("encounter is not active")
	}
	if userLongitute < -180 || userLongitute > 180 {
		return false, errors.New("invalid longitude")
	}
	if userLatitude < -90 || userLatitude > 90 {
		return false, errors.New("invalid latitude")
	}
	earthRadius := 6371000
	encounterRadius := 50 //meters
	latitude1 := encounter.Latitude * math.Pi / 180
	longitude1 := encounter.Longitude * math.Pi / 180
	latitude2 := userLatitude * math.Pi / 180
	longitude2 := userLongitute * math.Pi / 180

	latitudeDistance := latitude2 - latitude1
	longitudeDistance := longitude2 - longitude1

	a := math.Sin(latitudeDistance/2)*math.Sin(latitudeDistance/2) +
		math.Cos(latitude1)*math.Cos(latitude2)*
			math.Sin(longitudeDistance/2)*math.Sin(longitudeDistance/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := float64(earthRadius) * c

	return (distance < float64(encounterRadius)), nil
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
