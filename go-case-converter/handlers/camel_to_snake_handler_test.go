package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCamelToSnakeHandler(t *testing.T) {
	reqBody, _ := json.Marshal(map[string]string{"input": "CamelCase"})
	req := httptest.NewRequest(http.MethodPost, "/convert/camel-to-snake", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	CamelToSnakeHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.Status)
	}

	var respBody map[string]string
	json.NewDecoder(resp.Body).Decode(&respBody)
	if respBody["result"] != "camel_case" {
		t.Errorf("Expected 'camel_case', got %v", respBody["result"])
	}
}

func TestCamelToSnakeHandler_InvalidInput(t *testing.T) {
	reqBody, _ := json.Marshal(map[string]string{"input": "Invalid Input!"})
	req := httptest.NewRequest(http.MethodPost, "/convert/camel-to-snake", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	CamelToSnakeHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status BAD REQUEST, got %v", resp.Status)
	}
}
