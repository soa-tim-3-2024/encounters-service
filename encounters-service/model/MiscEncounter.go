package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type MiscEncounter struct {
	EncounterID   uuid.UUID `gorm:"primaryKey"`
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
