package tests

import (
	"alarmservice/handlers"
	"alarmservice/storage"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouterAndStore() (*gin.Engine, *storage.MemoryStore) {
	store := storage.NewMemoryStore()
	router := gin.Default()
	handlers.RegisterRoutes(router, store)
	return router, store
}

func TestCreateAlarm(t *testing.T) {
	r, _ := setupRouterAndStore()
	alarm := map[string]interface{}{
		"name":      "High CPU",
		"condition": "cpu > 95",
	}
	body, _ := json.Marshal(alarm)
	req, _ := http.NewRequest("POST", "/alarm", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, 201, resp.Code)
}

func TestGetAlarms(t *testing.T) {
	r, _ := setupRouterAndStore()
	req, _ := http.NewRequest("GET", "/alarm", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

func TestCRUDOperations(t *testing.T) {
	r, _ := setupRouterAndStore()

	// Create Alarm
	alarm := map[string]interface{}{
		"name":      "High Memory",
		"condition": "mem > 90",
	}
	body, _ := json.Marshal(alarm)
	req, _ := http.NewRequest("POST", "/alarm", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, 201, resp.Code)

	var created map[string]interface{}
	_ = json.Unmarshal(resp.Body.Bytes(), &created)
	id := created["id"].(string)

	// Get by ID
	req, _ = http.NewRequest("GET", "/alarm/"+id, nil)
	resp = httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)

	// Update
	updated := map[string]interface{}{
		"name":      "Updated Alarm",
		"condition": "mem > 95",
		"state":     "ACKED",
	}
	updateBody, _ := json.Marshal(updated)
	req, _ = http.NewRequest("PUT", "/alarm/"+id, bytes.NewBuffer(updateBody))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)

	// Delete
	req, _ = http.NewRequest("DELETE", "/alarm/"+id, nil)
	resp = httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, 204, resp.Code)

	// Ensure it's deleted
	req, _ = http.NewRequest("GET", "/alarm/"+id, nil)
	resp = httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, 404, resp.Code)
}
