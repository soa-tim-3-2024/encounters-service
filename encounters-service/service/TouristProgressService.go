package service

import (
	"encounters/model"
	"encounters/repo"
	"fmt"
)

type TouristProgressService struct {
	TPRepo *repo.TouristProgressRepository
}

func (service *TouristProgressService) FindTouristProgress(id string) (*model.TouristProgress, error) {
	touristProgress, err := service.TPRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &touristProgress, nil
}

func (service *TouristProgressService) Create(touristProgress *model.TouristProgress) error {
	err := service.TPRepo.CreateTouristProgress(touristProgress)
	if err != nil {
		return err
	}
	return nil
}

func (service *TouristProgressService) Save(progress *model.TouristProgress) error {
	return service.TPRepo.Save(progress)
}

func (service *TouristProgressService) GiveXp(userId string, xp int) error {
	progress, err := service.TPRepo.FindById(userId)
	if err != nil {
		return err
	}
	progress.Xp += xp
	return service.TPRepo.Save(&progress)
}
