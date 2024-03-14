package service

import (
	"encounters/model"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type EncounterService struct {
	HiddenLocationEncounterService HiddenLocationEncounterService
	KeyPointEncounterService       KeyPointEncounterService
	MiscEncounterService           MiscEncounterService
	SocialEncounterService         SocialEncounterService
	TouristProgressService         TouristProgressService
	DoneEncounterService           DoneEncounterService
}

func (service *EncounterService) FindEncounter(encounterId uuid.UUID) (model.Encounter, error) {
	encounter1, err := service.HiddenLocationEncounterService.FindEncounter(encounterId.String())
	if err == nil {
		return encounter1.Encounter, nil
	}
	encounter2, err := service.KeyPointEncounterService.FindEncounter(encounterId.String())
	if err == nil {
		return encounter2.Encounter, nil
	}
	encounter3, err := service.MiscEncounterService.FindEncounter(encounterId.String())
	if err == nil {
		return encounter3.Encounter, nil
	}
	encounter4, err := service.SocialEncounterService.FindEncounter(encounterId.String())
	if err == nil {
		return encounter4.Encounter, nil
	}
	return model.Encounter{}, errors.New("no encounter found")
}

func (service *EncounterService) Activate(userId int, longitude int, latitude int, encounterId uuid.UUID) error {
	_, err := service.TouristProgressService.FindTouristProgress(fmt.Sprint(userId))
	if err != nil {
		err = service.TouristProgressService.Create(&model.TouristProgress{UserId: float64(userId), Xp: 0, Level: 0})
		if err != nil {
			return errors.New("unable to create tourist progress")
		}
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
			miscEncounter, err := service.MiscEncounterService.FindEncounter(encounter.ID.String())
			if err != nil {
				return errors.New("error searching for misc encounter")
			}
			err = service.MiscEncounterService.Save(miscEncounter)
			if err != nil {
				return errors.New("error saving misc encounter")
			}
			service.createDoneEncounter(userId, encounterId)
			return nil
		}
		if encounter.Status == model.EncounterStatus(model.Social) {
			socialEncounter, err := service.SocialEncounterService.FindEncounter(encounter.ID.String())
			if err != nil {
				return errors.New("error searching for social encounter")
			}
			err = service.SocialEncounterService.Save(socialEncounter)
			if err != nil {
				return errors.New("error saving social encounter")
			}
			service.createDoneEncounter(userId, encounterId)
			return nil
		}
		if encounter.Status == model.EncounterStatus(model.KeyPoint) {
			keyPointEncounter, err := service.KeyPointEncounterService.FindEncounter(encounter.ID.String())
			if err != nil {
				return errors.New("error searching for key point encounter")
			}
			err = service.KeyPointEncounterService.Save(keyPointEncounter)
			if err != nil {
				return errors.New("error saving key point encounter")
			}
			service.createDoneEncounter(userId, encounterId)
			return nil
		}
		if encounter.Status == model.EncounterStatus(model.Hidden) {
			hiddenEncounter, err := service.HiddenLocationEncounterService.FindEncounter(encounter.ID.String())
			if err != nil {
				return errors.New("error searching for hidden encounter")
			}
			err = service.HiddenLocationEncounterService.Save(hiddenEncounter)
			if err != nil {
				return errors.New("error saving hidden encounter")
			}
			service.createDoneEncounter(userId, encounterId)
			return nil
		}
	}
	return errors.New("unknown error")
}

func (service *EncounterService) createDoneEncounter(userId int, encounterId uuid.UUID) error {
	err := service.DoneEncounterService.Create(&model.DoneEncounter{EncounterId: encounterId, UserId: userId, DoneEncounterStatus: model.DoneEncounterStatus(model.InProgress)})
	if err != nil {
		return err
	}
	println("Created done encounter")
	return nil
}
