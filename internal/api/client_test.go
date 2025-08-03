package api

import (
	"log"
	"os"
	"testing"
)

var client *Client

// Tests if the client is configured with the default endpoint.
func TestIsDefaultEndpoint(t *testing.T) {
	client := Client{
		endpoint: DefaultEndpoint,
	}

	if client.endpoint != DefaultEndpoint {
		t.Errorf("Expected endpoint to be %s, got %s", DefaultEndpoint, client.endpoint)
	}
}

// Test getting data centers response structure.
func TestGetDataCenters(t *testing.T) {
	if client == nil {
		t.Skip("Skipping test which requires a live API connection")
	}

	response, err := client.GetDataCenters()
	if err != nil {
		t.Fatalf("Failed to get data centers: %v", err)
	}
	if len(response.DataCenters) == 17 {
		t.Error("Expected 17 data center, got none")
	}
	for _, dc := range response.DataCenters {
		log.Printf("%v", dc)
		if dc.ID == 0 || dc.Name == "" {
			t.Error("Data center ID or Name is empty")
		}
		if len(dc.VpsTariffs) == 0 && len(dc.DriveTariffs) == 0 {
			t.Error("Data center should have at least one tariff")
		}
	}
}

// Test getting OS response structure.
func TestGetOS(t *testing.T) {
	if client == nil {
		t.Skip("Skipping test which requires a live API connection")
	}

	response, err := client.GetOSList()
	if err != nil {
		t.Fatalf("Failed to get OS: %v", err)
	}
	if len(response.Items) < 3 {
		t.Error("Expected at least one OS, got none")
	}
	for _, os := range response.Items {
		log.Printf("%v", os)
		if os.ID == 0 || os.Name == "" {
			t.Error("OS ID or Name is empty")
		}
	}
}

func init() {
	// Environment variable can be set in VSCode settings:
	// in .vscode/settings.json add:
	// 	"go.testEnvVars": {
	//  	   "RUVDS_API_TOKEN": "<TOKEN>"
	// 	}
	token := os.Getenv("RUVDS_API_TOKEN")
	if token == "" {
		// client will be not initialized
		log.Println("RUVDS_API_TOKEN environment variable is not set, tests will not run")
		return
	}
	client = NewClient(token, "")
	if client.endpoint != DefaultEndpoint {
		panic("API client not initialized with default endpoint")
	}
}
