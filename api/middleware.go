package api

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func (jwts JwtService) AuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := jwts.ParseJWT(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		accountID, ok := claims["accountID"].(string)
		if !ok {
			http.Error(w, "Invalid account ID in token", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "accountID", accountID)

		handler(w, r.WithContext(ctx))
	}
}
