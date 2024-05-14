package adapters

import (
	"bashscripts/internal/models"
	"context"
	"github.com/gorilla/mux"
)

type HTTPHandle struct {
	hm HTTPManager
}

func NewHTTPHandle(hm HTTPManager) *HTTPHandle {
	return &HTTPHandle{
		hm: hm,
	}
}

type HTTPManager interface {
	CreateScript(context.Context, *models.Script) error
	GetScriptsList(context.Context) ([]*models.Script, error)
	GetScript(context.Context, string) (*models.Script, error)
}

func (h *HTTPHandle) Register(router *mux.Router) {
	router.HandleFunc("/create", h.CreateScript).Methods("POST")
	router.HandleFunc("/getlist", h.GetScriptsList).Methods("GET")
	router.HandleFunc("/getScript", h.GetScript).Methods("GET")
}
