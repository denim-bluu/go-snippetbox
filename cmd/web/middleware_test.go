package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecureHeaders(t *testing.T) {
	// Initialise a new response recorder and dummy HTTP request.
	recorder := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP handler to pass to the secureHeaders middleware,
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Pass the mock HTTP handler to the secureHeaders middleware and get the response by the handler.
	secureHeaders(next).ServeHTTP(recorder, r)
	rs := recorder.Result()

	// Validate that Content-Security-Policy header is set correctly.
	expectedValue := "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com"
	assert.Equal(t, rs.Header.Get("Content-Security-Policy"), expectedValue)

	// Validate that the middleware has correctly set the Referrer-Policy header
	expectedValue = "origin-when-cross-origin"
	assert.Equal(t, rs.Header.Get("Referrer-Policy"), expectedValue)

	// Validate that the middleware has correctly set the X-Content-Type-Options header
	expectedValue = "nosniff"
	assert.Equal(t, rs.Header.Get("X-Content-Type-Options"), expectedValue)

	// Validate that the middleware has correctly set the X-Frame-Options header
	expectedValue = "deny"
	assert.Equal(t, rs.Header.Get("X-Frame-Options"), expectedValue)

	// Validate that the middleware has correctly set the X-XSS-Protection header
	expectedValue = "0"
	assert.Equal(t, rs.Header.Get("X-XSS-Protection"), expectedValue)

	// Validate that the status code from the handler is 200.
	assert.Equal(t, rs.StatusCode, http.StatusOK)

	// Validate that the body of the response matches.
	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)

	assert.Equal(t, string(body), "OK")
}
