package config

import "strings"

// ParseBool parses various boolean representations
// Supports: true/false, yes/no, 1/0, on/off, enable/disable, enabled/disabled, y/n, t/f
func ParseBool(value string) bool {
	v := strings.ToLower(strings.TrimSpace(value))

	truthy := []string{
		"true", "yes", "1", "on", "enable", "enabled", "y", "t",
	}

	for _, t := range truthy {
		if v == t {
			return true
		}
	}

	return false
}
