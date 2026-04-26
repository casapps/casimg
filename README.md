# casimg

**Self-hosted file conversion service** - A free, open-source alternative to CloudConvert and Convertio with unlimited features, no ads, and no feature gates.

[![Release](https://img.shields.io/github/v/release/casapps/casimg)](https://github.com/casapps/casimg/releases)
[![Docker](https://img.shields.io/badge/docker-ghcr.io-blue)](https://ghcr.io/casapps/casimg)
[![License](https://img.shields.io/github/license/casapps/casimg)](LICENSE.md)

## Features

- **Comprehensive Format Support**: Documents, images, audio, video, archives, spreadsheets, presentations
- **No Limits**: Unlimited conversions, no file size restrictions
- **Privacy First**: Self-hosted, files never leave your infrastructure
- **No Ads or Tracking**: Completely ad-free and privacy-respecting
- **Batch Conversion**: Convert multiple files in one request
- **API & Web Interface**: Full REST API, GraphQL, and user-friendly web UI
- **Open Source**: MIT License, fully auditable code

## Quick Start

### Docker (Recommended)

```bash
# Pull and run
docker run -d \
  --name casimg \
  -p 64580:80 \
  -v ./config:/config \
  -v ./data:/data \
  ghcr.io/casapps/casimg:latest
```

Or use docker-compose:

```bash
cd docker
docker-compose up -d
```

Access at: http://localhost:64580

### Binary Installation

Download the latest release for your platform:

```bash
# Linux
wget https://github.com/casapps/casimg/releases/latest/download/casimg-linux-amd64
chmod +x casimg-linux-amd64
./casimg-linux-amd64

# macOS
wget https://github.com/casapps/casimg/releases/latest/download/casimg-darwin-amd64
chmod +x casimg-darwin-amd64
./casimg-darwin-amd64

# Windows
# Download casimg-windows-amd64.exe from releases
casimg-windows-amd64.exe
```

## Supported Formats

### Documents
PDF, DOCX, ODT, TXT, HTML, EPUB, RTF, Markdown

### Images
JPEG, PNG, GIF, BMP, TIFF, WEBP, SVG, ICO, PSD

### Audio
MP3, WAV, FLAC, AAC, OGG, M4A, WMA

### Video
MP4, AVI, MKV, MOV, WEBM, FLV, WMV

### Archives
ZIP, TAR, GZ, BZ2, 7Z, RAR

### Office
XLSX, ODS, CSV, XLS, PPTX, ODP, PPT

## API Usage

### Upload and Convert

```bash
# Upload file
curl -X POST http://localhost:64580/api/v1/convert \
  -F "file=@document.pdf" \
  -F "output_format=docx"

# Response includes job ID
{
  "id": "abc123",
  "status": "queued",
  "download_url": ""
}

# Check status
curl http://localhost:64580/api/v1/convert/abc123

# Download result when complete
curl -O http://localhost:64580/api/v1/convert/abc123/download
```

### List Supported Formats

```bash
curl http://localhost:64580/api/v1/formats
```

## Configuration

Configuration file: `/config/server.yml` (Docker) or `/etc/casimg/server.yml` (binary)

```yaml
server:
  address: 0.0.0.0
  port: 64580
  debug: false

conversion:
  max_file_size: 524288000  # 500MB
  timeout: 300
  max_concurrent: 5
  retention_hours: 24

cluster:
  enabled: false
```

## Building from Source

```bash
git clone https://github.com/casapps/casimg.git
cd casimg
make build
```

Binaries will be in `binaries/` directory.

## Documentation

- [Installation Guide](docs/installation.md)
- [Configuration](docs/configuration.md)
- [API Documentation](docs/api.md)
- [Admin Panel](docs/admin.md)
- [Development](docs/development.md)

Full documentation: https://casapps-casimg.readthedocs.io

## Requirements

### Docker
- Docker 20.10+
- Docker Compose 2.0+

### Binary
- Linux/macOS/Windows/FreeBSD
- Conversion tools (optional, enables formats):
  - ImageMagick (image conversions)
  - FFmpeg (audio/video conversions)
  - LibreOffice (document conversions)
  - Pandoc (markup conversions)

## License

MIT License - see [LICENSE.md](LICENSE.md)

## Author

casapps - https://github.com/casapps

## Contributing

Contributions welcome! Please read [CONTRIBUTING.md](.github/CONTRIBUTING.md) first.

## Support

- Issues: https://github.com/casapps/casimg/issues
- Discussions: https://github.com/casapps/casimg/discussions
