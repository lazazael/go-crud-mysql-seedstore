package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lazazael/go-crud-mysql-seedstore/pkg/models"
	"github.com/lazazael/go-crud-mysql-seedstore/pkg/utils"
	"net/http"
	"strconv"
)

var NewSeed models.Seed

func GetSeed(w http.ResponseWriter, r *http.Request) {
	newSeeds := models.GetAllSeeds()
	res, _ := json.Marshal(newSeeds)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetSeedById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	seedId := vars["seedId"]
	ID, err := strconv.ParseInt(seedId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")

	}
	seedDetails, _ := models.GetSeedById(ID)
	res, _ := json.Marshal(seedDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateSeed(w http.ResponseWriter, r *http.Request) {
	CreateSeed := &models.Seed{}
	utils.ParseBody(r, CreateSeed)
	s := CreateSeed.CreateSeed()
	res, _ := json.Marshal(s)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteSeed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	seedId := vars["seedId"]
	ID, err := strconv.ParseInt(seedId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	seed := models.DeleteSeed(ID)
	res, _ := json.Marshal(seed)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateSeed(w http.ResponseWriter, r *http.Request) {
	var updateSeed = &models.Seed{}
	utils.ParseBody(r, updateSeed)

	vars := mux.Vars(r)
	seedId := vars["seedId"]

	ID, err := strconv.ParseInt(seedId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	seedDetails, db := models.GetSeedById(ID)
	if updateSeed.SpeciesName != "" {
		seedDetails.SpeciesName = updateSeed.SpeciesName
	}
	if updateSeed.VarietyName != "" {
		seedDetails.VarietyName = updateSeed.VarietyName
	}
	if updateSeed.PacketContent != 0 {
		seedDetails.PacketContent = updateSeed.PacketContent
	}
	db.Save(&seedDetails)
	res, _ := json.Marshal(seedDetails)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)

}
