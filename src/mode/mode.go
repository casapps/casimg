package mode

import "os"

type Mode string

const (
	Production  Mode = "production"
	Development Mode = "development"
)

// Detect returns the application mode from environment or default
func Detect() Mode {
	mode := os.Getenv("CASIMG_MODE")
	if mode == "" {
		mode = os.Getenv("MODE")
	}

	switch mode {
	case "development", "dev":
		return Development
	default:
		return Production
	}
}

// IsDevelopment returns true if mode is development
func (m Mode) IsDevelopment() bool {
	return m == Development
}

// IsProduction returns true if mode is production
func (m Mode) IsProduction() bool {
	return m == Production
}

// String returns the string representation
func (m Mode) String() string {
	return string(m)
}
