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

func (service *EncounterService) Activate(userId int, longitude int, latitude int, encounterId string) error {
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

	canActivate, err := encounter.CanActivate(longitude, latitude)
	if err != nil {
		return errors.New("error with activating encounter")
	}
	if canActivate {
		if encounter.Status == model.EncounterStatus(model.Misc) {
			miscEncounter, err := service.MiscEncounterService.FindEncounter(fmt.Sprint(encounter.ID))
			if err != nil {
				return errors.New("error searching for misc encounter")
			}
			err = service.MiscEncounterService.Save(miscEncounter)
			if err != nil {
				return errors.New("error saving misc encounter")
			}
			service.createDoneEncounter(userId, encounterIdConverted)
			return nil
		}
		if encounter.Status == model.EncounterStatus(model.Social) {
			socialEncounter, err := service.SocialEncounterService.FindEncounter(fmt.Sprint(encounter.ID))
			if err != nil {
				return errors.New("error searching for social encounter")
			}
			err = service.SocialEncounterService.Save(socialEncounter)
			if err != nil {
				return errors.New("error saving social encounter")
			}
			service.createDoneEncounter(userId, encounterIdConverted)
			return nil
		}
		if encounter.Status == model.EncounterStatus(model.KeyPoint) {
			keyPointEncounter, err := service.KeyPointEncounterService.FindEncounter(fmt.Sprint(encounter.ID))
			if err != nil {
				return errors.New("error searching for key point encounter")
			}
			err = service.KeyPointEncounterService.Save(keyPointEncounter)
			if err != nil {
				return errors.New("error saving key point encounter")
			}
			service.createDoneEncounter(userId, encounterIdConverted)
			return nil
		}
		if encounter.Status == model.EncounterStatus(model.Hidden) {
			hiddenEncounter, err := service.HiddenLocationEncounterService.FindEncounter(fmt.Sprint(encounter.ID))
			if err != nil {
				return errors.New("error searching for hidden encounter")
			}
			err = service.HiddenLocationEncounterService.Save(hiddenEncounter)
			if err != nil {
				return errors.New("error saving hidden encounter")
			}
			service.createDoneEncounter(userId, encounterIdConverted)
			return nil
		}
	}
	return errors.New("unknown error")
}

func (service *EncounterService) createDoneEncounter(userId int, encounterId int) error {
	err := service.DoneEncounterService.Create(&model.DoneEncounter{EncounterId: encounterId, UserId: userId, DoneEncounterStatus: model.DoneEncounterStatus(model.InProgress)})
	if err != nil {
		return err
	}
	println("Created done encounter")
	return nil
}

func (service *EncounterService) GetEncounters() (*[]model.Encounter, error) {
	tours, err := service.EncounterRepository.GetEncounters()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", "-2"))
	}
	return &tours, nil
}
