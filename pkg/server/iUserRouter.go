package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jcasado94/gobuyright/pkg/entity"

	"github.com/gorilla/mux"
)

type iUserRouter struct {
	iUserService entity.IUserService
}

// NewIUserRouter adds an iUserRouter configuration to router.
func NewIUserRouter(s entity.IUserService, router *mux.Router) *mux.Router {
	userRouter := iUserRouter{s}

	router.HandleFunc("/", userRouter.createUserHandler).Methods("PUT")
	router.HandleFunc("/{username}", userRouter.getUserHandler).Methods("GET")

	return router
}

func (ur *iUserRouter) createUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := decodeUser(r)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = ur.iUserService.CreateUser(&user)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, err)
}

func (ur *iUserRouter) getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	user, err := ur.iUserService.GetByUsername(username)
	if err != nil {
		WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, user)
}

func decodeUser(r *http.Request) (entity.IUser, error) {
	var u entity.IUser
	if r.Body == nil {
		return u, errors.New("No request body")
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)

	return u, err
}
