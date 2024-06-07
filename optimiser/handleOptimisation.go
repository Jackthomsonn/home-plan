package optimiser

import (
	"net/http"
)

func HandleOptimisation(w http.ResponseWriter, r *http.Request, svc Optimiser) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
