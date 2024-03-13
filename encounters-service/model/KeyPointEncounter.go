package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type KeyPointEncounter struct {
	EncounterID uuid.UUID `gorm:"primaryKey"`
	Encounter   Encounter `json:",omitempty"`
	KeyPointId  int       `json:"keyPointId"`
}

func (a StringArray) ValueKeyPoint() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *StringArray) ScanKeyPoint(value interface{}) error {
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
