package controllers

import (
	"encoding/json"
	"net/http"

	"main.go/commons"
	"main.go/models"
)

func getAll(writer http.ResponseWriter, request *http.Request) {
	people := []models.People{}
	db := commons.Getconnection()
	defer db.Close()
	db.Find(&people)

	json, _ := json.Marshal(people)

	commons.SendResponse(writer, http.StatusOK, json)

}
