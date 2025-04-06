package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maggickkw/ecom/types"
	"github.com/maggickkw/ecom/utils"
)

type Handler struct{
	store *types.UserStore
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/login", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {\
// get JSON payload
var payload types.RegisterUserPayload
if err := utils.ParseJSON(r, payload) err != nil {
	utils.WriteError(w, http.StatusBadRequest, err)
}
//check if the user exists 


// if it doesn't , we create a new user
}
