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
	"time"

	"github.com/gorilla/mux"
)

type DoneEncounterHandler struct {
	DoneEncounterService *service.DoneEncounterService
}

func (handler *DoneEncounterHandler) Get(writer http.ResponseWriter, req *http.Request) {
	userId := mux.Vars(req)["userId"]
	log.Printf("DoneEncounter sa userId-em %s", userId)
	encounter, err := handler.DoneEncounterService.FindByUserId(userId)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(encounter)
}

func (handler *DoneEncounterHandler) Create(writer http.ResponseWriter, req *http.Request) {
	bodyBytes, err := io.ReadAll(io.TeeReader(req.Body, &bytes.Buffer{}))
	if err != nil {
		println("Error reading request body")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var doneEncounter model.DoneEncounter

	err = json.Unmarshal(bodyBytes, &doneEncounter)
	if err != nil {
		println("Error while parsing json ")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	doneEncounter.CompletionTime = time.Now()
	fmt.Println(doneEncounter)

	err = handler.DoneEncounterService.Create(&doneEncounter)
	if err != nil {
		println("Error while creating a new done encounter")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
