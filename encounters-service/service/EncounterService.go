package service

import (
	"encounters/model"
	"encounters/repo"
	"errors"
	"fmt"
	"strconv"
)

type EncounterService struct {
	HiddenLocationEncounterService HiddenLocationEncounterService
	KeyPointEncounterService       KeyPointEncounterService
	MiscEncounterService           MiscEncounterService
	SocialEncounterService         SocialEncounterService
	TouristProgressService         TouristProgressService
	DoneEncounterService           DoneEncounterService
	EncounterRepository            *repo.EncounterRepository
}

func (service *EncounterService) FindEncounter(encounterId string) (model.Encounter, error) {
	encounter1, err := service.HiddenLocationEncounterService.FindEncounter(encounterId)
	if err == nil {
		return encounter1.Encounter, nil
	}
	encounter2, err := service.KeyPointEncounterService.FindEncounter(encounterId)
	if err == nil {
		return encounter2.Encounter, nil
	}
	encounter3, err := service.MiscEncounterService.FindEncounter(encounterId)
	if err == nil {
		return encounter3.Encounter, nil
	}
	encounter4, err := service.SocialEncounterService.FindEncounter(encounterId)
	if err == nil {
		return encounter4.Encounter, nil
	}
	return model.Encounter{}, errors.New("no encounter found")
}

func (service *EncounterService) Activate(userId int, longitude float64, latitude float64, encounterId string) error {
	_, err := service.TouristProgressService.FindTouristProgress(fmt.Sprint(userId))
	if err != nil {
		err = service.TouristProgressService.Create(&model.TouristProgress{UserId: float64(userId), Xp: 0, Level: 0})
		if err != nil {
			return errors.New("unable to create tourist progress")
		}
	}

	encounterIdConverted, err := strconv.Atoi(encounterId)
	if err != nil {
		return fmt.Errorf("failed to convert encounterId to integer: %w", err)
	}

	encounter, err := service.FindEncounter(encounterId)
	if err != nil {
		return errors.New("encounter with given uuid does not exists")
	}

	_, err = service.DoneEncounterService.Find(fmt.Sprint(userId), encounterId)
	if err == nil {
		return errors.New("encounter already activated")
	}

	canActivate, err := encounter.CanActivate(longitude, latitude)
	if err != nil {
		return errors.New("error with activating encounter")
	}
	if canActivate {
		service.createDoneEncounter(userId, encounterIdConverted)
		return nil
	} else {
		return errors.New("can not activate encounter, you are too far")
	}
}

func (service *EncounterService) createDoneEncounter(userId int, encounterId int) error {
	err := service.DoneEncounterService.Create(&model.DoneEncounter{EncounterId: encounterId, UserId: userId, DoneEncounterStatus: model.DoneEncounterStatus(model.InProgress)})
	if err != nil {
		return err
	}
	println("Created done encounter")
	return nil
}

func (service *EncounterService) Complete(userId string, encounterId string) error {
	doneEncounter, err := service.DoneEncounterService.Find(userId, encounterId)
	if err != nil {
		return err
	}
	err = doneEncounter.Complete()
	if err != nil {
		return err
	}
	encounter, err := service.FindEncounter(encounterId)
	if err != nil {
		return err
	}
	err = service.TouristProgressService.GiveXp(userId, encounter.XpReward)
	if err != nil {
		return err
	}
	return service.DoneEncounterService.Save(doneEncounter)
}

func (service *EncounterService) GetEncounters() (*[]model.Encounter, error) {
	tours, err := service.EncounterRepository.GetEncounters()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", "-2"))
	}
	return &tours, nil
}

func (service *EncounterService) GetActive(userId string) (*[]model.Encounter, error) {
	var activeEncounters []model.Encounter
	doneEncounters, err := service.DoneEncounterService.GetActiveByUserId(userId)
	if err != nil {
		return &activeEncounters, err
	}

	for _, doneEncounter := range *doneEncounters {
		encounter, err := service.HiddenLocationEncounterService.FindEncounter(fmt.Sprint(doneEncounter.EncounterId))
		if err == nil {
			activeEncounters = append(activeEncounters, encounter.Encounter)
			continue
		}
		encounter1, err := service.SocialEncounterService.FindEncounter(fmt.Sprint(doneEncounter.EncounterId))
		if err == nil {
			activeEncounters = append(activeEncounters, encounter1.Encounter)
			continue
		}
		encounter2, err := service.MiscEncounterService.FindEncounter(fmt.Sprint(doneEncounter.EncounterId))
		if err == nil {
			activeEncounters = append(activeEncounters, encounter2.Encounter)
			continue
		}
		return &activeEncounters, errors.New("internal error, no encounter found for done encounter")
	}
	return &activeEncounters, nil
}

func (service *EncounterService) GetCompleted(userId string) (*[]model.Encounter, error) {
	var completedEncounters []model.Encounter
	doneEncounters, err := service.DoneEncounterService.GetCompletedByUserId(userId)
	if err != nil {
		return &completedEncounters, err
	}

	for _, doneEncounter := range *doneEncounters {
		encounter, err := service.HiddenLocationEncounterService.FindEncounter(fmt.Sprint(doneEncounter.EncounterId))
		if err == nil {
			completedEncounters = append(completedEncounters, encounter.Encounter)
			continue
		}
		encounter1, err := service.SocialEncounterService.FindEncounter(fmt.Sprint(doneEncounter.EncounterId))
		if err == nil {
			completedEncounters = append(completedEncounters, encounter1.Encounter)
			continue
		}
		encounter2, err := service.MiscEncounterService.FindEncounter(fmt.Sprint(doneEncounter.EncounterId))
		if err == nil {
			completedEncounters = append(completedEncounters, encounter2.Encounter)
			continue
		}
		return &completedEncounters, errors.New("internal error, no encounter found for done encounter")
	}
	return &completedEncounters, nil
}

func (service *EncounterService) Cancel(userId string, encounterId string) error {
	return service.DoneEncounterService.Delete(userId, encounterId)
}

func (service *EncounterService) IsCompleted(userId string, encounterId string) bool {
	return service.DoneEncounterService.IsCompleted(userId, encounterId)
}
