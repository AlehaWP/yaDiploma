package middlewares

import (
	"context"
	"net/http"

	"github.com/AlehaWP/yaDiploma.git/internal/models"
)

var userRepo models.UserRepo

func SetCookieUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("UserID")
		cv := ""
		if err == nil {
			cv = c.Value
		}
		if ok := userRepo.Find(cv); !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		c = &http.Cookie{
			Name:  "UserID",
			Value: cv,
		}
		http.SetCookie(w, c)

		ctx := context.WithValue(r.Context(), models.UKeyName, cv)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
