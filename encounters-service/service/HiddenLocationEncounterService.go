package service

import (
	"encounters/model"
	"encounters/repo"
	"fmt"
)

type HiddenLocationEncounterService struct {
	EncounterRepo *repo.HiddenLocationEncounterRepository
}

func (service *HiddenLocationEncounterService) FindEncounter(id string) (*model.HiddenLocationEncounter, error) {
	encounter, err := service.EncounterRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}

	return &encounter, nil
}

func (service *HiddenLocationEncounterService) Create(encounter *model.HiddenLocationEncounter) error {
	err := service.EncounterRepo.CreateEncounter(encounter)
	if err != nil {
		return err
	}
	return nil
}
