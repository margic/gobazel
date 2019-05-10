package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"

	"github.com/stretchr/testify/assert"
)

func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createHandlers().HealthHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, "OK", rr.Body.String(), "Body not ok")
}

func createHandlers() *Handlers {
	logger, _ := zap.NewDevelopment()
	h := &Handlers{
		logger: logger,
	}
	return h
}
