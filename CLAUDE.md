# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

MCPJungle is a self-hosted MCP (Model Context Protocol) Server registry that acts as a centralized proxy for managing multiple MCP servers. It provides both a CLI tool and web dashboard for developers to register, manage, and invoke tools from multiple MCP servers through a single interface.

**Key Architecture**: Client-Server with dual interfaces:
- **Registry Server**: Web API + Dashboard for managing MCP servers (`/api/v0/*`)
- **MCP Proxy Server**: Unified MCP endpoint that proxies to registered servers (`/mcp`)
- **CLI Client**: Command-line interface for server management and tool invocation

## Development Commands

### Building and Running
```bash
# Build single binary for current platform
goreleaser build --single-target --clean --snapshot

# Build for all platforms
goreleaser release --snapshot --clean

# Run server locally (uses SQLite by default)
go run main.go start

# Run with PostgreSQL
export DATABASE_URL=postgres://mcpjungle:mcpjungle@localhost:5432/mcpjungle
go run main.go start

# Run via Docker Compose (recommended for development)
docker-compose up -d
```

### Frontend Development
```bash
cd web
npm install
npm run dev        # Development server
npm run build      # Production build
npm run lint       # ESLint check
```

### CLI Usage
```bash
# Register a MCP server
go run main.go register --name calculator --description "Math tools" --url http://localhost:8000/mcp

# List registered servers and tools
go run main.go list servers
go run main.go list tools

# Invoke a tool (format: server-name/tool-name)
go run main.go invoke calculator/multiply --input '{"a": 100, "b": 50}'

# Remove a server
go run main.go deregister calculator
```

## Core Architecture

### Command Structure (`cmd/`)
- `start.go`: Registry server startup with MCP proxy
- `register.go`: Add MCP servers to registry
- `list.go`: Display servers and available tools
- `invoke.go`: Call tools through the proxy
- `deregister.go`: Remove servers from registry

### Internal Services (`internal/`)
- `api/`: Gin HTTP router and REST endpoint handlers
- `service/`: Business logic for MCP server management and tool proxying
- `model/`: Database models (GORM) for server metadata
- `migrations/`: Database schema evolution

### Key Integration Points
- **MCP Proxy**: Uses `github.com/mark3labs/mcp-go` for MCP protocol handling
- **Database**: GORM with PostgreSQL (production) or SQLite (development)
- **Web Frontend**: Next.js 14 with TypeScript, served by Go server at `/`
- **API Routes**: REST endpoints under `/api/v0/` for server management

### Transport Support
MCPJungle currently supports:
- **HTTP Transport**: Streamable HTTP for remote MCP servers
- **SSE Transport**: Server-Sent Events for real-time communication  
- **Stdio Transport**: Direct process communication for local servers

### Environment Configuration
Key variables in `.env`:
- `DATABASE_URL`: PostgreSQL connection string
- `DEBUG_LEVEL`: Logging level (debug/info/warn/error)
- `PORT`: Server port (default 8080)
- `MCP_CLIENT_TIMEOUT`: Timeout for MCP server calls
- `MCP_CONNECTION_POOL_SIZE`: Connection pooling configuration

### Client Configuration Examples
For AI clients to connect to MCPJungle's unified MCP endpoint:

```json
// Cursor/Claude Desktop
{
  "mcpServers": {
    "mcpjungle": {
      "url": "http://localhost:8080/mcp"
    }
  }
}
```

## Development Notes

### Authentication
- Bearer token support for downstream MCP servers
- Tokens stored in database (marked for encryption improvement)
- Static token auth via `--bearer-token` flag during registration

### Tool Naming Convention
Tools are accessed using canonical names: `<server-name>/<tool-name>`
- Example: `github/git_commit`, `calculator/multiply`
- This namespacing prevents tool name conflicts between servers

### Database Schema
- `mcp_servers`: Server metadata and connection details
- `server_metrics`: Performance and usage tracking (future enhancement)
- Transport-specific fields for HTTP/SSE/Stdio configurations

### Release Process
MCPJungle uses GoReleaser for automated builds and releases:
- Creates binaries for multiple platforms
- Publishes Docker images
- Updates Homebrew formula automatically