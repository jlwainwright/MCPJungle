package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientType string

const (
	ClientTypeClaude   ClientType = "claude"
	ClientTypeCursor   ClientType = "cursor"
	ClientTypeWindsurf ClientType = "windsurf"
	ClientTypeCline    ClientType = "cline"
)

type ClientConfig struct {
	ID          uuid.UUID  `json:"-" gorm:"type:uuid;primaryKey"`
	ClientType  ClientType `json:"client_type" gorm:"not null"`
	DisplayName string     `json:"display_name" gorm:"not null"`
	ConfigPath  string     `json:"config_path"`
	Icon        string     `json:"icon"`
	Enabled     bool       `json:"enabled" gorm:"default:true"`
}

func (c *ClientConfig) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return nil
}

// ClientServerMapping represents which servers are enabled for which clients
type ClientServerMapping struct {
	ID           uuid.UUID `json:"-" gorm:"type:uuid;primaryKey"`
	ClientType   ClientType `json:"client_type" gorm:"not null"`
	McpServerID  uuid.UUID `json:"mcp_server_id" gorm:"type:uuid;not null"`
	Enabled      bool      `json:"enabled" gorm:"default:false"`
	
	// Relations
	McpServer    McpServer `json:"mcp_server" gorm:"foreignKey:McpServerID"`
}

func (csm *ClientServerMapping) BeforeCreate(tx *gorm.DB) (err error) {
	csm.ID = uuid.New()
	return nil
}