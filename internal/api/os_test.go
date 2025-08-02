// go
package api

import (
	"testing"
)

func TestGetOSCode(t *testing.T) {
	tests := []struct {
		name     string
		input    OS
		expected string
	}{
		{
			name: "ID and name with spaces and parentheses",
			input: OS{
				ID:   42,
				Name: "Ubuntu 20.04 (LTS)",
			},
			expected: "42-ubuntu-20.04-lts",
		},
		{
			name: "ID zero, name with spaces",
			input: OS{
				ID:   0,
				Name: "CentOS 7",
			},
			expected: "centos-7",
		},
		{
			name: "ID positive, name with no spaces",
			input: OS{
				ID:   7,
				Name: "Debian",
			},
			expected: "7-debian",
		},
		{
			name: "Name with uppercase and parentheses",
			input: OS{
				ID:   1,
				Name: "WINDOWS (Server)",
			},
			expected: "1-windows-server",
		},
		{
			name: "Name with multiple spaces",
			input: OS{
				ID:   5,
				Name: "Free BSD  13",
			},
			expected: "5-free-bsd--13",
		},
		{
			name: "Name with special characters",
			input: OS{
				ID:   9,
				Name: "Alpine@3.15 (Edge)",
			},
			expected: "9-alpine@3.15-edge",
		},
		{
			name: "ID zero, name with parentheses only",
			input: OS{
				ID:   0,
				Name: "(Test)",
			},
			expected: "test",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.input.GetCode()
			if got != tc.expected {
				t.Errorf("GetOSCode(%+v) = %q; want %q", tc.input, got, tc.expected)
			}
		})
	}
}
