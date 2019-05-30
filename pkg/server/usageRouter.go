package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jcasado94/gobuyright/pkg/entity"
)

type usageRouter struct {
	usagesService entity.UsageService
}

// NewUsageRouter adds a usageRouter configuration to router.
func NewUsageRouter(s entity.UsageService, router *mux.Router) *mux.Router {
	ur := usageRouter{s}

	router.HandleFunc("/all", ur.getAllUsagesHandler).Methods("GET")

	return router
}

func (ur *usageRouter) getAllUsagesHandler(w http.ResponseWriter, r *http.Request) {
	usages, err := ur.usagesService.GetAllUsages()
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, usages)
}
