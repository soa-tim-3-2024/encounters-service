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
	return service.DoneEncounterRepo.Create(doneEncounter)
}

func (service *DoneEncounterService) Find(userId string, encounterId string) (model.DoneEncounter, error) {
	return service.DoneEncounterRepo.Find(userId, encounterId)
}

func (service *DoneEncounterService) Save(doneEncounter model.DoneEncounter) error {
	return service.DoneEncounterRepo.Save(doneEncounter)
}

func (service *DoneEncounterService) Delete(userId string, encounterId string) error {
	return service.DoneEncounterRepo.Delete(userId, encounterId)
}

func (service *DoneEncounterService) GetCompletedByUserId(userId string) (*[]model.DoneEncounter, error) {
	return service.DoneEncounterRepo.GetCompletedByUserId(userId)
}

func (service *DoneEncounterService) GetActiveByUserId(userId string) (*[]model.DoneEncounter, error) {
	return service.DoneEncounterRepo.GetActiveByUserId(userId)
}

func (service *DoneEncounterService) IsCompleted(userId string, encounterId string) bool {
	return service.DoneEncounterRepo.IsCompleted(userId, encounterId)
}
