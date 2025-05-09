package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSnakeToKebabHandler(t *testing.T) {
	reqBody, _ := json.Marshal(map[string]string{"input": "snake_case"})
	req := httptest.NewRequest(http.MethodPost, "/convert/snake-to-kebab", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	SnakeToKebabHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.Status)
	}

	var respBody map[string]string
	json.NewDecoder(resp.Body).Decode(&respBody)
	if respBody["result"] != "snake-case" {
		t.Errorf("Expected 'snake-case', got %v", respBody["result"])
	}
}

func TestSnakeToKebabHandler_InvalidInput(t *testing.T) {
	reqBody, _ := json.Marshal(map[string]string{"input": "Invalid Input!"})
	req := httptest.NewRequest(http.MethodPost, "/convert/snake-to-kebab", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	SnakeToKebabHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status BAD REQUEST, got %v", resp.Status)
	}
}
