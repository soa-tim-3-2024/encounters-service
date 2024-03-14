package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"time"
)

type DoneEncounterStatus int

const (
	InProgress DoneEncounterStatus = iota
	Completed
)

type DoneEncounter struct {
	EncounterId         uuid.UUID           `json:"encounterId"`
	UserId              int                 `json:"userId"`
	DoneEncounterStatus DoneEncounterStatus `json:"doneEncounterStatus"`
	CompletionTime      time.Time           `json:"completionTime"`
}

func (doneEncounter *DoneEncounter) BeforeCreate(scope *gorm.DB) error {
	return nil
}

func (doneEncounter *DoneEncounter) Complete() error {
	if doneEncounter.DoneEncounterStatus == DoneEncounterStatus(Completed) {
		return errors.New("counter already completed")
	}
	doneEncounter.CompletionTime = time.Now()
	doneEncounter.DoneEncounterStatus = DoneEncounterStatus(Completed)
	return nil
}

func (a StringArray) ValueDoneEncounter() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *StringArray) ScanDoneEncounter(value interface{}) error {
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
