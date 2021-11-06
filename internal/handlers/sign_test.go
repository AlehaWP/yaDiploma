package handlers

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/AlehaWP/yaDiploma.git/internal/models"
)

func TestHandlerLogin(t *testing.T) {
	// type args struct {
	// 	w http.ResponseWriter
	// 	r *http.Request
	// }
	// tests := []struct {
	// 	name string
	// 	args args
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		HandlerLogin(tt.args.w, tt.args.r)
	// 	})
	// }
}

func TestHandlerRegistration(t *testing.T) {
	type args struct {
		ur models.UsersRepo
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandlerRegistration(tt.args.ur); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandlerRegistration() = %v, want %v", got, tt.want)
			}
		})
	}
}
