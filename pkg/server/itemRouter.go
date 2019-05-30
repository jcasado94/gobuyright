package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jcasado94/gobuyright/pkg/entity"
)

type itemRouter struct {
	itemService entity.ItemService
}

// NewItemRouter adds an itemRouter configuration to router.
func NewItemRouter(s entity.ItemService, router *mux.Router) *mux.Router {
	ir := itemRouter{s}

	router.HandleFunc("/all", ir.getAllItemsHandler).Methods("GET")

	return router
}

func (ir *itemRouter) getAllItemsHandler(w http.ResponseWriter, r *http.Request) {
	items, err := ir.itemService.GetAllItems()
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, items)
}
