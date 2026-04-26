package server

import (
"context"
"fmt"
"net/http"
"os"
"os/signal"
"syscall"
"time"

"github.com/casapps/casimg/src/config"
"github.com/casapps/casimg/src/server/handler"
"github.com/casapps/casimg/src/swagger"
)

type Server struct {
config    *config.Config
version   string
commitID  string
buildDate string
startTime time.Time
httpSrv   *http.Server
}

func New(cfg *config.Config, version, commitID, buildDate string) *Server {
return &Server{
config:    cfg,
version:   version,
commitID:  commitID,
buildDate: buildDate,
startTime: time.Now(),
}
}

func (s *Server) Start() error {
mux := http.NewServeMux()

healthHandler := handler.NewHealthHandler(s.version, s.config.Mode.String())
convertHandler := handler.NewConvertHandler(s.config.Paths.DataDir)

mux.HandleFunc("/healthz", healthHandler.HandleHealth)
mux.HandleFunc("/api/v1/healthz", healthHandler.HandleHealthJSON)

mux.HandleFunc("/api/v1/convert", convertHandler.HandleUpload)
mux.HandleFunc("/api/v1/convert/", func(w http.ResponseWriter, r *http.Request) {
if r.URL.Path[len("/api/v1/convert/"):] == "" {
convertHandler.HandleUpload(w, r)
return
}
if len(r.URL.Path) > len("/api/v1/convert/") && r.URL.Path[len(r.URL.Path)-9:] == "/download" {
convertHandler.HandleDownload(w, r)
} else {
convertHandler.HandleStatus(w, r)
}
})
mux.HandleFunc("/api/v1/formats", convertHandler.HandleFormats)

// Swagger/OpenAPI routes - See AI.md PART 14
mux.HandleFunc("/openapi", swagger.ServeSwaggerUI)
mux.HandleFunc("/openapi.json", swagger.ServeOpenAPIJSON)

mux.HandleFunc("/", s.handleHome)

addr := fmt.Sprintf("%s:%d", s.config.Server.Address, s.config.Server.Port)
s.httpSrv = &http.Server{
Addr:    addr,
Handler: mux,
}

fmt.Printf("casimg %s starting...\n", s.version)
fmt.Printf("Mode: %s\n", s.config.Mode)
fmt.Printf("Listening on http://%s\n", addr)

errChan := make(chan error, 1)
go func() {
if err := s.httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
errChan <- err
}
}()

sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

select {
case err := <-errChan:
return err
case <-sigChan:
fmt.Println("\nShutting down gracefully...")
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
return s.httpSrv.Shutdown(ctx)
}
}

func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {
if r.URL.Path != "/" {
http.NotFound(w, r)
return
}

html := `<!DOCTYPEhtml>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>casimg - File Conversion Service</title>
    <style>
        body {
            font-family: system-ui, -apple-system, sans-serif;
            max-width: 900px;
            margin: 50px auto;
            padding: 20px;
            line-height: 1.6;
        }
        h1 { color: #333; }
        .status { color: #28a745; font-weight: bold; }
        .card {
            background: #f8f9fa;
            border: 1px solid #dee2e6;
            border-radius: 8px;
            padding: 20px;
            margin: 20px 0;
        }
        .endpoints { list-style: none; padding: 0; }
        .endpoints li { padding: 8px 0; }
        .endpoint { font-family: monospace; background: #e9ecef; padding: 4px 8px; border-radius: 4px; }
        .formats { display: flex; flex-wrap: wrap; gap: 10px; }
        .format-badge {
            background: #007bff;
            color: white;
            padding: 4px 12px;
            border-radius: 12px;
            font-size: 14px;
        }
    </style>
</head>
<body>
    <h1>🔄 casimg - File Conversion Service</h1>
    <p class="status">✓ Server is running</p>
    <p><strong>Version:</strong> ` + s.version + ` | <strong>Mode:</strong> ` + s.config.Mode.String() + `</p>

    <div class="card">
        <h2>Quick Start</h2>
        <p>Upload and convert files using the API:</p>
        <pre><code>curl -X POST http://localhost:` + fmt.Sprintf("%d", s.config.Server.Port) + `/api/v1/convert \
  -F "file=@document.pdf" \
  -F "output_format=docx"</code></pre>
    </div>

    <div class="card">
        <h2>API Endpoints</h2>
        <ul class="endpoints">
            <li><span class="endpoint">GET /healthz</span> - Health check (HTML)</li>
            <li><span class="endpoint">GET /api/v1/healthz</span> - Health check (JSON)</li>
            <li><span class="endpoint">GET /api/v1/formats</span> - List supported formats</li>
            <li><span class="endpoint">POST /api/v1/convert</span> - Upload and convert file</li>
            <li><span class="endpoint">GET /api/v1/convert/{id}</span> - Check conversion status</li>
            <li><span class="endpoint">GET /api/v1/convert/{id}/download</span> - Download result</li>
        </ul>
    </div>

    <div class="card">
        <h2>Supported Formats</h2>
        <div class="formats">
            <span class="format-badge">PDF</span>
            <span class="format-badge">DOCX</span>
            <span class="format-badge">JPG</span>
            <span class="format-badge">PNG</span>
            <span class="format-badge">MP3</span>
            <span class="format-badge">MP4</span>
            <span class="format-badge">+more</span>
        </div>
        <p><a href="/api/v1/formats">View all formats →</a></p>
    </div>

    <div class="card">
        <h2>Documentation</h2>
        <ul>
            <li><a href="https://github.com/casapps/casimg">GitHub Repository</a></li>
            <li><a href="https://casapps-casimg.readthedocs.io">Documentation</a></li>
            <li><a href="/healthz">Server Health</a></li>
        </ul>
    </div>
</body>
</html>`

w.Header().Set("Content-Type", "text/html; charset=utf-8")
w.Write([]byte(html))
}
