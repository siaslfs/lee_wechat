package services

import "testing"

func TestGetAccessToken(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "获取accessToken"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetAccessToken()
		})
	}
}
