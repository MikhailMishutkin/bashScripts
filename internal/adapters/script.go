package adapters

import (
	"bashscripts/internal/models"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

// ...
func (h *HTTPHandle) CreateScript(w http.ResponseWriter, r *http.Request) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	var script *models.Script
	err = json.Unmarshal(content, &script)
	if err != nil {
		w.Write([]byte("corrupt json data" + err.Error()))
	}
	err = h.hm.CreateScript(context.Background(), script)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
	}
}

// ...
func (h *HTTPHandle) GetScript(w http.ResponseWriter, r *http.Request) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	var script *models.Script
	err = json.Unmarshal(content, &script)
	if err != nil {
		w.Write([]byte("corrupt json data" + err.Error()))
	}

	script, err = h.hm.GetScript(context.Background(), script.Name)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))

	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Connection:", "close")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		err = json.NewEncoder(w).Encode(script)

	}
}

// ...
func (h *HTTPHandle) GetScriptsList(w http.ResponseWriter, r *http.Request) {

	list, err := h.hm.GetScriptsList(context.Background())

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))

	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Connection:", "close")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		err = json.NewEncoder(w).Encode(list)

	}
}
