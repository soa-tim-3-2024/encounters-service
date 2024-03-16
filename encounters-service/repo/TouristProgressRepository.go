package repo

import (
	"encounters/model"

	"gorm.io/gorm"
)

type TouristProgressRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TouristProgressRepository) FindById(id string) (model.TouristProgress, error) {
	progress := model.TouristProgress{}
	dbResult := repo.DatabaseConnection.First(&progress, "user_id = ?", id)
	if dbResult != nil {
		return progress, dbResult.Error
	}
	return progress, nil
}

func (repo *TouristProgressRepository) CreateTouristProgress(progress *model.TouristProgress) error {
	dbResult := repo.DatabaseConnection.Create(progress)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TouristProgressRepository) Save(progress *model.TouristProgress) error {
	dbResult := repo.DatabaseConnection.Where("id = ?", progress.ID).Save(&progress)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affeccted: ", dbResult.RowsAffected)
	return nil
}
