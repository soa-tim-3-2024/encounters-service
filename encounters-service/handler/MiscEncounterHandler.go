package handler

import (
	"bytes"
	"encoding/json"
	"encounters/model"
	"encounters/service"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type MiscEncounterHandler struct {
	MiscEncounterService *service.MiscEncounterService
}

func (handler *MiscEncounterHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Encounter sa id-em %s", id)
	encounter, err := handler.MiscEncounterService.FindEncounter(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(encounter)
}

func (handler *MiscEncounterHandler) Create(writer http.ResponseWriter, req *http.Request) {
	bodyBytes, err := io.ReadAll(io.TeeReader(req.Body, &bytes.Buffer{}))
	if err != nil {
		println("Error reading request body")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var encounter model.MiscEncounter
	err = json.Unmarshal(bodyBytes, &encounter)
	if err != nil {
		println("Error while parsing json ")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var baseEncounter model.Encounter
	err2 := json.Unmarshal(bodyBytes, &baseEncounter)
	if err2 != nil {
		println("Error while parsing json base")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	encounter.EncounterID = baseEncounter.ID
	encounter.Encounter = baseEncounter

	err = handler.MiscEncounterService.Create(&encounter)
	if err != nil {
		println("Error while creating a new encounter")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
