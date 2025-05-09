package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSnakeToCamelHandler(t *testing.T) {
	reqBody, _ := json.Marshal(map[string]string{"input": "snake_case"})
	req := httptest.NewRequest(http.MethodPost, "/convert/snake-to-camel", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	SnakeToCamelHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.Status)
	}

	var respBody map[string]string
	json.NewDecoder(resp.Body).Decode(&respBody)
	if respBody["result"] != "snakeCase" {
		t.Errorf("Expected 'snakeCase', got %v", respBody["result"])
	}
}

func TestSnakeToCamelHandler_InvalidInput(t *testing.T) {
	reqBody, _ := json.Marshal(map[string]string{"input": "Invalid Input!"})
	req := httptest.NewRequest(http.MethodPost, "/convert/snake-to-camel", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	SnakeToCamelHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status BAD REQUEST, got %v", resp.Status)
	}
}
