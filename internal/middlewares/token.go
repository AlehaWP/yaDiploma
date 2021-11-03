package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/AlehaWP/yaDiploma.git/internal/models"
)

var userRepo models.UsersRepo

func CheckAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenName := "bearer "
		t := r.Header.Get("Authorization")
		key := ""
		if strings.HasPrefix(strings.ToLower(t), tokenName) {
			key = t[len(tokenName):]
		}

		if len(key) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.Header().Add("Authorization", t)
		ctx := context.WithValue(r.Context(), models.UKeyName, key)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
