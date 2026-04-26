package main

import (
"flag"
"fmt"
"os"
"path/filepath"

"github.com/casapps/casimg/src/config"
"github.com/casapps/casimg/src/mode"
"github.com/casapps/casimg/src/paths"
"github.com/casapps/casimg/src/server"
)

var (
Version   = "dev"
CommitID  = "unknown"
BuildDate = "unknown"
)

func main() {
binaryName := filepath.Base(os.Args[0])

helpFlag := flag.Bool("help", false, "Show help")
hFlag := flag.Bool("h", false, "Show help")
versionFlag := flag.Bool("version", false, "Show version")
vFlag := flag.Bool("v", false, "Show version")
modeFlag := flag.String("mode", "", "Application mode")
configFlag := flag.String("config", "", "Configuration directory")
dataFlag := flag.String("data", "", "Data directory")
logFlag := flag.String("log", "", "Log directory")
pidFlag := flag.String("pid", "", "PID file path")
addressFlag := flag.String("address", "", "Listen address")
portFlag := flag.Int("port", 0, "Listen port")
debugFlag := flag.Bool("debug", false, "Enable debug mode")
statusFlag := flag.Bool("status", false, "Show status")

flag.Parse()

if *helpFlag || *hFlag {
fmt.Printf(`%s - Self-hosted file conversion service

Usage: %s [options]

Options:
  --help, -h              Show help
  --version, -v           Show version
  --mode                  Application mode (production|development)
  --config                Configuration directory
  --data                  Data directory
  --log                   Log directory
  --address               Listen address (default: 0.0.0.0)
  --port                  Listen port (default: 64580)
  --debug                 Enable debug mode
  --status                Show status

See AI.md for complete documentation.
`, binaryName, binaryName)
os.Exit(0)
}

if *versionFlag || *vFlag {
fmt.Printf("%s %s (commit: %s, built: %s)\n", binaryName, Version, CommitID, BuildDate)
os.Exit(0)
}

appMode := mode.Detect()
if *modeFlag != "" {
switch *modeFlag {
case "development", "dev":
appMode = mode.Development
case "production", "prod":
appMode = mode.Production
}
}

defaultPaths := paths.GetDefault()
if *configFlag != "" {
defaultPaths.ConfigDir = *configFlag
}
if *dataFlag != "" {
defaultPaths.DataDir = *dataFlag
}
if *logFlag != "" {
defaultPaths.LogDir = *logFlag
}
if *pidFlag != "" {
defaultPaths.PIDFile = *pidFlag
}

if err := defaultPaths.EnsureDirs(); err != nil {
fmt.Fprintf(os.Stderr, "Error: %v\n", err)
os.Exit(1)
}

cfg, err := config.Load(defaultPaths.ConfigDir)
if err != nil {
fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
os.Exit(1)
}
cfg.Mode = appMode
cfg.Paths = defaultPaths

if *addressFlag != "" {
cfg.Server.Address = *addressFlag
}
if *portFlag != 0 {
cfg.Server.Port = *portFlag
}
if *debugFlag {
cfg.Server.Debug = true
}

if *statusFlag {
fmt.Printf("Configuration:\n")
fmt.Printf("  Mode: %s\n", cfg.Mode)
fmt.Printf("  Config: %s\n", cfg.Paths.ConfigDir)
fmt.Printf("  Data: %s\n", cfg.Paths.DataDir)
fmt.Printf("  Server: %s:%d\n", cfg.Server.Address, cfg.Server.Port)
os.Exit(0)
}

srv := server.New(cfg, Version, CommitID, BuildDate)
if err := srv.Start(); err != nil {
fmt.Fprintf(os.Stderr, "Error: %v\n", err)
os.Exit(1)
}
}
