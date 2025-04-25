package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

const timeFormat = "2006-01-02 15:04:05 -0700 MST"

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	tz := r.URL.Query().Get("tz")
	w.Header().Set("Content-Type", "application/json")

	if tz == "" {
		now := time.Now().UTC()
		json.NewEncoder(w).Encode(map[string]string{
			"current_time": now.Format(timeFormat),
		})
		return
	}

	zones := strings.Split(tz, ",")
	result := map[string]string{}

	for _, zone := range zones {
		zoneName := strings.TrimSpace(zone)
		loc, err := time.LoadLocation(zoneName)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid timezone"})
			return
		}
		now := time.Now().In(loc)
		result[zoneName] = now.Format(timeFormat)
	}

	json.NewEncoder(w).Encode(result)
}
