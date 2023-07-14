package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Authenticate(t *testing.T) {
	
	postBody := map[string]interface{}{
		"email":    "me@test.com",
		"password": "password",
	}

	body, _ := json.Marshal(postBody)

	req, err := http.NewRequest("POST", "/authentication", bytes.NewReader(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.Authentication.Authenticate)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
