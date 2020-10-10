package controllers

import (
	"encoding/json"
	"net/http"

	"log"

	"../common"
	"../models"
	"github.com/gorilla/mux" // URL router and dispatcher.
)

var Tasks = new(tasksController)

type tasksController struct{}

// POST
func (tc *tasksController) Create(w http.ResponseWriter, r *http.Request) {
	var t models.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	task, err := models.Tasks.Create(t.Name, t.Desc)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(task)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonOk(w, res, http.StatusCreated)
}

// GET
func (tc *tasksController) Get(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.Tasks.FindAll()
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(tasks)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonOk(w, res, http.StatusOK)
}

// GET
func (tc *tasksController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	task, err := models.Tasks.FindOne(id)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(task)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonOk(w, res, http.StatusOK)
}

// PUT
func (tc *tasksController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var t models.Task
	// verifica a consistencia do request
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		log.Println("ERRO1===>" + err.Error())
		return
	}

	// faz o update propriamente dito
	if err := models.Tasks.Update(id, t.Name, t.Desc); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		log.Println("ERRO2===>" + err.Error())
		return
	}

	common.JsonStatus(w, http.StatusNoContent)
}

// DELETE
func (tc *tasksController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := models.Tasks.DeleteById(id); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	common.JsonStatus(w, http.StatusNoContent)
}
