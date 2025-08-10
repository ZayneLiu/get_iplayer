# get_iplayer Go Implementation

This directory contains a Go 1.24 implementation of the get_iplayer BBC iPlayer/BBC Sounds downloading tool.

## Status

This is a **work-in-progress** Go port of the original Perl get_iplayer tool. It currently provides:

- ✅ Basic CLI interface compatible with original get_iplayer
- ✅ Command structure (--help, --version, --long-help)
- ✅ Search functionality (with mock data for demonstration)
- ✅ Configuration management
- ✅ Programme listing with both short and long formats
- ✅ Basic download preparation (structure only)
- 🔄 Real BBC iPlayer/Sounds API integration (in development)
- 🔄 Actual downloading functionality (in development)

## Building

```bash
go build -o get_iplayer_go .
```

## Usage

The Go implementation aims to be compatible with the original Perl version:

```bash
# Show version
./get_iplayer_go --version

# Show help
./get_iplayer_go --help
./get_iplayer_go --long-help

# Search (currently with mock data)
./get_iplayer_go "doctor who"
./get_iplayer_go --long ".*"
./get_iplayer_go --channel="BBC One" ".*"

# Download preparation (structure only)
./get_iplayer_go --get 123
./get_iplayer_go --pid=b01sc0wf
```

## Architecture

```
cmd/           - CLI interface and commands
pkg/
  api/         - BBC iPlayer/Sounds API client
  config/      - Configuration management
  models/      - Data structures
internal/
  cache/       - Programme cache (planned)
  downloader/  - Download functionality (planned)
```

## Go 1.24 Features Used

- Modern Go module structure
- Native error handling
- Type safety for BBC API data structures
- Concurrent operations support (planned for downloads)

## Development Status

This implementation serves as a foundation for a full Go port. For production use, please continue using the original Perl version at `./get_iplayer`.

The mock data demonstrates the intended functionality and CLI compatibility while the real API integration is developed.

## Testing

```bash
go test -v ./...
```

## Contributing

This Go implementation follows the same goals as the original:
- Maintain CLI compatibility  
- Provide reliable BBC content access
- Support cross-platform usage
- Keep the codebase maintainable

For the original Perl implementation, see the main repository documentation.