package service

import (
	"encounters/model"
	"encounters/repo"
	"fmt"
)

type SocialEncounterService struct {
	EncounterRepo *repo.SocialEncounterRepository
}

func (service *SocialEncounterService) FindEncounter(id string) (*model.SocialEncounter, error) {
	encounter, err := service.EncounterRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}

	return &encounter, nil
}

func (service *SocialEncounterService) Create(encounter *model.SocialEncounter) error {
	err := service.EncounterRepo.CreateEncounter(encounter)
	if err != nil {
		return err
	}
	return nil
}
