package repo

import (
	"encounters/model"

	"gorm.io/gorm"
)

type KeyPointEncounterRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *KeyPointEncounterRepository) FindById(id string) (model.KeyPointEncounter, error) {
	encounter := model.KeyPointEncounter{}
	dbResult := repo.DatabaseConnection.Preload("Encounter").First(&encounter, "encounter_id = ?", id)
	if dbResult != nil {
		return encounter, dbResult.Error
	}

	return encounter, nil
}

func (repo *KeyPointEncounterRepository) CreateEncounter(encounter *model.KeyPointEncounter) error {
	dbResult := repo.DatabaseConnection.Create(encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *KeyPointEncounterRepository) Save(encounter *model.KeyPointEncounter) error {
	dbResult := repo.DatabaseConnection.Save(encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
