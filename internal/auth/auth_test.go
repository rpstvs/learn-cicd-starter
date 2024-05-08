package auth

import (
	"net/http"
	"testing"
)

// TestGetAPIKey checks various scenarios for the GetAPIKey function.
func TestGetAPIKey(t *testing.T) {
	// Define test cases as a slice of struct.
	tests := []struct {
		name           string      // Descriptive name of the test case
		headers        http.Header // Mock headers to simulate requests
		expectedAPIKey string      // Expected API key result
		expectError    bool        // Whether an error is expected
	}{
		{
			name:           "No Auth Header",
			headers:        http.Header{}, // Empty header to simulate missing auth
			expectedAPIKey: "",
			expectError:    true, // Expect an error in this case
		},
		{
			name: "Malformed Auth Header",
			headers: http.Header{
				"Authorization": []string{"BadHeader"}, // Incorrect format
			},
			expectedAPIKey: "",
			expectError:    true,
		},
		{
			name: "Valid Auth Header",
			headers: http.Header{
				"Authorization": []string{"ApiKey 12345"}, // Correct format
			},
			expectedAPIKey: "12345", // The expected API key
			expectError:    false,   // No error should occur here
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tc.headers)
			if tc.expectError && err == nil {
				t.Errorf("Expected an error but got none")
			}
			if !tc.expectError && err != nil {
				t.Errorf("Did not expect an error but got one: %v", err)
			}
			if apiKey != tc.expectedAPIKey {
				t.Errorf("Expected API Key '%s', got '%s'", tc.expectedAPIKey, apiKey)
			}
		})
	}
}
