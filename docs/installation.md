# Installation

casimg can be installed using Docker (recommended) or as a standalone binary.

## Docker Installation

### Using docker run

```bash
docker run -d \
  --name casimg \
  -p 64580:80 \
  -e TZ=America/New_York \
  -e MODE=production \
  -v ./config:/config \
  -v ./data:/data \
  --restart unless-stopped \
  ghcr.io/casapps/casimg:latest
```

### Using docker-compose

1. Create `docker-compose.yml`:

```yaml
version: '3.8'

services:
  casimg:
    image: ghcr.io/casapps/casimg:latest
    container_name: casimg
    restart: unless-stopped
    ports:
      - "64580:80"
    environment:
      - TZ=America/New_York
      - MODE=production
    volumes:
      - ./config:/config
      - ./data:/data
```

2. Start the service:

```bash
docker-compose up -d
```

## Binary Installation

### Download

Download the latest release for your platform from the [releases page](https://github.com/casapps/casimg/releases).

### Linux

```bash
# Download
wget https://github.com/casapps/casimg/releases/latest/download/casimg-linux-amd64

# Make executable
chmod +x casimg-linux-amd64

# Move to system path
sudo mv casimg-linux-amd64 /usr/local/bin/casimg

# Install conversion tools (optional)
sudo apt install imagemagick ffmpeg libreoffice pandoc ghostscript
```

### macOS

```bash
# Download
wget https://github.com/casapps/casimg/releases/latest/download/casimg-darwin-amd64

# Make executable
chmod +x casimg-darwin-amd64

# Move to system path
sudo mv casimg-darwin-amd64 /usr/local/bin/casimg

# Install conversion tools (optional)
brew install imagemagick ffmpeg libreoffice pandoc ghostscript
```

### Windows

1. Download `casimg-windows-amd64.exe` from the releases page
2. Rename to `casimg.exe`
3. Add to your PATH or run directly

## Running

### Binary

```bash
casimg --mode production --port 64580
```

### Systemd Service

Create `/etc/systemd/system/casimg.service`:

```ini
[Unit]
Description=casimg File Conversion Service
After=network.target

[Service]
Type=simple
User=casimg
ExecStart=/usr/local/bin/casimg --mode production
Restart=always

[Install]
WantedBy=multi-user.target
```

Enable and start:

```bash
sudo systemctl daemon-reload
sudo systemctl enable casimg
sudo systemctl start casimg
```

## Verification

Check that casimg is running:

```bash
curl http://localhost:64580/api/v1/healthz
```

You should see a JSON response indicating the service is healthy.

## Next Steps

- [Configuration](configuration.md) - Configure your installation
- [Admin Panel](admin.md) - Access the admin interface
- [API Documentation](api.md) - Start using the API
