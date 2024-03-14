package handler

import (
	"encoding/json"
	"encounters/service"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type EncounterHandler struct {
	EncounterService service.EncounterService
}

func (handler *EncounterHandler) Activate(writer http.ResponseWriter, req *http.Request) {
	var position struct {
		TouristId int
		Longitude int
		Latitude  int
	}
	encounterId := mux.Vars(req)["id"]
	err := json.NewDecoder(req.Body).Decode(&position)
	if err != nil {
		fmt.Println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	encounterUUID, err := uuid.FromBytes([]byte(encounterId))
	if err != nil {
		fmt.Println("Error while creating uuid")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.EncounterService.Activate(position.TouristId, position.Longitude, position.Latitude, encounterUUID)
	if err != nil {
		fmt.Println("Error while activating encounter")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)

}
