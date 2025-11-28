package api

import (
	"errors"
	"net/http"

	"github.com/NicoBernardes/auction-house-go.git/internal/api/usecase/user"
	"github.com/NicoBernardes/auction-house-go.git/internal/jsonutils"
	"github.com/NicoBernardes/auction-house-go.git/internal/services"
)

func (api *Api) handleSignupUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[user.CreateUserReq](r)
	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}
	id, err := api.UserService.CreateUser(r.Context(),
		data.UserName,
		data.Email,
		data.Password,
		data.Bio,
	)
	if err != nil {
		if errors.Is(err, services.ErrDuplicatedEmailOrUsername) {
			_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{
				"error": "email or username already exists",
			})
			return
		}
	}
	_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{
		"user_id": id,
	})
}

func (api *Api) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	panic("TODO - NOT IMPLEMENTED")
}

func (api *Api) handleLogoutUser(w http.ResponseWriter, r *http.Request) {
	panic("TODO - NOT IMPLEMENTED")
}
