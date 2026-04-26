# API Documentation

casimg provides a comprehensive REST API for file conversion.

## Base URL

```
http://localhost:64580/api/v1
```

## Authentication

Anonymous usage is supported with 24-hour result retention. For persistent storage and history, authentication is required.

## Endpoints

### Health Check

```http
GET /api/v1/healthz
```

Response:
```json
{
  "status": "healthy",
  "version": "0.1.0",
  "mode": "production",
  "uptime": "2d 5h 30m"
}
```

### List Formats

Get supported formats and conversion capabilities:

```http
GET /api/v1/formats
```

Response:
```json
{
  "formats": [
    {
      "format": "pdf",
      "name": "PDF Document",
      "category": "document",
      "extensions": [".pdf"],
      "converts_to": ["docx", "txt", "html"]
    }
  ]
}
```

### Upload and Convert

Submit a file for conversion:

```http
POST /api/v1/convert
Content-Type: multipart/form-data

file: [binary data]
output_format: docx
quality: high (optional)
```

Response:
```json
{
  "id": "abc123",
  "status": "queued",
  "input_format": "pdf",
  "output_format": "docx",
  "created_at": 1640000000,
  "expires_at": 1640086400
}
```

### Check Status

Poll for conversion progress:

```http
GET /api/v1/convert/{id}
```

Response:
```json
{
  "id": "abc123",
  "status": "completed",
  "progress": 100,
  "download_url": "/api/v1/convert/abc123/download",
  "completed_at": 1640001000
}
```

Status values: `queued`, `processing`, `completed`, `failed`

### Download Result

```http
GET /api/v1/convert/{id}/download
```

Returns the converted file.

### Batch Conversion

Convert multiple files:

```http
POST /api/v1/convert/batch
Content-Type: multipart/form-data

files[]: [binary data 1]
files[]: [binary data 2]
output_format: pdf
```

Response:
```json
{
  "batch_id": "batch_xyz",
  "jobs": [
    {"id": "abc123", "status": "queued"},
    {"id": "def456", "status": "queued"}
  ]
}
```

### Cancel Conversion

```http
DELETE /api/v1/convert/{id}
```

Response:
```json
{
  "success": true,
  "message": "Conversion cancelled"
}
```

## Error Responses

All errors follow this format:

```json
{
  "error": {
    "code": "ERR_FILE_TOO_LARGE",
    "message": "File size exceeds maximum limit",
    "details": {}
  }
}
```

Common error codes:
- `ERR_FILE_TOO_LARGE` - File exceeds size limit
- `ERR_INVALID_FORMAT` - Unsupported format
- `ERR_CONVERSION_FAILED` - Conversion error
- `ERR_NOT_FOUND` - Job not found
- `ERR_EXPIRED` - Result has expired

## Rate Limiting

API requests are rate limited:
- Anonymous: 60 requests/minute
- Authenticated: 300 requests/minute

Rate limit headers:
```
X-RateLimit-Limit: 60
X-RateLimit-Remaining: 45
X-RateLimit-Reset: 1640000060
```

## OpenAPI/Swagger

Interactive API documentation:

```
http://localhost:64580/openapi
```

## GraphQL

GraphQL endpoint:

```
http://localhost:64580/graphql
```

GraphiQL interface:

```
http://localhost:64580/graphql
```

## Next Steps

- [Admin Panel](admin.md) - Manage conversions
- [Configuration](configuration.md) - Configure limits
