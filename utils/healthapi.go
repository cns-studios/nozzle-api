package utils

import (
	"net/http"
)

func IsHealthy(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, map[string]interface{}{"status": "ok"})
}
