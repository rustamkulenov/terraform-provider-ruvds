package api

import (
	"testing"
)

func TestGetDatacenterCode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Russian with colon",
			input:    "Москва:RU-1",
			expected: "MOSKVA", // transliterate("Москва")
		},
		{
			name:     "Russian without colon",
			input:    "Санкт-Петербург",
			expected: "SANKT-PETERBURG", // transliterate("Санкт-Петербург")
		},
		{
			name:     "Latin with colon",
			input:    "LDN1:UK-1",
			expected: "LDN1",
		},
		{
			name:     "Latin without colon",
			input:    "Frankfurt",
			expected: "FRANKFURT",
		},
		{
			name:     "Empty name",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dc := DataCenter{Name: tt.input}
			got := dc.GetDatacenterCode()
			if got != tt.expected {
				t.Errorf("GetDatacenterCode() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestTransliterate(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Москва", "Moskva"},
		{"Санкт-Петербург", "Sankt-Peterburg"},
		{"London", "London"},
		{"", ""},
		{"123", "123"},
		{"тест123", "test123"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := transliterate(tt.input)
			if got != tt.expected {
				t.Errorf("transliterate(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestGetDatacenterCountryCode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Russian country in name",
			input:    "Москва, Россия",
			expected: "RU",
		},
		{
			name:     "German country in Russian",
			input:    "Франкфурт, Германия",
			expected: "DE",
		},
		{
			name:     "French country in Russian",
			input:    "Париж, Франция",
			expected: "",
		},
		{
			name:     "No country in name",
			input:    "Токио",
			expected: "",
		},
		{
			name:     "Only country, no city",
			input:    "Германия",
			expected: "DE",
		},
		{
			name:     "Kazakhstan country",
			input:    "Алматы, Казахстан",
			expected: "KZ",
		},
		{
			name:     "USA in Russian",
			input:    "Нью-Йорк, США",
			expected: "",
		},
		{
			name:     "Empty name",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dc := DataCenter{Name: tt.input}
			got := dc.GetDatacenterCountryCode()
			if got != tt.expected {
				t.Errorf("GetDatacenterCountryCode() = %q, want %q", got, tt.expected)
			}
		})
	}
}
