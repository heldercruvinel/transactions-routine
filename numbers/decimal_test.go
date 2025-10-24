package numbers

import (
	"testing"
)

func TestFormatToTwoDecimalPlaces(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"rounds up", 1.235, 1.24},
		{"rounds down", 1.234, 1.23},
		{"no rounding needed", 2.50, 2.50},
		{"zero", 0.0, 0.0},
		{"negative value", -1.236, -1.24},
		{"large value", 123456.789, 123456.79},
		{"small value", 0.004, 0.00},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatToTwoDecimalPlaces(tt.input)
			if result != tt.expected {
				t.Errorf("FormatToTwoDecimalPlaces(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
