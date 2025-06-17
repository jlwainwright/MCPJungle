package service

import (
	"context"
	"fmt"
	"github.com/duaraghav8/mcpjungle/internal/model"
	"gorm.io/gorm"
)

// RegisterMcpServer registers a new MCP server in the database.
// It also registers all the Tools provided by the server.
// Tool registration is on best-effort basis and does not fail the server registration.
// Registered tools are also added to the MCP proxy server.
func (m *MCPService) RegisterMcpServer(ctx context.Context, s *model.McpServer) error {
	if err := validateServerName(s.Name); err != nil {
		return err
	}

	// TODO: validate the URL to ensure it is a valid HTTP/HTTPS URL (streamable http compliant)

	// test that the server is reachable and is MCP-compliant
	c, err := createMcpServerConn(ctx, s)
	if err != nil {
		return fmt.Errorf("failed to connect to MCP server %s: %w", s.Name, err)
	}
	defer c.Close()

	// register the server in the DB
	if err := m.db.Create(s).Error; err != nil {
		return fmt.Errorf("failed to register mcp server: %w", err)
	}

	if err = m.registerServerTools(ctx, s, c); err != nil {
		return fmt.Errorf("failed to register tools for MCP server %s: %w", s.Name, err)
	}
	return nil
}

// DeregisterMcpServer deregisters an MCP server from the database.
// It also deregisters all the tools registered by the server.
// If even a singe tool fails to deregister, the server deregistration fails.
// A deregistered tool is also removed from the MCP proxy server.
func (m *MCPService) DeregisterMcpServer(name string) error {
	s, err := m.GetMcpServer(name)
	if err != nil {
		return fmt.Errorf("failed to get MCP server %s from DB: %w", name, err)
	}
	if err := m.deregisterServerTools(s); err != nil {
		return fmt.Errorf(
			"failed to deregister tools for server %s, cannot proceed with server deregistration: %w",
			name,
			err,
		)
	}
	if err := m.db.Delete(s).Error; err != nil {
		return fmt.Errorf("failed to deregister server %s: %w", name, err)
	}
	return nil
}

// ListMcpServers returns all registered MCP servers.
func (m *MCPService) ListMcpServers() ([]model.McpServer, error) {
	var servers []model.McpServer
	if err := m.db.Find(&servers).Error; err != nil {
		return nil, err
	}
	return servers, nil
}

// GetMcpServer fetches a server from the database by name.
func (m *MCPService) GetMcpServer(name string) (*model.McpServer, error) {
	var serverModel model.McpServer
	if err := m.db.Where("name = ?", name).First(&serverModel).Error; err != nil {
		return nil, err
	}
	return &serverModel, nil
}

// InitializeExampleServers creates example MCP servers for demo purposes
func (m *MCPService) InitializeExampleServers() error {
	exampleServers := []model.McpServer{
		{
			Name:        "github",
			Description: "GitHub integration tools for repository management",
			URL:         "https://api.github.com/mcp",
			BearerToken: "", // No token for demo
		},
		{
			Name:        "calculator", 
			Description: "Basic mathematical operations and calculations",
			URL:         "http://localhost:8000/mcp",
			BearerToken: "",
		},
		{
			Name:        "weather",
			Description: "Weather information and forecasting tools",
			URL:         "https://api.weather.com/mcp",
			BearerToken: "",
		},
		{
			Name:        "database",
			Description: "Database query and management utilities",
			URL:         "http://localhost:5432/mcp", 
			BearerToken: "",
		},
	}

	for _, server := range exampleServers {
		var existingServer model.McpServer
		err := m.db.Where("name = ?", server.Name).First(&existingServer).Error
		if err == gorm.ErrRecordNotFound {
			// Server doesn't exist, create it
			err = m.db.Create(&server).Error
			if err != nil {
				return fmt.Errorf("failed to create example server %s: %w", server.Name, err)
			}
			
			// Add some example tools for each server (without actually connecting)
			tools := getExampleToolsForServer(server.Name)
			for _, tool := range tools {
				tool.ServerID = server.ID
				err = m.db.Create(&tool).Error
				if err != nil {
					// Log error but continue
					fmt.Printf("Failed to create tool %s for server %s: %v\n", tool.Name, server.Name, err)
				}
			}
		}
	}

	return nil
}

// getExampleToolsForServer returns mock tools for demo servers
func getExampleToolsForServer(serverName string) []model.Tool {
	switch serverName {
	case "github":
		return []model.Tool{
			{Name: "create_repo", Description: "Create a new GitHub repository", InputSchema: []byte(`{"type":"object","properties":{"name":{"type":"string"},"description":{"type":"string"}},"required":["name"]}`)},
			{Name: "list_repos", Description: "List user repositories", InputSchema: []byte(`{"type":"object","properties":{"username":{"type":"string"}},"required":["username"]}`)},
			{Name: "create_issue", Description: "Create a new issue", InputSchema: []byte(`{"type":"object","properties":{"title":{"type":"string"},"body":{"type":"string"}},"required":["title"]}`)},
			{Name: "get_commits", Description: "Get commit history", InputSchema: []byte(`{"type":"object","properties":{"repo":{"type":"string"},"limit":{"type":"number"}},"required":["repo"]}`)},
		}
	case "calculator":
		return []model.Tool{
			{Name: "add", Description: "Add two numbers", InputSchema: []byte(`{"type":"object","properties":{"a":{"type":"number"},"b":{"type":"number"}},"required":["a","b"]}`)},
			{Name: "subtract", Description: "Subtract two numbers", InputSchema: []byte(`{"type":"object","properties":{"a":{"type":"number"},"b":{"type":"number"}},"required":["a","b"]}`)},
			{Name: "multiply", Description: "Multiply two numbers", InputSchema: []byte(`{"type":"object","properties":{"a":{"type":"number"},"b":{"type":"number"}},"required":["a","b"]}`)},
			{Name: "divide", Description: "Divide two numbers", InputSchema: []byte(`{"type":"object","properties":{"a":{"type":"number"},"b":{"type":"number"}},"required":["a","b"]}`)},
		}
	case "weather":
		return []model.Tool{
			{Name: "get_weather", Description: "Get current weather conditions", InputSchema: []byte(`{"type":"object","properties":{"location":{"type":"string"}},"required":["location"]}`)},
			{Name: "get_forecast", Description: "Get weather forecast", InputSchema: []byte(`{"type":"object","properties":{"location":{"type":"string"},"days":{"type":"number"}},"required":["location"]}`)},
			{Name: "get_alerts", Description: "Get weather alerts", InputSchema: []byte(`{"type":"object","properties":{"location":{"type":"string"}},"required":["location"]}`)},
		}
	case "database":
		return []model.Tool{
			{Name: "query", Description: "Execute SQL query", InputSchema: []byte(`{"type":"object","properties":{"sql":{"type":"string"}},"required":["sql"]}`)},
			{Name: "insert", Description: "Insert data into table", InputSchema: []byte(`{"type":"object","properties":{"table":{"type":"string"},"data":{"type":"object"}},"required":["table","data"]}`)},
			{Name: "update", Description: "Update table records", InputSchema: []byte(`{"type":"object","properties":{"table":{"type":"string"},"data":{"type":"object"},"where":{"type":"object"}},"required":["table","data"]}`)},
			{Name: "delete", Description: "Delete records from table", InputSchema: []byte(`{"type":"object","properties":{"table":{"type":"string"},"where":{"type":"object"}},"required":["table","where"]}`)},
		}
	default:
		return []model.Tool{}
	}
}
