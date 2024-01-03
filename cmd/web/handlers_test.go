package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	mockApp := newTestApplication(t)

	ts := newTestServer(t, mockApp.newRouter())
	defer ts.Close()

	code, _, body := ts.getMethod(t, "/ping")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}

func TestSnippetView(t *testing.T) {
	mockApp := newTestApplication(t)

	ts := newTestServer(t, mockApp.newRouter())
	defer ts.Close()

	tests := []struct {
		name         string
		url          string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Valid instance",
			url:          "/snippet/view/1",
			expectedCode: http.StatusOK,
			expectedBody: "Mock Content",
		},
		{
			name:         "Non-existent instance",
			url:          "/snippet/view/2",
			expectedCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			code, _, body := ts.getMethod(t, test.url)
			assert.Equal(t, code, test.expectedCode)
			if test.expectedBody != "" {
				assert.Contains(t, body, test.expectedBody)
			}
		})
	}
}
