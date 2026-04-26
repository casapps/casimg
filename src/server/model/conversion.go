package model

type ConversionJob struct {
ID           string                 `json:"id"`
UserID       string                 `json:"user_id"`
InputFile    string                 `json:"input_file"`
InputFormat  string                 `json:"input_format"`
OutputFormat string                 `json:"output_format"`
Status       string                 `json:"status"`
Progress     int                    `json:"progress"`
DownloadURL  string                 `json:"download_url"`
Error        string                 `json:"error,omitempty"`
Options      map[string]interface{} `json:"options"`
FileSize     int64                  `json:"file_size"`
CreatedAt    int64                  `json:"created_at"`
CompletedAt  int64                  `json:"completed_at,omitempty"`
ExpiresAt    int64                  `json:"expires_at"`
}

type FormatInfo struct {
Format        string   `json:"format"`
Name          string   `json:"name"`
Category      string   `json:"category"`
MimeType      string   `json:"mime_type"`
Extensions    []string `json:"extensions"`
ConvertsTo    []string `json:"converts_to"`
HasQuality    bool     `json:"has_quality"`
HasResolution bool     `json:"has_resolution"`
}

type BatchConversion struct {
ID         string   `json:"id"`
UserID     string   `json:"user_id"`
Jobs       []string `json:"job_ids"`
Status     string   `json:"status"`
TotalFiles int      `json:"total_files"`
Completed  int      `json:"completed"`
Failed     int      `json:"failed"`
CreatedAt  int64    `json:"created_at"`
}
