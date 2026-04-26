# casimg Documentation

Welcome to casimg - a self-hosted file conversion service that provides unlimited format conversions without ads or feature gates.

## What is casimg?

casimg is a free, open-source alternative to commercial file conversion services like CloudConvert and Convertio. It runs entirely on your own infrastructure, ensuring complete privacy and control over your data.

## Key Features

- **Comprehensive Format Support**: Convert between hundreds of file formats
- **No Limitations**: Unlimited conversions, no file size restrictions
- **Privacy First**: Files never leave your server
- **No Ads or Tracking**: Completely ad-free experience
- **Batch Processing**: Convert multiple files simultaneously
- **Full API Access**: REST, GraphQL, and Swagger interfaces
- **Open Source**: MIT licensed, fully auditable

## Supported Formats

casimg supports conversion between the following format categories:

- **Documents**: PDF, DOCX, ODT, TXT, HTML, EPUB, RTF, Markdown
- **Images**: JPEG, PNG, GIF, BMP, TIFF, WEBP, SVG, ICO
- **Audio**: MP3, WAV, FLAC, AAC, OGG, M4A, WMA
- **Video**: MP4, AVI, MKV, MOV, WEBM, FLV, WMV
- **Archives**: ZIP, TAR, GZ, BZ2, 7Z
- **Spreadsheets**: XLSX, ODS, CSV, XLS
- **Presentations**: PPTX, ODP, PPT

## Quick Start

The fastest way to get started is with Docker:

```bash
docker run -d \
  --name casimg \
  -p 64580:80 \
  -v ./config:/config \
  -v ./data:/data \
  ghcr.io/casapps/casimg:latest
```

Then visit http://localhost:64580 in your browser.

## Next Steps

- [Installation Guide](installation.md) - Detailed installation instructions
- [Configuration](configuration.md) - Configure casimg for your needs
- [API Documentation](api.md) - Integrate with your applications
- [Admin Panel](admin.md) - Manage your casimg instance

## Support

- **Issues**: [GitHub Issues](https://github.com/casapps/casimg/issues)
- **Discussions**: [GitHub Discussions](https://github.com/casapps/casimg/discussions)
- **Source Code**: [GitHub Repository](https://github.com/casapps/casimg)

## License

casimg is released under the MIT License. See [LICENSE.md](https://github.com/casapps/casimg/blob/main/LICENSE.md) for details.
