package service

import (
	"encounters/model"
	"encounters/repo"
	"fmt"
)

type DoneEncounterService struct {
	DoneEncounterRepo *repo.DoneEncounterRepository
}

func (service *DoneEncounterService) FindByUserId(userId string) ([]model.DoneEncounter, error) {
	doneEncounters, err := service.DoneEncounterRepo.FindByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", userId))
	}

	return doneEncounters, nil
}

func (service *DoneEncounterService) Create(doneEncounter *model.DoneEncounter) error {
	err := service.DoneEncounterRepo.Create(doneEncounter)
	return err
}
