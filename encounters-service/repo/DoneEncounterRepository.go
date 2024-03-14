package repo

import (
	"encounters/model"

	"gorm.io/gorm"
)

type DoneEncounterRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *DoneEncounterRepository) FindByUserId(userId string) ([]model.DoneEncounter, error) {
	var doneEncounters []model.DoneEncounter
	dbResult := repo.DatabaseConnection.Find(&doneEncounters, "userId = ?", userId)
	if dbResult != nil {
		return doneEncounters, dbResult.Error
	}
	return doneEncounters, nil
}

func (repo *DoneEncounterRepository) Create(doneEncounter *model.DoneEncounter) error {
	dbResult := repo.DatabaseConnection.Create(doneEncounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *DoneEncounterRepository) Delete(userId string, encounterId string) error {
	var deletedDoneEncounter model.DoneEncounter
	dbResult := repo.DatabaseConnection.Where("userId = ?", userId).Where("encounterId = ?", encounterId).Delete(&deletedDoneEncounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affeccted: ", dbResult.RowsAffected)
	return nil
}
