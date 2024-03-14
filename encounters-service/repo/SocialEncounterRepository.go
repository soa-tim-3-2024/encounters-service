package repo

import (
	"encounters/model"

	"gorm.io/gorm"
)

type SocialEncounterRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *SocialEncounterRepository) FindById(id string) (model.SocialEncounter, error) {
	encounter := model.SocialEncounter{}
	dbResult := repo.DatabaseConnection.Preload("Encounter").First(&encounter, "encounter_id = ?", id)
	if dbResult != nil {
		return encounter, dbResult.Error
	}

	return encounter, nil
}

func (repo *SocialEncounterRepository) CreateEncounter(encounter *model.SocialEncounter) error {
	dbResult := repo.DatabaseConnection.Create(encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *SocialEncounterRepository) Save(encounter *model.SocialEncounter) error {
	dbResult := repo.DatabaseConnection.Save(encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
