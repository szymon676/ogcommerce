package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/szymon676/ogcommerce/store"
	"github.com/szymon676/ogcommerce/types"
)

const secret = "secret"

type AuthHandler struct {
	store store.UsersStorager
}

func NewAuthHandler(astore store.UsersStorager) *AuthHandler {
	return &AuthHandler{
		store: astore,
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

	token, err := h.CreateJWT(&reqUser, r.Context())

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

func (ah AuthHandler) CreateJWT(user *types.ReqUser, ctx context.Context) (string, error) {
	account, _ := ah.store.GetUserByEmail(ctx, user.Email)

	claims := &jwt.MapClaims{
		"expiresAt": time.Now().Add(time.Hour * 24).Unix(),
		"accountID": account.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func (ah AuthHandler) ParseJWT(r *http.Request) (*jwt.Token, error) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		return nil, fmt.Errorf("%c", err)
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid JWT: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid JWT: token is not valid")
	}

	return token, nil
}

func CreateCookie(w http.ResponseWriter, token string) error {
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  expiration,
		Path:     "/",
		MaxAge:   604800,
		SameSite: http.SameSiteNoneMode,
	}

	http.SetCookie(w, &cookie)
	return nil
}
