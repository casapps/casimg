package paths

import (
	"os"
	"path/filepath"
	"runtime"
)

// Paths holds all application directories
type Paths struct {
	ConfigDir string
	DataDir   string
	LogDir    string
	PIDFile   string
}

// GetDefault returns default paths based on OS
func GetDefault() *Paths {
	p := &Paths{}

	switch runtime.GOOS {
	case "windows":
		p.ConfigDir = filepath.Join(os.Getenv("ProgramData"), "casapps", "casimg")
		p.DataDir = filepath.Join(os.Getenv("ProgramData"), "casapps", "casimg", "data")
		p.LogDir = filepath.Join(os.Getenv("ProgramData"), "casapps", "casimg", "logs")
		p.PIDFile = filepath.Join(os.Getenv("ProgramData"), "casapps", "casimg", "casimg.pid")
	case "darwin":
		p.ConfigDir = "/usr/local/etc/casimg"
		p.DataDir = "/usr/local/var/lib/casimg"
		p.LogDir = "/usr/local/var/log/casimg"
		p.PIDFile = "/usr/local/var/run/casimg.pid"
	default:
		p.ConfigDir = "/etc/casimg"
		p.DataDir = "/var/lib/casimg"
		p.LogDir = "/var/log/casimg"
		p.PIDFile = "/var/run/casimg.pid"
	}

	return p
}

// EnsureDirs creates all required directories
func (p *Paths) EnsureDirs() error {
	dirs := []string{
		p.ConfigDir,
		p.DataDir,
		p.LogDir,
		filepath.Dir(p.PIDFile),
		filepath.Join(p.DataDir, "db"),
		filepath.Join(p.DataDir, "uploads"),
		filepath.Join(p.DataDir, "converted"),
		filepath.Join(p.DataDir, "backups"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	return nil
}
