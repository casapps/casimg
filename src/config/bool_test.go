package config

import "testing"

func TestParseBool(t *testing.T) {
tests := []struct {
input    string
expected bool
}{
{"true", true},
{"True", true},
{"TRUE", true},
{"yes", true},
{"Yes", true},
{"1", true},
{"on", true},
{"enable", true},
{"enabled", true},
{"false", false},
{"no", false},
{"0", false},
{"off", false},
{"disabled", false},
{"", false},
{"invalid", false},
}

for _, tt := range tests {
result := ParseBool(tt.input)
if result != tt.expected {
t.Errorf("ParseBool(%q) = %v; want %v", tt.input, result, tt.expected)
}
}
}
