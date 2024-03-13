package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type SocialEncounter struct {
	EncounterID  uuid.UUID `gorm:"primaryKey"`
	Encounter    Encounter `json:",omitempty"`
	PeopleNumber int       `json:"peopleNumber"`
}

func (a StringArray) ValueSocial() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *StringArray) ScanSocial(value interface{}) error {
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
