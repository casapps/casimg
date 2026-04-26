# Configuration

casimg is configured via `server.yml` and environment variables.

## Configuration File

**Location:**
- Docker: `/config/server.yml`
- Binary: `/etc/casimg/server.yml` (Linux), `~/.config/casimg/server.yml` (user)

## Basic Configuration

```yaml
server:
  address: 0.0.0.0
  port: 64580
  debug: false

conversion:
  max_file_size: 524288000  # 500MB
  timeout: 300              # 5 minutes
  max_concurrent: 5
  retention_hours: 24
  cleanup_interval: 3600    # 1 hour

cluster:
  enabled: false
```

## Environment Variables

All settings can be overridden with environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `CASIMG_MODE` | Application mode (production/development) | `production` |
| `CASIMG_PORT` | Listen port | `64580` |
| `CASIMG_ADDRESS` | Listen address | `0.0.0.0` |
| `TZ` | Timezone | `America/New_York` |

## Command Line Flags

```bash
casimg --help
```

Options:
- `--mode {production|development}` - Application mode
- `--port {port}` - Listen port
- `--address {address}` - Listen address
- `--config {dir}` - Configuration directory
- `--data {dir}` - Data directory
- `--log {dir}` - Log directory
- `--debug` - Enable debug mode

## Conversion Settings

### File Size Limits

Control maximum upload size:

```yaml
conversion:
  max_file_size: 1073741824  # 1GB
```

### Timeouts

Set conversion timeout per job:

```yaml
conversion:
  timeout: 600  # 10 minutes
```

### Concurrency

Limit simultaneous conversions:

```yaml
conversion:
  max_concurrent: 10
```

### File Retention

Configure how long converted files are kept:

```yaml
conversion:
  retention_hours: 48  # 2 days
```

## Cluster Mode

For high availability:

```yaml
cluster:
  enabled: true
  database:
    type: postgresql
    host: localhost
    port: 5432
    name: casimg
    user: casimg
    password: secret
```

## Security

### SSL/TLS

```yaml
server:
  ssl:
    enabled: true
    cert: /config/cert.pem
    key: /config/key.pem
```

### Rate Limiting

```yaml
server:
  rate_limit:
    enabled: true
    requests_per_minute: 60
```

## Logging

```yaml
logging:
  level: info
  format: json
  file: /data/log/server.log
  max_size: 100  # MB
  max_age: 30    # days
```

## Next Steps

- [Admin Panel](admin.md) - Configure via web interface
- [API Documentation](api.md) - Use the API
