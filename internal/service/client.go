package service

import (
	"fmt"

	"github.com/duaraghav8/mcpjungle/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientService struct {
	db *gorm.DB
}

func NewClientService(db *gorm.DB) *ClientService {
	return &ClientService{db: db}
}

type ClientServerMatrix map[string]map[string]bool

func (s *ClientService) ListClients() ([]model.ClientConfig, error) {
	var clients []model.ClientConfig
	err := s.db.Find(&clients).Error
	return clients, err
}

func (s *ClientService) GetClientServers(clientType model.ClientType) ([]model.ClientServerMapping, error) {
	var mappings []model.ClientServerMapping
	err := s.db.Preload("McpServer").Where("client_type = ?", clientType).Find(&mappings).Error
	return mappings, err
}

func (s *ClientService) ToggleServerForClient(clientType model.ClientType, serverID string) error {
	serverUUID, err := uuid.Parse(serverID)
	if err != nil {
		return fmt.Errorf("invalid server ID: %v", err)
	}

	var mapping model.ClientServerMapping
	err = s.db.Where("client_type = ? AND mcp_server_id = ?", clientType, serverUUID).First(&mapping).Error
	
	if err == gorm.ErrRecordNotFound {
		// Create new mapping if it doesn't exist
		mapping = model.ClientServerMapping{
			ClientType:  clientType,
			McpServerID: serverUUID,
			Enabled:     true,
		}
		return s.db.Create(&mapping).Error
	} else if err != nil {
		return err
	}

	// Toggle existing mapping
	mapping.Enabled = !mapping.Enabled
	return s.db.Save(&mapping).Error
}

func (s *ClientService) GenerateClientConfig(clientType model.ClientType) (map[string]interface{}, error) {
	var mappings []model.ClientServerMapping
	err := s.db.Preload("McpServer").Where("client_type = ? AND enabled = ?", clientType, true).Find(&mappings).Error
	if err != nil {
		return nil, err
	}

	mcpServers := make(map[string]interface{})
	for _, mapping := range mappings {
		mcpServers[mapping.McpServer.Name] = map[string]interface{}{
			"url": fmt.Sprintf("http://localhost:8080/mcp/%s", mapping.McpServer.Name),
		}
	}

	config := map[string]interface{}{
		"mcpServers": mcpServers,
	}

	return config, nil
}

func (s *ClientService) GetClientServerMatrix() (ClientServerMatrix, error) {
	var servers []model.McpServer
	err := s.db.Find(&servers).Error
	if err != nil {
		return nil, err
	}

	var mappings []model.ClientServerMapping
	err = s.db.Find(&mappings).Error
	if err != nil {
		return nil, err
	}

	// Create mapping lookup
	mappingLookup := make(map[string]map[string]bool)
	for _, mapping := range mappings {
		if mappingLookup[string(mapping.ClientType)] == nil {
			mappingLookup[string(mapping.ClientType)] = make(map[string]bool)
		}
		mappingLookup[string(mapping.ClientType)][mapping.McpServerID.String()] = mapping.Enabled
	}

	// Build matrix
	matrix := make(ClientServerMatrix)
	clientTypes := []model.ClientType{
		model.ClientTypeClaude,
		model.ClientTypeCursor,
		model.ClientTypeWindsurf,
		model.ClientTypeCline,
		model.ClientTypeClaudeCode,
		model.ClientTypeRooCode,
		model.ClientTypeKiloCode,
	}

	for _, clientType := range clientTypes {
		matrix[string(clientType)] = make(map[string]bool)
		for _, server := range servers {
			enabled := false
			if clientMappings, exists := mappingLookup[string(clientType)]; exists {
				enabled = clientMappings[server.ID.String()]
			}
			matrix[string(clientType)][server.ID.String()] = enabled
		}
	}

	return matrix, nil
}

func (s *ClientService) InitializeDefaultClients() error {
	defaultClients := []model.ClientConfig{
		{
			ClientType:  model.ClientTypeClaude,
			DisplayName: "Claude Desktop",
			ConfigPath:  "~/Library/Application Support/Claude/claude_desktop_config.json",
			Icon:        "🧠",
			Enabled:     true,
		},
		{
			ClientType:  model.ClientTypeCursor,
			DisplayName: "Cursor",
			ConfigPath:  "~/.cursor/mcp_config.json",
			Icon:        "↗️",
			Enabled:     true,
		},
		{
			ClientType:  model.ClientTypeWindsurf,
			DisplayName: "Windsurf",
			ConfigPath:  "~/.windsurf/mcp_config.json",
			Icon:        "🏄",
			Enabled:     true,
		},
		{
			ClientType:  model.ClientTypeCline,
			DisplayName: "Cline",
			ConfigPath:  "~/.cline/mcp_config.json",
			Icon:        "📋",
			Enabled:     true,
		},
		{
			ClientType:  model.ClientTypeClaudeCode,
			DisplayName: "Claude Code CLI",
			ConfigPath:  "~/.config/claude-code/mcp_config.json",
			Icon:        "💻",
			Enabled:     true,
		},
		{
			ClientType:  model.ClientTypeRooCode,
			DisplayName: "Roo.Code",
			ConfigPath:  "~/.roocode/mcp_config.json",
			Icon:        "🦘",
			Enabled:     true,
		},
		{
			ClientType:  model.ClientTypeKiloCode,
			DisplayName: "Kilo.Code",
			ConfigPath:  "~/.kilocode/mcp_config.json",
			Icon:        "⚡",
			Enabled:     true,
		},
	}

	for _, client := range defaultClients {
		var existingClient model.ClientConfig
		err := s.db.Where("client_type = ?", client.ClientType).First(&existingClient).Error
		if err == gorm.ErrRecordNotFound {
			err = s.db.Create(&client).Error
			if err != nil {
				return err
			}
		}
	}

	return nil
}