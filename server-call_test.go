package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockServer => Creates a mock server for unit testing
func MockServer(status int, encodeValue interface{}) *httptest.Server {

	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(encodeValue)
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

// TestGetData => Test Method for 201 status code
func TestGetData(t *testing.T) {

	t.Log("Testing GetMethod which returns mocked data")
	{
		serverResult := make(map[string]string)
		serverResult["ip"] = "dummy IP"

		server := MockServer(201, serverResult)
		//Here we are overwriting the url defined in main method
		// So now the call will go on this mock server url
		url = server.URL
		defer server.Close()
		result, err := GetData()
		if err != nil {
			t.Errorf("Expected nil, got ‘%s’", err)
		}
		if result.IP != "dummy IP" {
			t.Errorf("Expected ‘dummy IP’ request, got ‘%s’", result)
		}
	}
}

//TestGetDataError => Test Method for 400 Bad request
func TestGetDataError(t *testing.T) {

	t.Log("Testing GetMethod which returns 400 Bad Request")
	{
		serverResult := make(map[string]string)
		serverResult["ip"] = "dummy IP"

		server := MockServer(400, nil)
		url = server.URL

		defer server.Close()
		result, err := GetData()

		if result.IP != "" {
			t.Errorf("Expected '', got ‘%s’", result.IP)
		}
		if err == nil {
			t.Errorf("Expected error, got ‘%s’", err)
		}

		if err.Error() != "400 Bad Request" {
			t.Errorf("Expected ‘400 Bad Request’ request, got ‘%s’", err)
		}

		if err.Error() == "400 Bad Request" {
			t.Logf("Expected ‘400 Bad Request’ request, got ‘%s’", err)
		}
	}
}
