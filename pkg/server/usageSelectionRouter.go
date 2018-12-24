package server

import (
	"encoding/json"
	"errors"
	"gobuyright/pkg/entity"
	"net/http"

	"github.com/gorilla/mux"
)

type usageSelectionRouter struct {
	usageSelectionService entity.UsageSelectionService
}

// NewUsageSelectionRouter creates a new usageSelectionRouter with the provided Service and Router.
func NewUsageSelectionRouter(uss entity.UsageSelectionService, router *mux.Router) *mux.Router {
	usr := usageSelectionRouter{uss}

	router.HandleFunc("/", usr.createUsageSelectionHandler).Methods("PUT")

	return router
}

func (usr *usageSelectionRouter) createUsageSelectionHandler(w http.ResponseWriter, r *http.Request) {
	us, err := decodeUsageSelection(r)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid request payload")
	}

	err = usr.usageSelectionService.CreateUsageSelection(&us)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, err)
}

func decodeUsageSelection(r *http.Request) (entity.UsageSelection, error) {
	var us entity.UsageSelection
	if r.Body == nil {
		return us, errors.New("No request body")
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&us)

	return us, err
}
