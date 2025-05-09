package handlers

import (
	"encoding/json"
	"net/http"

	"go-case-converter/converters"
	"go-case-converter/validators"
)

type ConversionRequest struct {
	Input string `json:"input"`
}

type ConversionResponse struct {
	Result string `json:"result"`
}

// CamelToSnakeHandler handles the conversion of camel case to snake case.
func CamelToSnakeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req ConversionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input
	if err := validators.ValidateSingleWord(req.Input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := converters.ConvertCamelToSnake(req.Input)
	resp := ConversionResponse{Result: result}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
