package weather

import (
	"fmt"
	"net/http"
)

func HandleWeather(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Weather endpoint hit")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
