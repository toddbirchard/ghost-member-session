package main

import (
	"net/http"
	"testing"
	"time"
)

func TestGetSessionInfo(t *testing.T) {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Make request
	resp, err := client.Get("https://hackersandslackers.app/members/api/member/")
	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 200 (ok)
	if resp.StatusCode != http.StatusOK {
		t.Log("Status should be ok, got %i", resp.StatusCode)
	}

}
