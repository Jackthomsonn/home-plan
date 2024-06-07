package optimiser

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jackthomsonn/home-plan/connections"
)

type Event string

const (
	Raise Event = "raise"
	Lower Event = "lower"
	Skip  Event = "skip"
)

type Plan struct {
	Event Event  `json:"event"`
	Value uint32 `json:"value"`
}

func HandleOptimisation(w http.ResponseWriter, r *http.Request, svc Optimiser) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading request body", http.StatusBadRequest)
		return
	}

	var reqBody OptimiserRequest
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		http.Error(w, "error unmarshalling request body", http.StatusBadRequest)
		return
	}

	if reqBody.Type == "nest" {
		svc.connection = connections.NewNestConnection()
	} else {
		http.Error(w, "invalid connection type", http.StatusBadRequest)
	}

	plan := CreatePlan(svc)

	svc.connection.ControlTemperature(plan.Value)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(plan)
}

func CreatePlan(svc Optimiser) Plan {
	svc.weatherSvc.Collect()
	svc.gridSvc.Collect()
	// Do some calculations
	// ...
	result := Skip

	if result == Skip {
		return Plan{
			Event: Skip,
		}
	}

	return Plan{
		Event: Raise,
		Value: 10,
	}
}
