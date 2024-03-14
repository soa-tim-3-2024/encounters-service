package handler

import (
	"bytes"
	"encoding/json"
	"encounters/model"
	"encounters/service"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HiddenLocationEncounterHandler struct {
	HiddenLocationEncounterService *service.HiddenLocationEncounterService
}

func (handler *HiddenLocationEncounterHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Encounter sa id-em %s", id)
	encounter, err := handler.HiddenLocationEncounterService.FindEncounter(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(encounter)
}

func (handler *HiddenLocationEncounterHandler) Create(writer http.ResponseWriter, req *http.Request) {
	bodyBytes, err := io.ReadAll(io.TeeReader(req.Body, &bytes.Buffer{}))
	if err != nil {
		println("Error reading request body")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var encounter model.HiddenLocationEncounter

	err = json.Unmarshal(bodyBytes, &encounter)
	if err != nil {
		println("Error while parsing json ")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(encounter)
	var baseEncounter model.Encounter
	json.Unmarshal(bodyBytes, &baseEncounter)
	fmt.Println(baseEncounter)
	encounter.EncounterID = baseEncounter.ID
	encounter.Encounter = baseEncounter
	encounter.Encounter.Type = 1

	err = handler.HiddenLocationEncounterService.Create(&encounter)
	if err != nil {
		println("Error while creating a new encounter")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
