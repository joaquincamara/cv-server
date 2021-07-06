package api

import (
	"encoding/json"
	"net/http"

	"github.com/joaquincamara/cv-server/internal/devTechs"
)

type IDevTechHandler interface {
	Post(http.ResponseWriter, *http.Request)
}

type handler struct {
	devTechService devTechs.Service
}

func NewDevTechHandler(devTechService devTechs.Service) IDevTechHandler {
	return &handler{devTechService: devTechService}
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	p := &devTechs.DevTech{}
	err := json.NewDecoder(r.Body).Decode(p)

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
	json.NewEncoder(w).Encode(p)
}
