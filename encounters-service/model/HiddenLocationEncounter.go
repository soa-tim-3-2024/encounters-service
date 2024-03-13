package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type HiddenLocationEncounter struct {
	EncounterID      uuid.UUID `gorm:"primaryKey"`
	Encounter        Encounter `json:",omitempty"`
	PictureLongitude float64   `json:"pictureLongitude"`
	PictureLatitude  float64   `json:"pictureLatitude"`
}

func (a StringArray) ValueHiddenLocation() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *StringArray) ScanHiddenLocation(value interface{}) error {
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
