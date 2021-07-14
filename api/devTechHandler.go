package api

import (
	"encoding/json"
	"net/http"

	"github.com/joaquincamara/cv-server/internal/devTechs"
)

type IDevTechHandler interface {
	Post(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
}

type devTechHandler struct {
	devTechService devTechs.Service
}

func NewDevTechHandler(devTechService devTechs.Service) IDevTechHandler {
	return &devTechHandler{devTechService: devTechService}
}

func (h *devTechHandler) Post(w http.ResponseWriter, r *http.Request) {
	p := &devTechs.DevTech{}
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = h.devTechService.Add(p)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&p)
}

func (h *devTechHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	err = h.devTechService.Delete(devTechId.Id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(devTechId)
}

func (h *devTechHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	res, err := h.devTechService.FindAll()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&res)
}

func (h *devTechHandler) Put(w http.ResponseWriter, r *http.Request) {

	p := &devTechs.DevTech{}
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = h.devTechService.Update(p)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&p)

}
