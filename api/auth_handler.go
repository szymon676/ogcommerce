package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/szymon676/ogcommerce/store"
	"github.com/szymon676/ogcommerce/types"
)

const secret = "secret"

type AuthHandler struct {
	store      store.UsersStorager
	jwtservice JwtService
}

func NewAuthHandler(astore store.UsersStorager, jwtservice JwtService) *AuthHandler {
	return &AuthHandler{
		store:      astore,
		jwtservice: jwtservice,
	}
}

func (h *AuthHandler) handleLoginUser(w http.ResponseWriter, r *http.Request) error {
	var reqUser types.ReqUser

	if err := json.NewDecoder(r.Body).Decode(&reqUser); err != nil {
		return err
	}

	if err := h.ValidateUser(r.Context(), reqUser, w); err != nil {
		return err
	}

	token, err := h.jwtservice.CreateJWT(&reqUser, r.Context())

	if err != nil {
		return err
	}

	if err := CreateCookie(w, token); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, "user logged in successfully")
}

func (h AuthHandler) ValidateUser(ctx context.Context, vuser types.ReqUser, w http.ResponseWriter) error {
	user, err := h.store.GetUserByEmail(ctx, vuser.Email)
	if err != nil {
		return err
	}
	if user.Password != vuser.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return errors.New("invaild password")
	}
	return nil
}
