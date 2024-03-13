package service

import (
	"encounters/model"
	"encounters/repo"
	"fmt"
)

type KeyPointEncounterService struct {
	EncounterRepo *repo.KeyPointEncounterRepository
}

func (service *KeyPointEncounterService) FindEncounter(id string) (*model.KeyPointEncounter, error) {
	encounter, err := service.EncounterRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}

	return &encounter, nil
}

func (service *KeyPointEncounterService) Create(encounter *model.KeyPointEncounter) error {
	err := service.EncounterRepo.CreateEncounter(encounter)
	if err != nil {
		return err
	}
	return nil
}
