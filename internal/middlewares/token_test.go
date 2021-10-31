package middlewares

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestCheckAuthorization(t *testing.T) {
// 	type args struct {
// 		next http.Handler
// 	}
//     handler :=
// 	r := httptest.NewRequest("GET", "/", strings.NewReader(""))
// 	r.Header.Add("Authorization", "")
// 	w := httptest.NewRecorder()

// 	tests := []struct {
// 		name string
// 		args args
// 		want http.Handler
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := (tt.args.next); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("CheckAuthorization() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

type testHandler struct{}

func (t *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}

func TestCheckAuthorization(t *testing.T) {

	testValue := map[string]struct {
		token     string
		resStatus int
	}{
		"test1": {
			token:     "",
			resStatus: http.StatusUnauthorized,
		},
		"test2": {
			token:     "Bearer asdflkajfhkajdf",
			resStatus: http.StatusAccepted,
		},
		"test3": {
			token:     "Bearer ",
			resStatus: http.StatusUnauthorized,
		},
	}

	r := httptest.NewRequest("GET", "/", strings.NewReader(""))

	tHandler := new(testHandler)
	handler := CheckAuthorization(tHandler)
	var w *httptest.ResponseRecorder

	for _, v := range testValue {
		w = httptest.NewRecorder()
		r.Header.Set("Authorization", v.token)
		handler.ServeHTTP(w, r)
		res := w.Result()

		assert.Equal(t, v.resStatus, res.StatusCode, "Не верный код ответа GET")
	}

}
