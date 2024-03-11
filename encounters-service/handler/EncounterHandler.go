package handler

import (
	"encoding/json"
	"encounters/model"
	"encounters/service"
	"fmt"
	"net/http"
)

type EncounterHandler struct {
	EncounterService *service.EncounterService
}

func (handler *EncounterHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var encounter model.Encounter
	err := json.NewDecoder(req.Body).Decode(&encounter)
	fmt.Printf("%+v\n", encounter)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.EncounterService.Update(&encounter)
	if err != nil {
		println("Error while updating tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounter)
}

func (handler *EncounterHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var encounter model.Encounter
	err := json.NewDecoder(req.Body).Decode(&encounter)
	fmt.Printf("%+v\n", encounter)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.EncounterService.Create(&encounter)
	if err != nil {
		println("Error while creating a new tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounter)
}
