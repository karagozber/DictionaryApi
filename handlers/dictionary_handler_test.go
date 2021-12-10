package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetValue(t *testing.T) {
	req, err := http.NewRequest("GET", "api/get", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("Key", "berkay")
	req.URL.RawQuery = q.Encode()
	dictionaryHandler := NewDictionaryHandlers()
	dictionaryHandler.Data["berkay"] = "karagöz"

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(dictionaryHandler.GetValue)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `{"key":"berkay","value":"karagöz"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestSetValue(t *testing.T) {
	req, err := http.NewRequest("POST", "api/set", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("Key", "berkay")
	q.Add("Value", "karagöz")
	req.URL.RawQuery = q.Encode()
	dictionaryHandler := NewDictionaryHandlers()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(dictionaryHandler.SetValue)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if dictionaryHandler.Data["berkay"] != "karagöz" {
		t.Errorf("handler returned unexpected body: got %v want %v", dictionaryHandler.Data["berkay"], "karagöz")
	}
}
