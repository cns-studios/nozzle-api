package utils

import (
	"net/http"
)

func IsHealthy(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]interface{}{"health": "ok"})
}
