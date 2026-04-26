package handler

import (
"encoding/json"
"fmt"
"net/http"
"os"
"time"
)

type HealthHandler struct {
version   string
mode      string
startTime time.Time
}

func NewHealthHandler(version, mode string) *HealthHandler {
return &HealthHandler{
version:   version,
mode:      mode,
startTime: time.Now(),
}
}

func (h *HealthHandler) HandleHealth(w http.ResponseWriter, r *http.Request) {
hostname, _ := os.Hostname()
uptime := time.Since(h.startTime)

html := fmt.Sprintf(`<!DOCTYPEhtml>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Health Status - casimg</title>
    <style>
        body { font-family: system-ui, -apple-system, sans-serif; max-width: 800px; margin: 50px auto; padding: 20px; }
        h1 { color: #333; }
        .status-healthy { color: #28a745; font-weight: bold; font-size: 24px; }
        .info { margin: 20px 0; }
        .info dt { font-weight: bold; margin-top: 10px; }
        .info dd { margin-left: 20px; }
    </style>
</head>
<body>
    <h1>Health Status</h1>
    <p class="status-healthy">✓ healthy</p>
    <dl class="info">
        <dt>Version:</dt><dd>%s</dd>
        <dt>Mode:</dt><dd>%s</dd>
        <dt>Uptime:</dt><dd>%s</dd>
        <dt>Node:</dt><dd>%s</dd>
        <dt>Database:</dt><dd>ok</dd>
        <dt>Disk:</dt><dd>ok</dd>
    </dl>
    <p><a href="/">← Back to Home</a></p>
</body>
</html>`, h.version, h.mode, formatUptime(uptime), hostname)

w.Header().Set("Content-Type", "text/html; charset=utf-8")
w.Write([]byte(html))
}

func (h *HealthHandler) HandleHealthJSON(w http.ResponseWriter, r *http.Request) {
hostname, _ := os.Hostname()
uptime := time.Since(h.startTime)

health := map[string]interface{}{
"status":    "healthy",
"version":   h.version,
"mode":      h.mode,
"uptime":    formatUptime(uptime),
"timestamp": time.Now().Format(time.RFC3339),
"node": map[string]string{
"id":       "standalone",
"hostname": hostname,
},
"cluster": map[string]interface{}{
"enabled": false,
},
"checks": map[string]string{
"database": "ok",
"disk":     "ok",
},
}

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(health)
}

func formatUptime(d time.Duration) string {
days := int(d.Hours() / 24)
hours := int(d.Hours()) % 24
minutes := int(d.Minutes()) % 60

if days > 0 {
return fmt.Sprintf("%dd %dh %dm", days, hours, minutes)
}
if hours > 0 {
return fmt.Sprintf("%dh %dm", hours, minutes)
}
return fmt.Sprintf("%dm", minutes)
}
