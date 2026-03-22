# DTUI - Docker Terminal User Interface

A sleek, terminal-based user interface for managing Docker containers, images, and volumes. Built with Go and the Charm Bubble Tea framework.

## Features

- **Multi-tab Interface**: Switch seamlessly between Containers, Images, Volumes, and Work in Progress views
- **Interactive Management**: Start, stop, remove, and manage Docker resources with keyboard shortcuts
- **Real-time Updates**: Live spinners and progress indicators for long-running operations
- **Pull Images**: Download Docker images directly from the TUI
- **Batch Operations**: Select and operate on multiple resources at once
- **Confirmation Prompts**: Safe deletion with yes/no confirmations
- **Status Messages**: Clear feedback for all operations (success, error, info)
- **Help Overlay**: Press `?` to see all available keybindings

## Screenshots

> _Screenshots coming soon_

## Prerequisites

- Go 1.22 or later
- Docker CLI installed and configured
- Terminal with mouse support (optional, for mouse cell motion)

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/TheAimHero/dtui.git
cd dtui

# Build the application
go build -o dtui .

# Run
./dtui
```

### Run Directly

```bash
go run .
```

## Usage

### Navigation

| Key | Action |
|-----|--------|
| `←` / `→` | Switch between tabs |
| `↑` / `↓` | Navigate list items |
| `↑` / `k` | Move up |
| `↓` / `j` | Move down |

### Container Management

| Key | Action |
|-----|--------|
| `Enter` | Start selected container |
| `ctrl+d` | Stop selected container |
| `d` | Delete selected container (with confirmation) |
| `D` | Force delete (no confirmation) |
| `l` | View container logs |
| `ctrl+r` | Refresh container list |

### Image Management

| Key | Action |
|-----|--------|
| `p` | Pull a new image |
| `d` | Remove selected image |
| `ctrl+p` | Prune unused images |
| `ctrl+r` | Refresh image list |

### Volume Management

| Key | Action |
|-----|--------|
| `d` | Remove selected volume |
| `D` | Force remove (no confirmation) |
| `ctrl+p` | Prune unused volumes |
| `ctrl+r` | Refresh volume list |

### General

| Key | Action |
|-----|--------|
| `Space` | Toggle selection |
| `?` | Toggle help overlay |
| `q` | Quit |
| `ctrl+c` | Quit (alternative) |

## Architecture

DTUI follows a clean architecture pattern:

```
cmd/tui/           # Application entry points and UI modules
  ├── tabs/        # Tab management (Container, Image, Volume, WIP)
  ├── manageContainer/
  ├── manageImage/
  ├── manageVolume/
  └── wip/        # Work in progress features

internal/
  ├── docker/      # Docker API/CLI abstraction layer
  ├── ui/          # Shared UI components
  │   └── components/
  └── utils/       # Utility functions
```

### Key Components

- **Bubble Tea**: Handles the Elm-style TEA (Terminal, Elang Architecture) pattern
- **Lipgloss**: Provides terminal styling and colors
- **BaseModel**: Shared model structure embedded across all management modules
- **Service Interfaces**: Clean abstraction over Docker operations

For detailed architecture documentation, see [docs/architecture.md](docs/architecture.md).

## Development

### Project Structure

```
dtui/
├── cmd/tui/           # CLI commands
│   ├── tabs/          # Main TUI orchestration
│   ├── manageContainer/
│   ├── manageImage/
│   ├── manageVolume/
│   └── wip/
├── internal/
│   ├── docker/        # Docker service implementations
│   ├── ui/           # Shared UI components
│   └── utils/        # Utilities
├── docs/             # Documentation
└── main.go           # Entry point
```

### Build Commands

```bash
# Build
go build -o dtui .

# Run with hot reload (requires air)
air

# Run tests
go test ./...

# Tidy dependencies
go mod tidy
```

### Code Style

- Follow Go conventions
- Use `gofmt` for formatting
- Group imports: standard library → third party → internal

## Configuration

DTUI uses sensible defaults and requires no configuration file. Terminal dimensions are automatically detected.

## Troubleshooting

### Docker Not Found

Ensure Docker is installed and the CLI is in your PATH:
```bash
docker --version
```

### Permission Denied

If you encounter permission issues, ensure your user has Docker permissions:
```bash
sudo usermod -aG docker $USER
# Log out and log back in for changes to take effect
```

### Build Errors

Make sure you have Go 1.22+:
```bash
go version
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Roadmap

- [ ] Build Mode (build images from Dockerfiles)
- [ ] Image Mode (detailed image management)
- [ ] Log Mode (streaming container logs viewer)
- [ ] Configuration file support
- [ ] Customizable keybindings
- [ ] Dark/light theme toggle

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [Charm](https://charm.sh/) for the excellent Bubble Tea and Lipgloss libraries
- [Docker](https://www.docker.com/) for the container platform
- All contributors and users of this project
