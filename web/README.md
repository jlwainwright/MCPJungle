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

### Completed Features ✅
- [x] **Console Logging**: Detailed API call logging with grouped output and visual indicators
- [x] **Tool-level Configuration**: Expandable server rows with individual tool toggles
- [x] **Modern Web Dashboard**: Responsive UI with server toggle matrix
- [x] **Client Management**: Multi-client support (Claude Desktop, Cursor, Windsurf, Cline)
- [x] **Database Foundation**: Complete analytics schema and migration system
- [x] **Docker Integration**: Full stack deployment with Docker Compose
- [x] **Future Architecture**: Foundations for server discovery and health monitoring

### High Priority (Next Quarter)
- [ ] **Docker MCP Integration**: Leverage Docker's MCP Catalog and Toolkit as deployment backend
- [ ] **Server Discovery**: Pull from GitHub registry + Docker Hub catalog (100+ verified servers)
- [ ] **Real-time Health Monitoring**: Functional heartbeat indicators with container status
- [ ] **Anthropic SDK Integration**: Token usage tracking and cost calculations
- [ ] **Product Rebrand**: Transition to commercial product (ToolFlow/AgentHub/ControlPlane AI)
- [ ] **SaaS Infrastructure**: Authentication, billing, multi-tenancy for commercial offering

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

### Analytics & Business Intelligence
- [x] **Analytics Database Schema**: Complete metrics collection framework
- [x] **Cost Calculation Engine**: Anthropic pricing models and token cost tracking
- [ ] **Real-time Dashboards**: Visual analytics for usage patterns and costs
- [ ] **Token Usage Tracking**: Live monitoring per model, client, and tool
- [ ] **Business Intelligence**: ROI analysis, optimization recommendations
- [ ] **Export & Reporting**: CSV/PDF reports for accounting and billing
- [ ] **Alert System**: Cost thresholds, usage spikes, performance issues

### Commercial Product Features
- [ ] **Multi-tenancy**: Isolated environments for different organizations
- [ ] **RBAC & SSO**: Role-based access control and single sign-on
- [ ] **API Management**: Rate limiting, usage quotas, webhook integrations
- [ ] **Enterprise Integrations**: Slack notifications, JIRA ticketing, Datadog monitoring
- [ ] **Compliance**: SOC2, GDPR, audit logs, data governance
- [ ] **White-label Options**: Custom branding for enterprise customers

### Advanced Features
- [ ] **Configuration Templates**: Pre-defined setups for common workflows  
- [ ] **Bulk Operations**: Multi-select server and client management
- [ ] **Dark Mode**: Theme switching and accessibility improvements
- [ ] **Marketplace**: Plugin ecosystem for custom integrations

### Performance & UX
- [ ] **Real-time Updates**: WebSocket integration for live status updates
- [ ] **Keyboard Shortcuts**: Quick actions and navigation
- [ ] **Search & Filtering**: Advanced server and tool discovery
- [ ] **Drag & Drop**: Intuitive server assignment to clients