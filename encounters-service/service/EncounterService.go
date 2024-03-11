package service

import (
	"encounters/model"
	"encounters/repo"
)

type EncounterService struct {
	EncounterRepo *repo.EncounterRepository
}

func (service *EncounterService) Create(encounter *model.Encounter) error {
	err := service.EncounterRepo.CreateEncounter(encounter)
	if err != nil {
		return err
	}
	return nil
}

func (service *EncounterService) Update(encounter *model.Encounter) error {
	err := service.EncounterRepo.UpdateEncounter(encounter)
	if err != nil {
		return err
	}
	return nil
}
