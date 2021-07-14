package api

import (
	"encoding/json"
	"net/http"

	"github.com/joaquincamara/cv-server/internal/personalProjects"
)

type IPersonalProjectHandler interface {
	Post(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
}

type personalProjectHandler struct {
	personalProjectService personalProjects.Service
}

func NewPersonalProjectHandler(personalProjectService personalProjects.Service) IPersonalProjectHandler {
	return &personalProjectHandler{personalProjectService: personalProjectService}
}

func (p *personalProjectHandler) Post(w http.ResponseWriter, r *http.Request) {
	pp := &personalProjects.PersonalProjects{}
	err := json.NewDecoder(r.Body).Decode(&pp)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = p.personalProjectService.Add(pp)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&pp)
}

func (p *personalProjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
	type id struct {
		Id int
	}
	personalProjectId := &id{}
	err := json.NewDecoder(r.Body).Decode(personalProjectId)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	err = p.personalProjectService.Delete(personalProjectId.Id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(personalProjectId)
}

func (p *personalProjectHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	res, err := p.personalProjectService.FindAll()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&res)
}

func (p *personalProjectHandler) Put(w http.ResponseWriter, r *http.Request) {

	pp := &personalProjects.PersonalProjects{}
	err := json.NewDecoder(r.Body).Decode(&pp)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = p.personalProjectService.Update(pp)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&pp)

}
