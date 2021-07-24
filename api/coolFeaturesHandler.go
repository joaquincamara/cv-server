package api

import (
	"encoding/json"
	"net/http"

	coolfeatures "github.com/joaquincamara/cv-server/internal/coolFeatures"
)

type ICoolFeaturesHandler interface {
	Post(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
}

type coolFeaturesHandler struct {
	coolFeaturesService coolfeatures.Service
}

func NewCoolFeaturesHandler(coolFeaturesServic coolfeatures.Service) ICoolFeaturesHandler {
	return &coolFeaturesHandler{coolFeaturesService: coolFeaturesServic}
}

func (c *coolFeaturesHandler) Post(w http.ResponseWriter, r *http.Request) {
	p := &coolfeatures.Coolfeatures{}
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = c.coolFeaturesService.Add(p)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&p)
}

func (c *coolFeaturesHandler) Delete(w http.ResponseWriter, r *http.Request) {
	type id struct {
		Id int
	}
	coolFeatureId := &id{}
	err := json.NewDecoder(r.Body).Decode(coolFeatureId)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	err = c.coolFeaturesService.Delete(coolFeatureId.Id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(coolFeatureId)
}

func (c *coolFeaturesHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	res, err := c.coolFeaturesService.FindAll()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&res)
}

func (c *coolFeaturesHandler) Put(w http.ResponseWriter, r *http.Request) {

	p := &coolfeatures.Coolfeatures{}
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = c.coolFeaturesService.Update(p)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&p)

}
