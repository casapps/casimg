package handler

import (
"encoding/json"
"fmt"
"io"
"net/http"
"os"
"path/filepath"
"time"

"github.com/casapps/casimg/src/server/model"
)

type ConvertHandler struct {
dataDir string
}

func NewConvertHandler(dataDir string) *ConvertHandler {
uploadsDir := filepath.Join(dataDir, "uploads")
convertedDir := filepath.Join(dataDir, "converted")
os.MkdirAll(uploadsDir, 0755)
os.MkdirAll(convertedDir, 0755)

return &ConvertHandler{
dataDir: dataDir,
}
}

func (h *ConvertHandler) HandleUpload(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodPost {
http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
return
}

r.ParseMultipartForm(500 << 20) // 500MB max

file, header, err := r.FormFile("file")
if err != nil {
writeJSON(w, http.StatusBadRequest, map[string]interface{}{
"error": map[string]string{
"code":    "ERR_NO_FILE",
"message": "No file uploaded",
},
})
return
}
defer file.Close()

outputFormat := r.FormValue("output_format")
if outputFormat == "" {
writeJSON(w, http.StatusBadRequest, map[string]interface{}{
"error": map[string]string{
"code":    "ERR_NO_OUTPUT_FORMAT",
"message": "Output format required",
},
})
return
}

jobID := generateID()
inputPath := filepath.Join(h.dataDir, "uploads", jobID+"_"+header.Filename)

dst, err := os.Create(inputPath)
if err != nil {
writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
"error": map[string]string{
"code":    "ERR_SAVE_FAILED",
"message": "Failed to save file",
},
})
return
}
defer dst.Close()

written, err := io.Copy(dst, file)
if err != nil {
writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
"error": map[string]string{
"code":    "ERR_SAVE_FAILED",
"message": "Failed to save file",
},
})
return
}

inputExt := filepath.Ext(header.Filename)
if len(inputExt) > 0 {
inputExt = inputExt[1:]
}

now := time.Now().Unix()
job := &model.ConversionJob{
ID:           jobID,
InputFile:    header.Filename,
InputFormat:  inputExt,
OutputFormat: outputFormat,
Status:       "queued",
Progress:     0,
FileSize:     written,
CreatedAt:    now,
ExpiresAt:    now + 86400, // 24 hours
Options:      make(map[string]interface{}),
}

go h.processConversion(job, inputPath)

writeJSON(w, http.StatusCreated, job)
}

func (h *ConvertHandler) HandleStatus(w http.ResponseWriter, r *http.Request) {
jobID := filepath.Base(r.URL.Path)

writeJSON(w, http.StatusOK, map[string]interface{}{
"id":       jobID,
"status":   "completed",
"progress": 100,
"download_url": fmt.Sprintf("/api/v1/convert/%s/download", jobID),
})
}

func (h *ConvertHandler) HandleDownload(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/octet-stream")
w.Write([]byte("Converted file placeholder"))
}

func (h *ConvertHandler) HandleFormats(w http.ResponseWriter, r *http.Request) {
formats := []model.FormatInfo{
{
Format:     "pdf",
Name:       "PDF Document",
Category:   "document",
MimeType:   "application/pdf",
Extensions: []string{".pdf"},
ConvertsTo: []string{"docx", "txt", "html", "png", "jpg"},
},
{
Format:     "docx",
Name:       "Word Document",
Category:   "document",
MimeType:   "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
Extensions: []string{".docx"},
ConvertsTo: []string{"pdf", "txt", "html", "odt"},
},
{
Format:     "jpg",
Name:       "JPEG Image",
Category:   "image",
MimeType:   "image/jpeg",
Extensions: []string{".jpg", ".jpeg"},
ConvertsTo: []string{"png", "gif", "webp", "bmp", "pdf"},
},
{
Format:     "png",
Name:       "PNG Image",
Category:   "image",
MimeType:   "image/png",
Extensions: []string{".png"},
ConvertsTo: []string{"jpg", "gif", "webp", "bmp", "pdf"},
},
{
Format:        "mp3",
Name:          "MP3 Audio",
Category:      "audio",
MimeType:      "audio/mpeg",
Extensions:    []string{".mp3"},
ConvertsTo:    []string{"wav", "flac", "aac", "ogg"},
HasQuality:    true,
HasResolution: false,
},
{
Format:        "mp4",
Name:          "MP4 Video",
Category:      "video",
MimeType:      "video/mp4",
Extensions:    []string{".mp4"},
ConvertsTo:    []string{"avi", "mkv", "mov", "webm"},
HasQuality:    true,
HasResolution: true,
},
}

writeJSON(w, http.StatusOK, map[string]interface{}{
"formats": formats,
})
}

func (h *ConvertHandler) processConversion(job *model.ConversionJob, inputPath string) {
job.Status = "processing"
job.Progress = 50

time.Sleep(2 * time.Second)

job.Status = "completed"
job.Progress = 100
job.CompletedAt = time.Now().Unix()
job.DownloadURL = fmt.Sprintf("/api/v1/convert/%s/download", job.ID)
}

func generateID() string {
return fmt.Sprintf("%d", time.Now().UnixNano())
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(status)
json.NewEncoder(w).Encode(data)
}
