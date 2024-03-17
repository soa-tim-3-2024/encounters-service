package handler

import (
	"encoding/json"
	"encounters/service"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type EncounterHandler struct {
	EncounterService service.EncounterService
}

func (handler *EncounterHandler) Activate(writer http.ResponseWriter, req *http.Request) {
	var position struct {
		TouristId int
		Longitude float64
		Latitude  float64
	}
	encounterId := mux.Vars(req)["id"]
	err := json.NewDecoder(req.Body).Decode(&position)
	if err != nil {
		fmt.Println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	/*encounterUUID, err := uuid.FromBytes([]byte(encounterId))
	if err != nil {
		fmt.Println("Error while creating uuid")
		writer.WriteHeader(http.StatusBadRequest)
		return
	} */
	err = handler.EncounterService.Activate(position.TouristId, position.Longitude, position.Latitude, encounterId)
	if err != nil {
		fmt.Println("Error while activating encounter")
		fmt.Println(err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (handler *EncounterHandler) Get(writer http.ResponseWriter, req *http.Request) {
	encounters, err := handler.EncounterService.GetEncounters()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(encounters)
}

func (handler *EncounterHandler) Cancel(writer http.ResponseWriter, req *http.Request) {
	userId := mux.Vars(req)["userId"]
	encounterId := mux.Vars(req)["encounterId"]
	err := handler.EncounterService.Cancel(userId, encounterId)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (handler *EncounterHandler) GetCompletedByUser(writer http.ResponseWriter, req *http.Request) {
	userId := mux.Vars(req)["userId"]
	encounters, err := handler.EncounterService.GetCompleted(userId)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
	println(*encounters)
	json.NewEncoder(writer).Encode(*encounters)
}

func (handler *EncounterHandler) IsCompleted(writer http.ResponseWriter, req *http.Request) {
	userId := mux.Vars(req)["userId"]
	encounterId := mux.Vars(req)["encounterId"]
	completed := handler.EncounterService.IsCompleted(userId, encounterId)
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(completed)
}

func (handler *EncounterHandler) Complete(writer http.ResponseWriter, req *http.Request) {
	userId := mux.Vars(req)["userId"]
	encounterId := mux.Vars(req)["encounterId"]
	err := handler.EncounterService.Complete(userId, encounterId)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
