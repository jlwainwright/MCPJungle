# MCPJungle Web Dashboard

A modern web interface for managing MCP servers across different AI clients.

## Features

- **Server Management**: View all registered MCP servers with their status and tools
- **Client Configuration**: Toggle servers on/off for different AI clients (Claude Desktop, Cursor, Windsurf, Cline)
- **Config Generation**: Automatically generate client-specific configuration files
- **Real-time Updates**: Live toggle functionality to manage server access per client

## Development

```bash
# Install dependencies
npm install

# Run development server
npm run dev
```

The development server will start on http://localhost:3001

## Production

```bash
# Build the application
npm run build

# Start production server  
npm start
```

## Docker

```bash
# Build and run with Docker Compose
docker-compose up --build mcpjungle-web
```

## API Integration

The web interface connects to the MCPJungle backend API at:
- Development: `http://localhost:8080/api/v0`
- Production: `/api/v0` (proxied through Next.js)

## Architecture

- **Frontend**: Next.js 14 with TypeScript and Tailwind CSS
- **UI Components**: Headless UI and Heroicons
- **State Management**: React hooks (useState, useEffect)
- **API Client**: Native fetch with TypeScript interfaces

## TODO: Future Enhancements

### Immediate Improvements
- [x] **Console Logging**: Add detailed API call logging to show server communication in browser console
- [x] **Tool-level Configuration**: Expandable server rows with individual tool toggles
- [ ] **Server Discovery**: Pull maintained servers from https://github.com/modelcontextprotocol/servers
- [ ] **Real-time Health Monitoring**: Functional heartbeat indicators for server status
- [ ] **Error Handling**: Improve error states and user feedback for failed API calls
- [ ] **Loading States**: Add skeleton loaders and better loading indicators

### Client Support Expansion
- [ ] **Claude Code CLI**: Add support for Claude Code CLI configuration
- [ ] **Roo.Code**: Integrate Roo.Code AI assistant configuration
- [ ] **Kilo.Code**: Support for Kilo.Code development environment

### Server Discovery & Management
- [ ] **MCP Server Registry**: Pull from https://github.com/modelcontextprotocol/servers
- [ ] **Server Search**: Search and filter available MCP servers by category/functionality
- [ ] **One-click Installation**: Install servers directly from the registry
- [ ] **Server Marketplace**: Browse curated server collections
- [ ] **Version Management**: Handle server updates and versioning
- [ ] **Dependency Resolution**: Auto-install server dependencies

### Health & Monitoring
- [ ] **Real-time Heartbeat**: Functional green/red dots showing live server status
- [ ] **Health Dashboard**: Comprehensive server health monitoring
- [ ] **Performance Metrics**: Response times, success rates, error tracking
- [ ] **Alerting System**: Notifications for server failures or issues
- [ ] **Historical Monitoring**: Server uptime and performance trends

### Advanced Features
- [ ] **Usage Analytics**: Track tool usage patterns across clients
- [ ] **Bulk Operations**: Multi-select server management
- [ ] **Configuration Templates**: Pre-defined client setups for common workflows
- [ ] **Import/Export**: Backup and restore client configurations
- [ ] **Dark Mode**: Theme switching support
- [ ] **Multi-user Support**: User authentication and role-based access

### Performance & UX
- [ ] **Real-time Updates**: WebSocket integration for live status updates
- [ ] **Keyboard Shortcuts**: Quick actions and navigation
- [ ] **Search & Filtering**: Advanced server and tool discovery
- [ ] **Drag & Drop**: Intuitive server assignment to clients