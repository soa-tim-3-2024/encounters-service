package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type HiddenLocationEncounter struct {
	EncounterID      int       `gorm:"primaryKey"`
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

func (encounter *HiddenLocationEncounter) BeforeCreate(scope *gorm.DB) error {
	//encounter.EncounterID = uuid.New()
	return nil
}
