package repo

import (
	"encounters/model"

	"gorm.io/gorm"
)

type HiddenLocationEncounterRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *HiddenLocationEncounterRepository) FindById(id string) (model.HiddenLocationEncounter, error) {
	encounter := model.HiddenLocationEncounter{}
	dbResult := repo.DatabaseConnection.Preload("Encounter").First(&encounter, "encounter_id = ?", id)
	if dbResult != nil {
		return encounter, dbResult.Error
	}

	return encounter, nil
}

func (repo *HiddenLocationEncounterRepository) CreateEncounter(encounter *model.HiddenLocationEncounter) error {
	dbResult := repo.DatabaseConnection.Create(encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
