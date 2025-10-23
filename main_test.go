package main

import "testing"

func TestGenarateMessage(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "Alice",
			expected: "Hello Alice\n",
		},
		{
			name:     "",
			expected: "Hello World\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := genarateMessage(tt.name)
			if result != tt.expected {
				t.Errorf("genarateMessage(%q) = %q; want %q", tt.name, result, tt.expected)
			}
		})
	}
}
