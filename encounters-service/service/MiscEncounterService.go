package service

import (
	"encounters/model"
	"encounters/repo"
	"fmt"
)

type MiscEncounterService struct {
	EncounterRepo *repo.MiscEncounterRepository
}

func (service *MiscEncounterService) FindEncounter(id string) (*model.MiscEncounter, error) {
	encounter, err := service.EncounterRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}

	return &encounter, nil
}

func (service *MiscEncounterService) Create(encounter *model.MiscEncounter) error {
	err := service.EncounterRepo.CreateEncounter(encounter)
	if err != nil {
		return err
	}
	return nil
}
