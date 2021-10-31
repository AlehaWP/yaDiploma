package middlewares

import (
	"context"
	"net/http"
	"regexp"

	"github.com/AlehaWP/yaDiploma.git/internal/models"
)

var userRepo models.UserRepo

func CheckAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := r.Header.Get("Authorization")

		re := regexp.MustCompile(`Bearer\s(.+)`)
		st := re.FindStringSubmatch(t) //FindAString(t)

		if len(st) < 2 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.Header().Add("Authorization", t)
		ctx := context.WithValue(r.Context(), models.UKeyName, st[1])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// cv := ""
// if err == nil {
// 	cv = c.Value
// }
// if ok := userRepo.Find(cv); !ok {
// 	w.WriteHeader(http.StatusUnauthorized)
// 	return
// }
// bearer := "Bearer " + encription.EncriptStr(cv)
