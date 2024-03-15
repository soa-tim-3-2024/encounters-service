package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type MiscEncounter struct {
	EncounterID   int       `gorm:"primaryKey"`
	Encounter     Encounter `json:",omitempty"`
	ChallengeDone bool      `json:"challengeDone"`
}

func (a StringArray) ValueMisc() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *StringArray) ScanMisc(value interface{}) error {
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

func (encounter *MiscEncounter) BeforeCreate(scope *gorm.DB) error {
	//encounter.EncounterID = uuid.New()
	return nil
}
