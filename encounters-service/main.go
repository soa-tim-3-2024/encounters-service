package main

import (
	"encounters/handler"
	"encounters/model"
	"encounters/repo"
	"encounters/service"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=super dbname=gormen port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.Student{})

	database.AutoMigrate(&model.Encounter{})
	database.AutoMigrate(&model.MiscEncounter{})
	database.AutoMigrate(&model.SocialEncounter{})
	database.AutoMigrate(&model.KeyPointEncounter{})
	database.AutoMigrate(&model.HiddenLocationEncounter{})
	database.AutoMigrate(&model.TouristProgress{})

	database.Exec("INSERT INTO tourist_progresses VALUES ('aec7e123-233d-4a09-a289-75308ea5b7e6', '-4', '85', '12')")

	return database
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	studentRepo := &repo.StudentRepository{DatabaseConnection: database}
	studentService := &service.StudentService{StudentRepo: studentRepo}
	studentHandler := &handler.StudentHandler{StudentService: studentService}

	miscEncounterRepo := &repo.MiscEncounterRepository{DatabaseConnection: database}
	miscEncounterService := &service.MiscEncounterService{EncounterRepo: miscEncounterRepo}
	miscEncounterHandler := &handler.MiscEncounterHandler{MiscEncounterService: miscEncounterService}

	socialEncounterRepo := &repo.SocialEncounterRepository{DatabaseConnection: database}
	socialEncounterService := &service.SocialEncounterService{EncounterRepo: socialEncounterRepo}
	socialEncounterHandler := &handler.SocialEncounterHandler{SocialEncounterService: socialEncounterService}

	keyPointEncounterRepo := &repo.KeyPointEncounterRepository{DatabaseConnection: database}
	keyPointEncounterService := &service.KeyPointEncounterService{EncounterRepo: keyPointEncounterRepo}
	keyPointEncounterHandler := &handler.KeyPointEncounterHandler{KeyPointEncounterService: keyPointEncounterService}

	hiddenLocationEncounterRepo := &repo.HiddenLocationEncounterRepository{DatabaseConnection: database}
	hiddenLocationEncounterService := &service.HiddenLocationEncounterService{EncounterRepo: hiddenLocationEncounterRepo}
	hiddenLocationEncounterHandler := &handler.HiddenLocationEncounterHandler{HiddenLocationEncounterService: hiddenLocationEncounterService}

	touristProgressRepo := &repo.TouristProgressRepository{DatabaseConnection: database}
	touristProgressService := &service.TouristProgressService{TPRepo: touristProgressRepo}
	touristProgressHandler := &handler.TouristProgressHandler{TouristProgressService: touristProgressService}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/students/{id}", studentHandler.Get).Methods("GET")
	router.HandleFunc("/students", studentHandler.Create).Methods("POST")
	router.HandleFunc("/misc/encounters/{id}", miscEncounterHandler.Get).Methods("GET")
	router.HandleFunc("/misc/encounters", miscEncounterHandler.Create).Methods("POST")
	router.HandleFunc("/social/encounters/{id}", socialEncounterHandler.Get).Methods("GET")
	router.HandleFunc("/social/encounters", socialEncounterHandler.Create).Methods("POST")
	router.HandleFunc("/hidden/location/encounters/{id}", hiddenLocationEncounterHandler.Get).Methods("GET")
	router.HandleFunc("/hidden/location/encounters", hiddenLocationEncounterHandler.Create).Methods("POST")
	router.HandleFunc("/keyPoint/encounters/{id}", keyPointEncounterHandler.Get).Methods("GET")
	router.HandleFunc("/keyPoint/encounters", keyPointEncounterHandler.Create).Methods("POST")

	router.HandleFunc("/progress/{id}", touristProgressHandler.Get).Methods("GET")

	// Set up CORS middleware
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8082", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)))
}
