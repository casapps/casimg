package swagger

import (
	"encoding/json"
	"net/http"
	"strings"
)

// OpenAPISpec generates the OpenAPI 3.0 specification
// See AI.md PART 14: API STRUCTURE - Swagger & GraphQL Sync
func OpenAPISpec() map[string]interface{} {
	return map[string]interface{}{
		"openapi": "3.0.0",
		"info": map[string]interface{}{
			"title":       "casimg API",
			"description": "Self-hosted file conversion service - A free, open-source alternative to CloudConvert and Convertio",
			"version":     "1.0.0",
			"license": map[string]string{
				"name": "MIT",
				"url":  "https://github.com/casapps/casimg/blob/main/LICENSE.md",
			},
		},
		"servers": []map[string]interface{}{
			{
				"url":         "{protocol}://{host}",
				"description": "Current server",
				"variables": map[string]interface{}{
					"protocol": map[string]interface{}{
						"default": "http",
						"enum":    []string{"http", "https"},
					},
					"host": map[string]interface{}{
						"default": "localhost:64580",
					},
				},
			},
		},
		"paths": map[string]interface{}{
			"/api/v1/healthz": map[string]interface{}{
				"get": map[string]interface{}{
					"tags":        []string{"Health"},
					"summary":     "Health check",
					"description": "Returns server health status",
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Server is healthy",
							"content": map[string]interface{}{
								"application/json": map[string]interface{}{
									"schema": map[string]interface{}{
										"type": "object",
										"properties": map[string]interface{}{
											"status": map[string]string{
												"type":    "string",
												"example": "ok",
											},
											"version": map[string]string{
												"type":    "string",
												"example": "1.0.0",
											},
											"mode": map[string]string{
												"type":    "string",
												"example": "production",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"/api/v1/formats": map[string]interface{}{
				"get": map[string]interface{}{
					"tags":        []string{"Conversion"},
					"summary":     "List supported formats",
					"description": "Returns all supported file formats and their conversion capabilities",
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "List of supported formats",
							"content": map[string]interface{}{
								"application/json": map[string]interface{}{
									"schema": map[string]interface{}{
										"type": "object",
										"properties": map[string]interface{}{
											"formats": map[string]interface{}{
												"type": "array",
												"items": map[string]interface{}{
													"$ref": "#/components/schemas/FormatInfo",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"/api/v1/convert": map[string]interface{}{
				"post": map[string]interface{}{
					"tags":        []string{"Conversion"},
					"summary":     "Upload and convert file",
					"description": "Upload a file for format conversion",
					"requestBody": map[string]interface{}{
						"required": true,
						"content": map[string]interface{}{
							"multipart/form-data": map[string]interface{}{
								"schema": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"file": map[string]string{
											"type":        "string",
											"format":      "binary",
											"description": "File to convert",
										},
										"output_format": map[string]string{
											"type":        "string",
											"description": "Target format (e.g., pdf, jpg, mp4)",
										},
									},
									"required": []string{"file", "output_format"},
								},
							},
						},
					},
					"responses": map[string]interface{}{
						"201": map[string]interface{}{
							"description": "Conversion job created",
							"content": map[string]interface{}{
								"application/json": map[string]interface{}{
									"schema": map[string]interface{}{
										"$ref": "#/components/schemas/ConversionJob",
									},
								},
							},
						},
						"400": map[string]interface{}{
							"description": "Invalid request",
							"content": map[string]interface{}{
								"application/json": map[string]interface{}{
									"schema": map[string]interface{}{
										"$ref": "#/components/schemas/Error",
									},
								},
							},
						},
					},
				},
			},
			"/api/v1/convert/{job_id}": map[string]interface{}{
				"get": map[string]interface{}{
					"tags":        []string{"Conversion"},
					"summary":     "Get conversion job status",
					"description": "Check the status and progress of a conversion job",
					"parameters": []map[string]interface{}{
						{
							"name":        "job_id",
							"in":          "path",
							"required":    true,
							"description": "Conversion job ID",
							"schema": map[string]string{
								"type": "string",
							},
						},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Job status",
							"content": map[string]interface{}{
								"application/json": map[string]interface{}{
									"schema": map[string]interface{}{
										"$ref": "#/components/schemas/ConversionJob",
									},
								},
							},
						},
					},
				},
			},
			"/api/v1/convert/{job_id}/download": map[string]interface{}{
				"get": map[string]interface{}{
					"tags":        []string{"Conversion"},
					"summary":     "Download converted file",
					"description": "Download the result of a completed conversion",
					"parameters": []map[string]interface{}{
						{
							"name":        "job_id",
							"in":          "path",
							"required":    true,
							"description": "Conversion job ID",
							"schema": map[string]string{
								"type": "string",
							},
						},
					},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Converted file",
							"content": map[string]interface{}{
								"application/octet-stream": map[string]interface{}{
									"schema": map[string]string{
										"type":   "string",
										"format": "binary",
									},
								},
							},
						},
					},
				},
			},
		},
		"components": map[string]interface{}{
			"schemas": map[string]interface{}{
				"ConversionJob": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"id": map[string]string{
							"type":        "string",
							"description": "Unique job identifier",
						},
						"input_file": map[string]string{
							"type":        "string",
							"description": "Original filename",
						},
						"input_format": map[string]string{
							"type":        "string",
							"description": "Source format",
						},
						"output_format": map[string]string{
							"type":        "string",
							"description": "Target format",
						},
						"status": map[string]interface{}{
							"type":        "string",
							"description": "Job status",
							"enum":        []string{"queued", "processing", "completed", "failed"},
						},
						"progress": map[string]interface{}{
							"type":        "integer",
							"description": "Progress percentage (0-100)",
							"minimum":     0,
							"maximum":     100,
						},
						"file_size": map[string]string{
							"type":        "integer",
							"description": "File size in bytes",
						},
						"download_url": map[string]string{
							"type":        "string",
							"description": "Download URL when completed",
						},
						"error": map[string]string{
							"type":        "string",
							"description": "Error message if failed",
						},
						"created_at": map[string]string{
							"type":        "integer",
							"description": "Unix timestamp",
						},
						"completed_at": map[string]string{
							"type":        "integer",
							"description": "Unix timestamp",
						},
					},
				},
				"FormatInfo": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"format": map[string]string{
							"type":        "string",
							"description": "Format extension",
						},
						"name": map[string]string{
							"type":        "string",
							"description": "Human-readable name",
						},
						"category": map[string]string{
							"type":        "string",
							"description": "Format category",
						},
						"mime_type": map[string]string{
							"type":        "string",
							"description": "MIME type",
						},
						"converts_to": map[string]interface{}{
							"type":        "array",
							"description": "Compatible target formats",
							"items": map[string]string{
								"type": "string",
							},
						},
					},
				},
				"Error": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"error": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"code": map[string]string{
									"type":        "string",
									"description": "Error code",
								},
								"message": map[string]string{
									"type":        "string",
									"description": "Error message",
								},
							},
						},
					},
				},
			},
		},
	}
}

// ServeOpenAPIJSON serves the OpenAPI spec as JSON
func ServeOpenAPIJSON(w http.ResponseWriter, r *http.Request) {
	spec := OpenAPISpec()
	
	// Format JSON with 2-space indentation per spec
	data, err := json.MarshalIndent(spec, "", "  ")
	if err != nil {
		http.Error(w, "Failed to generate OpenAPI spec", http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	w.Write([]byte("\n")) // Single trailing newline per spec
}

// ServeSwaggerUI serves the Swagger UI HTML page
// See AI.md PART 14: Swagger & GraphQL Theming
func ServeSwaggerUI(w http.ResponseWriter, r *http.Request) {
	// Detect theme from cookie or default to dark
	theme := "dark"
	if cookie, err := r.Cookie("theme"); err == nil && cookie.Value != "" {
		theme = cookie.Value
	}
	
	// Get protocol and host for server URL
	protocol := "http"
	if r.TLS != nil {
		protocol = "https"
	}
	host := r.Host
	
	html := `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>casimg API Documentation</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5/swagger-ui.css">
    <style>
      /* Dark theme - DEFAULT per AI.md PART 14 */
      .swagger-ui.theme-dark {
        background: #282a36;
        color: #f8f8f2;
      }
      .swagger-ui.theme-dark .topbar {
        background: #1e1f29;
        border-bottom: 1px solid #44475a;
      }
      .swagger-ui.theme-dark .info .title,
      .swagger-ui.theme-dark .opblock-tag {
        color: #f8f8f2;
      }
      .swagger-ui.theme-dark .opblock.opblock-get {
        background: rgba(139, 233, 253, 0.1);
        border-color: #8be9fd;
      }
      .swagger-ui.theme-dark .opblock.opblock-get .opblock-summary-method {
        background: #8be9fd;
        color: #282a36;
      }
      .swagger-ui.theme-dark .opblock.opblock-post {
        background: rgba(80, 250, 123, 0.1);
        border-color: #50fa7b;
      }
      .swagger-ui.theme-dark .opblock.opblock-post .opblock-summary-method {
        background: #50fa7b;
        color: #282a36;
      }
      .swagger-ui.theme-dark .opblock.opblock-put {
        background: rgba(255, 184, 108, 0.1);
        border-color: #ffb86c;
      }
      .swagger-ui.theme-dark .opblock.opblock-delete {
        background: rgba(255, 85, 85, 0.1);
        border-color: #ff5555;
      }
      .swagger-ui.theme-dark input,
      .swagger-ui.theme-dark textarea,
      .swagger-ui.theme-dark select {
        background: #44475a;
        color: #f8f8f2;
        border: 1px solid #6272a4;
      }
      .swagger-ui.theme-dark .btn {
        background: #6272a4;
        color: #f8f8f2;
        border: none;
      }
      .swagger-ui.theme-dark .btn.execute {
        background: #bd93f9;
        color: #282a36;
      }
      .swagger-ui.theme-dark .response-col_status {
        color: #50fa7b;
      }
      .swagger-ui.theme-dark table thead tr th,
      .swagger-ui.theme-dark table thead tr td {
        background: #1e1f29;
        color: #f8f8f2;
        border-color: #44475a;
      }
      .swagger-ui.theme-dark .model-box,
      .swagger-ui.theme-dark .responses-wrapper {
        background: #1e1f29;
        border-color: #44475a;
      }
      
      /* Light theme */
      .swagger-ui.theme-light {
        background: #ffffff;
        color: #24292e;
      }
      .swagger-ui.theme-light .topbar {
        background: #f6f8fa;
        border-bottom: 1px solid #e1e4e8;
      }
      
      /* Theme toggle button */
      .theme-toggle {
        position: fixed;
        top: 10px;
        right: 10px;
        z-index: 9999;
        padding: 8px 16px;
        background: #6272a4;
        color: #f8f8f2;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        font-size: 14px;
      }
      .theme-toggle:hover {
        background: #bd93f9;
      }
    </style>
  </head>
  <body>
    <button class="theme-toggle" onclick="toggleTheme()">🌓 Toggle Theme</button>
    <div id="swagger-ui" class="theme-` + theme + `"></div>
    <script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5/swagger-ui-standalone-preset.js"></script>
    <script>
      window.onload = function() {
        const spec = ` + strings.Replace(string(must(json.Marshal(OpenAPISpec()))), "`", "\\`", -1) + `;
        
        // Replace server URL with current host
        spec.servers[0].url = "` + protocol + `://` + host + `";
        
        window.ui = SwaggerUIBundle({
          spec: spec,
          dom_id: '#swagger-ui',
          deepLinking: true,
          presets: [
            SwaggerUIBundle.presets.apis,
            SwaggerUIStandalonePreset
          ],
          layout: "StandaloneLayout"
        });
      };
      
      function toggleTheme() {
        const container = document.getElementById('swagger-ui');
        const currentTheme = container.className.includes('theme-dark') ? 'dark' : 'light';
        const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
        container.className = 'theme-' + newTheme;
        document.cookie = 'theme=' + newTheme + '; path=/; max-age=31536000';
      }
    </script>
  </body>
</html>
`
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

func must(data []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return data
}
