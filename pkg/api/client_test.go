package api

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	if client == nil {
		t.Error("NewClient should return a valid client")
	}
	if client.baseURL == "" {
		t.Error("Client should have a base URL")
	}
}

func TestSearch(t *testing.T) {
	client := NewClient()
	result, err := client.Search("doctor", "tv", "")
	if err != nil {
		t.Errorf("Search should not return error: %v", err)
	}
	if result == nil {
		t.Error("Search should return a result")
	}
	if result.Query != "doctor" {
		t.Errorf("Expected query 'doctor', got '%s'", result.Query)
	}
}

func TestGetProgrammeByPID(t *testing.T) {
	client := NewClient()
	prog, err := client.GetProgrammeByPID("test123")
	if err != nil {
		t.Errorf("GetProgrammeByPID should not return error: %v", err)
	}
	if prog == nil {
		t.Error("GetProgrammeByPID should return a programme")
	}
	if prog.PID != "test123" {
		t.Errorf("Expected PID 'test123', got '%s'", prog.PID)
	}
}
