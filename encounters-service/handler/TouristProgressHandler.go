package handler

import (
	"encoding/json"
	"encounters/model"
	"encounters/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TouristProgressHandler struct {
	TouristProgressService *service.TouristProgressService
}

func (handler *TouristProgressHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Progress sa id-em %s", id)
	progress, err := handler.TouristProgressService.FindTouristProgress(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(progress)
}

func (handler *TouristProgressHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var touristProgress model.TouristProgress
	err := json.NewDecoder(req.Body).Decode(&touristProgress)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.TouristProgressService.Create(&touristProgress)
	if err != nil {
		println("Error while creating a new progress")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
