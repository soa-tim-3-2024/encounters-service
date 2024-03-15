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

	/*encounterUUID, err := uuid.FromBytes([]byte(encounterId))
	if err != nil {
		fmt.Println("Error while creating uuid")
		writer.WriteHeader(http.StatusBadRequest)
		return
	} */
	err = handler.EncounterService.Activate(position.TouristId, position.Longitude, position.Latitude, encounterId)
	if err != nil {
		fmt.Println("Error while activating encounter")
		writer.WriteHeader(http.StatusInternalServerError)
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
