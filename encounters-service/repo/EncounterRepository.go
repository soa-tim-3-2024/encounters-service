package repo

import (
	"encounters/model"

	"gorm.io/gorm"
)

type EncounterRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *EncounterRepository) GetEncounters() ([]model.Encounter, error) {
	encounters := []model.Encounter{}
	dbResult := repo.DatabaseConnection.Find(&encounters)

	if dbResult != nil {
		return encounters, dbResult.Error
	}
	return encounters, nil
}
