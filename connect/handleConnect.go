package connect

import (
	"net/http"
)

func HandleConnect(w http.ResponseWriter, r *http.Request, svc Connect) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
