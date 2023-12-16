package main

import (
	"encoding/json"
	"net/http"
)

// ... (previous code)

// /process-single endpoint handler
func ProcessSingle(w http.ResponseWriter, r *http.Request) {
	var payload RequestPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	timeTaken := sequentialSorting(payload.ToSort)

	response := struct {
		SortedArrays [][]int       `json:"sorted_arrays"`
		TimeNS       int64         `json:"time_ns"`
	}{
		SortedArrays: payload.ToSort,
		TimeNS:       timeTaken.Nanoseconds(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// /process-concurrent endpoint handler
func ProcessConcurrent(w http.ResponseWriter, r *http.Request) {
	var payload RequestPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	timeTaken := concurrentSorting(payload.ToSort)

	response := struct {
		SortedArrays [][]int       `json:"sorted_arrays"`
		TimeNS       int64         `json:"time_ns"`
	}{
		SortedArrays: payload.ToSort,
		TimeNS:       timeTaken.Nanoseconds(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
