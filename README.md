<h1 align="center">
  :deciduous_tree: MCPJungle :deciduous_tree:
</h1>
<p align="center">
  Self-hosted MCP Server registry for your private AI agents
</p>

MCPJungle is a single source-of-truth registry for all [Model Context Protocol](https://modelcontextprotocol.io/introduction) based Servers running in your Organisation.

🧑‍💻 Developers use it to register & manage MCP servers and the tools they provide from a central place.

🤖 AI Agents use it to discover and consume all these tools from a single MCP Server.

![diagram](./assets/mcpjungle-diagram/mcpjungle-diagram.png)

<p align="center">MCPJungle is the only MCP Server your agents need to connect to!</p>

## Who should use MCPJungle?
1. Devs using MCP Clients like Claude, Cursor, & Windsurf that need to connect to **multiple** MCP servers for calling tools.
2. Devs building AI Agents that need to access **multiple** MCP servers for calling tools.
3. People who want to view and manage all their MCP servers from one centralized place. Secure & Private 🔒

## Installation

> [!WARNING]
> MCPJungle is **BETA** software.
>
> We're actively working to make it production-ready.
> You can provide your feedback by [creating an Issue](https://github.com/duaraghav8/MCPJungle/issues) in this repository.

MPCJungle is shipped as a stand-alone binary.

You can either download it from the [Releases](https://github.com/duaraghav8/MCPJungle/releases) Page or use [Homebrew](https://brew.sh/) to install it:

```bash
$ brew install duaraghav8/mcpjungle/mcpjungle
```

Verify your installation by running

```bash
$ mcpjungle version
```

> [!IMPORTANT]
> On MacOS, you will have to use homebrew because the compiled binary is not [Notarized](https://developer.apple.com/documentation/security/notarizing-macos-software-before-distribution).


## Usage

MCPJungle has a Client-Server architecture and the binary lets you run both a Server and a Client.

### Server
For running the MCPJungle server locally, docker compose is the recommended way:
```shell
curl -O https://raw.githubusercontent.com/duaraghav8/MCPJungle/refs/heads/main/docker-compose.yaml
docker-compose up -d
```

Otherwise, you can run the server directly using the binary:
```bash
$ mcpjungle start
```

This starts the main registry server responsible for managing all MCP servers. It is accessible on port `8080` by default.

The server also exposes its own MCP server at `/mcp` for AI Agents to discover and call Tools provided by the registered MCP Servers.

It relies on a database and by default, creates a SQLite DB in the current working directory.
Alternatively, you can supply a DSN for a Postgresql database to the server:

```bash
$ export DATABASE_URL=postgres://admin:root@localhost:5432/mcpjungle_db
$ mcpjungle start
```

If you use docker-compose, the DB is automatically created and managed for you.

### Client
Once the server is up, you can use the CLI to interact with it.

Let's say you're already running a MCP server locally at `http://127.0.0.1:8000/mcp` which provides basic math tools like `add`, `subtract`, etc.

You can register this MCP server with MCPJungle:
```bash
$ mcpjungle register --name calculator --description "Provides some basic math tools" --url http://127.0.0.1:8000/mcp
```

If you used docker-compose to run the server, and you're not on Linux, you will have to use `host.docker.internal` instead of your local loopback address.
```bash
$ mcpjungle register --name calculator --description "Provides some basic math tools" --url http://host.docker.internal:8000/mcp
```

The registry will now start tracking this MCP server and load its tools.

![register a MCP server in MCPJungle](./assets/register-mcp-server.png)

**Note**: MCPJungle currently only supports MCP Servers using the [Streamable HTTP Transport](https://modelcontextprotocol.io/specification/2025-03-26/basic/transports#streamable-http).

All tools provided by this server are now accessible via MCPJungle:

```bash
$ mcpjungle list tools

# Check tool usage
$ mcpjungle usage calculator/multiply

# Call a tool
$ mcpjungle invoke calculator/multiply --input '{"a": 100, "b": 50}'

```

![Call a tool via MCPJungle Proxy MCP server](./assets/tool-call.png)

> [!NOTE]
> A tool in MCPJungle must be referred to by its canonical name which follows the pattern `<mcp-server-name>/<tool-name>`.
>
> eg- If you register a MCP server `github` which provides a tool called `git_commit`, you can invoke it in MCPJungle using the name `github/git_commit`.
> 
> Your AI Agent must also use this canonical name to call the tool via MCPJungle.


Finally, you can remove a MCP server from the registry:
```bash
$ mcpjungle deregister calculator
```

After running this, the registry will stop tracking this server and its tools will no longer be available to use.

### Cursor Configuration
Use the following configuration in Cursor to connect to MCPJungle:

```json
{
  "mcpServers": {
    "mcpjungle": {
      "url": "http://localhost:8080/mcp"
    }
  }
}
```

(Assuming that MCPJungle is running on `localhost:8080`)

### Authentication
MCPJungle currently supports authentication if your MCP Server accepts static tokens for auth.

This is useful when using SaaS-provided MCP Servers like HuggingFace, Stripe, etc. which require your API token for authentication.

You can supply your token while registering the MCP server:
```bash
# If you specify the `--bearer-token` flag, MCPJungle will add the `Authorization: Bearer <token>` header to all requests made to this MCP server.
$ mcpjungle register --name huggingface --description "HuggingFace MCP Server" --url https://huggingface.co/mcp --bearer-token <your-hf-api-token>
```

Support for other auth methods like Oauth is coming soon!

## Development

This section contains notes for maintainers and contributors of MCPJungle.

### Build for local testing
```bash
# Single binary for your current system
$ goreleaser build --single-target --clean --snapshot

# Test the full release assets (binaries, docker image) without publishing
goreleaser release --clean --snapshot --skip publish

# Binaries for all supported platforms
$ goreleaser release --snapshot --clean
```

### Create a new release
1. Create a Git Tag with the new version

```bash
git tag -a v0.1.0 -m "Release version 0.1.0"
git push origin v0.1.0
```

2. Release
```bash
# Make sure GPG is present on your system and you have a default key which is added to Github.

# set your github access token
export GITHUB_TOKEN="<your GH token>"

goreleaser release --clean
```

This will create a new release under Releases and also make it available via Homebrew.

## Modern Web Dashboard

MCPJungle now includes a modern web dashboard for managing MCP servers across different AI clients.

### Features
- **Visual Server Matrix**: Toggle servers on/off for different AI clients
- **Real-time Configuration**: Generate client-specific config files instantly  
- **Multi-Client Support**: Claude Desktop, Cursor, Windsurf, Cline, Claude Code CLI, Roo.Code, Kilo.Code
- **Modern UI**: Responsive design with Tailwind CSS
- **API Logging**: Detailed console logging for debugging

### Quick Start
```bash
# Start the full stack (API + Web Dashboard)
docker-compose up --build

# Access the dashboard
open http://localhost:3001
```

## Current Status & Roadmap

### ✅ Completed Features (v1.0)
- [x] **Modern Web Dashboard**: Responsive UI with server toggle matrix and real-time control
- [x] **Multi-Client Management**: Support for Claude Desktop, Cursor, Windsurf, Cline, Claude Code CLI, Roo.Code, Kilo.Code
- [x] **Granular Tool Control**: Expandable server rows with individual tool-level toggles per client
- [x] **Console Logging**: Detailed API call tracking with visual indicators for debugging
- [x] **Analytics Foundation**: Complete database schema for usage metrics, cost tracking, and performance monitoring
- [x] **Docker Integration**: Full-stack deployment with PostgreSQL, Go backend, and Next.js frontend
- [x] **Future-Ready Architecture**: Foundations for server discovery, health monitoring, and Anthropic SDK integration

### 🎯 Next Phase: Commercial Product Evolution

**Strategic Direction**: Transition from open-source tool to commercial SaaS platform (potential names: ToolFlow, AgentHub, ControlPlane AI)

### High Priority (Q1 2025)
- [ ] **Docker MCP Integration**: Leverage Docker's MCP Catalog and Toolkit (100+ verified servers)
- [ ] **Enhanced Server Discovery**: GitHub registry + Docker Hub catalog integration
- [ ] **Real-time Health Monitoring**: Functional heartbeat indicators with container status
- [ ] **Anthropic SDK Integration**: Live token usage tracking and cost calculations
- [ ] **Product Rebrand & Positioning**: Commercial product launch preparation
- [ ] **SaaS Infrastructure**: Multi-tenancy, authentication, billing system

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

### Revenue-Generating Features (Q2-Q3 2025)
- [ ] **Usage Analytics Dashboard**: Visual insights for cost optimization and ROI analysis
- [ ] **Enterprise Authentication**: SSO, RBAC, audit logs, compliance (SOC2, GDPR)
- [ ] **Advanced Alerting**: Cost thresholds, usage spikes, performance degradation
- [ ] **Business Intelligence**: Token consumption analysis, optimization recommendations
- [ ] **White-label Solutions**: Custom branding for enterprise customers
- [ ] **API Management**: Rate limiting, usage quotas, webhook integrations
- [ ] **Marketplace Ecosystem**: Plugin system for custom integrations

### Market Opportunity
**Target Market**: AI-forward development teams, consultants, and enterprises managing multiple AI clients
**Revenue Model**: Freemium SaaS ($19 Pro, $49 Team, $199 Enterprise per month)
**Competitive Advantage**: Only solution providing granular, client-specific MCP tool management with cost analytics
**Market Size**: $2B+ AI developer tools market growing 40% annually

### Integration Partnerships
- [ ] **Docker MCP Ecosystem**: Official integration with Docker's MCP Catalog
- [ ] **AI Client Partnerships**: Direct integrations with Anthropic, Cursor, Windsurf teams
- [ ] **Enterprise Platforms**: Slack, Microsoft Teams, Datadog, New Relic integrations

### Performance & UX
- [ ] **Real-time Updates**: WebSocket integration for live status updates
- [ ] **Keyboard Shortcuts**: Quick actions and navigation
- [ ] **Search & Filtering**: Advanced server and tool discovery
- [ ] **Drag & Drop**: Intuitive server assignment to clients

---

**Created by**: Jacques Wainwright  
**Support**: https://buymeacoffee.com/jlwainwright
