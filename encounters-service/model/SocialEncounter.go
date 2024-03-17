package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type SocialEncounter struct {
	EncounterID  int       `gorm:"primaryKey"`
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

func (encounter *SocialEncounter) BeforeCreate(scope *gorm.DB) error {
	//encounter.EncounterID = uuid.New()
	return nil
}
