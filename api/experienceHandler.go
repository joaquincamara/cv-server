package api

import (
	"encoding/json"
	"net/http"

	"github.com/joaquincamara/cv-server/internal/experience"
)

type IExperienceHandler interface {
	Post(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
}

type experienceHandler struct {
	experienceService experience.Service
}

func NewExperienceHandler(experienceService experience.Service) IExperienceHandler {
	return &experienceHandler{experienceService: experienceService}
}

func (e *experienceHandler) Post(w http.ResponseWriter, r *http.Request) {
	p := &experience.Experience{}
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = e.experienceService.Add(p)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&p)
}

func (e *experienceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	type id struct {
		Id int
	}
	devTechId := &id{}
	err := json.NewDecoder(r.Body).Decode(devTechId)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	err = e.experienceService.Delete(devTechId.Id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(devTechId)
}

func (e *experienceHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	res, err := e.experienceService.FindAll()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&res)
}

func (e *experienceHandler) Put(w http.ResponseWriter, r *http.Request) {

	p := &experience.Experience{}
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = e.experienceService.Update(p)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&p)

}
