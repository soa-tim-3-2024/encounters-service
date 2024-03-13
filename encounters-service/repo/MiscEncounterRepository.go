package repo

import (
	"encounters/model"

	"gorm.io/gorm"
)

type MiscEncounterRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *MiscEncounterRepository) FindById(id string) (model.MiscEncounter, error) {
	encounter := model.MiscEncounter{}
	dbResult := repo.DatabaseConnection.Preload("Encounter").First(&encounter, "encounter_id = ?", id)
	if dbResult != nil {
		return encounter, dbResult.Error
	}

	return encounter, nil
}

func (repo *MiscEncounterRepository) CreateEncounter(encounter *model.MiscEncounter) error {
	dbResult := repo.DatabaseConnection.Create(encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
