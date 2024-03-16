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
	dbResult := repo.DatabaseConnection.Where("user_id = ? and encounter_id = ?", userId, encounterId).Delete(&deletedDoneEncounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affeccted: ", dbResult.RowsAffected)
	return nil
}

func (repo *DoneEncounterRepository) Save(doneEncounter model.DoneEncounter) error {
	dbResult := repo.DatabaseConnection.Where("user_id = ? and encounter_id = ?", doneEncounter.UserId, doneEncounter.EncounterId).Save(&doneEncounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affeccted: ", dbResult.RowsAffected)
	return nil
}

func (repo *DoneEncounterRepository) Find(userId string, encounterId string) (model.DoneEncounter, error) {
	var doneEncounter model.DoneEncounter
	dbResult := repo.DatabaseConnection.Where("user_id = ? and encounter_id = ?", userId, encounterId).First(&doneEncounter)
	if dbResult.Error != nil {
		return model.DoneEncounter{}, dbResult.Error
	}
	return doneEncounter, nil
}

func (repo *DoneEncounterRepository) GetActiveByUserId(userId string) (*[]model.DoneEncounter, error) {
	var doneEncounters []model.DoneEncounter
	dbResult := repo.DatabaseConnection.Find(&doneEncounters, "user_id = ? and done_encounter_status = 0", userId)
	return &doneEncounters, dbResult.Error
}

func (repo *DoneEncounterRepository) GetCompletedByUserId(userId string) (*[]model.DoneEncounter, error) {
	var doneEncounters []model.DoneEncounter
	dbResult := repo.DatabaseConnection.Find(&doneEncounters, "user_id = ? and done_encounter_status = 1", userId)
	return &doneEncounters, dbResult.Error
}

func (repo *DoneEncounterRepository) IsCompleted(userId string, encounterId string) bool {
	_, err := repo.Find(userId, encounterId)
	return err != nil
}
