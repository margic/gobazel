package main

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

// Handlers is a custom type for http handlers that provides handlers as pointer receivers
// with access to logger
type Handlers struct {
	logger *zap.Logger
}

// HealthHandler returns ok when called
func (h *Handlers) HealthHandler(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("Health Handler", zap.String("path", r.URL.String()))
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
