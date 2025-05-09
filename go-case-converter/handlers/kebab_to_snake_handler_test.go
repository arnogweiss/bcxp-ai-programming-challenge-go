package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestKebabToSnakeHandler(t *testing.T) {
	reqBody, _ := json.Marshal(map[string]string{"input": "kebab-case"})
	req := httptest.NewRequest(http.MethodPost, "/convert/kebab-to-snake", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	KebabToSnakeHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.Status)
	}

	var respBody map[string]string
	json.NewDecoder(resp.Body).Decode(&respBody)
	if respBody["result"] != "kebab_case" {
		t.Errorf("Expected 'kebab_case', got %v", respBody["result"])
	}
}

func TestKebabToSnakeHandler_InvalidInput(t *testing.T) {
	reqBody, _ := json.Marshal(map[string]string{"input": "Invalid Input!"})
	req := httptest.NewRequest(http.MethodPost, "/convert/kebab-to-snake", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	KebabToSnakeHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status BAD REQUEST, got %v", resp.Status)
	}
}
