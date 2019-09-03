package main

import (
	"io"
	"net/http/httptest"
	"testing"
)

// Testing Handling Requests
func HandlingRequests(t *testing.T) {
	i := io.LimitReader(nil, 0)
	assert := HandleRequests(httptest.NewRecorder(), httptest.NewRequest("GET", "/api/pages", i))

	if assert != 1 {
		t.Errorf("Abs(-1) = %d; want 1", assert)
	}
}
