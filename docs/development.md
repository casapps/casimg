# Development

Guide for developers working on casimg.

## Building from Source

```bash
git clone https://github.com/casapps/casimg.git
cd casimg
make build
```

## Project Structure

```
casimg/
├── src/              # Go source code
│   ├── main.go       # Entry point
│   ├── config/       # Configuration
│   ├── server/       # HTTP server
│   └── ...
├── docker/           # Docker files
├── docs/             # Documentation
└── tests/            # Tests
```

## Running Tests

```bash
make test
```

## Local Development

```bash
# Build and run
go build -o casimg ./src
./casimg --mode development --port 64580

# Or use make
make dev
```

## Docker Development

```bash
cd docker
docker-compose -f docker-compose.dev.yml up
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests
5. Submit a pull request

## Code Style

Follow standard Go conventions:
- Run `go fmt`
- Run `go vet`
- Add tests for new features

## Documentation

Update documentation when adding features:
- Update AI.md PART 36
- Update README.md
- Update docs/ files
- Update API documentation
