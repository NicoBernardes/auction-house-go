package api

import (
	"net/http"

	"github.com/NicoBernardes/auction-house-go.git/internal/services"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
)

type Api struct {
	Router      *chi.Mux
	UserService services.UserService
	Sessions    *scs.SessionManager
}

func (api *Api) handleCreateUser(w http.ResponseWriter, r *http.Request) {}
