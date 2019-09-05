package server

import (
	"net/http"
	"testing"
)

func TestStartFactory(t *testing.T) {
	type args struct {
		port string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StartFactory(tt.args.port)
		})
	}
}

func Test_handleRequests(t *testing.T) {
	type args struct {
		r http.ResponseWriter
		s *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleRequests(tt.args.r, tt.args.s)
		})
	}
}
