package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/open-source-publications/cqrs-golang-core/model"
	"github.com/open-source-publications/cqrs-golang-core/repository"
	"gopkg.in/mgo.v2/bson"
)

var eventStormingRespository = repository.EventStormingRepository{}

func respondWithERROR(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//Save in database
func Save(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var eventStormingModel model.EventStormingModel

	if err := json.NewDecoder(r.Body).Decode(&eventStormingModel); err != nil {
		respondWithERROR(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	eventStormingModel.ID = bson.NewObjectId()
	eventStormingModel.CreatedAt = time.Now()
	eventStormingModel.UpdatedAt = time.Now()

	if err := eventStormingRespository.Save(eventStormingModel); err != nil {
		respondWithERROR(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, eventStormingModel)
}

//Delete a eventstorming of database
func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// gets params sended
	params := mux.Vars(r)

	if err := eventStormingRespository.Delete(params["id"]); err != nil {
		respondWithERROR(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

//Update a eventstorming in database
func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// get params sended
	params := mux.Vars(r)

	var eventStormingModel model.EventStormingModel

	if err := json.NewDecoder(r.Body).Decode(&eventStormingModel); err != nil {
		respondWithERROR(w, http.StatusBadRequest, "Invalid request Payload")
		return
	}

	eventStormingModel.UpdatedAt = time.Now()

	if err := eventStormingRespository.Update(params["id"], eventStormingModel); err != nil {
		respondWithERROR(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

//GetByID a eventstorming
func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	eventStormingModel, err := eventStormingRespository.GetByID(params["id"])

	if err != nil {
		respondWithERROR(w, http.StatusBadRequest, "Invalid EventStorming ID")
	}

	respondWithJSON(w, http.StatusOK, eventStormingModel)

}

//GetAll elements in database
func GetAll(w http.ResponseWriter, r *http.Request) {
	eventStorminModel, err := eventStormingRespository.GetAll()

	if err != nil {
		respondWithERROR(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, eventStorminModel)
}
