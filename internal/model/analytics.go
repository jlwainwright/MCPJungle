package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UsageMetric represents token usage and cost tracking
type UsageMetric struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Model        string    `json:"model" gorm:"not null;index"`
	ClientType   string    `json:"client_type" gorm:"not null;index"`
	InputTokens  int       `json:"input_tokens" gorm:"not null"`
	OutputTokens int       `json:"output_tokens" gorm:"not null"`
	TotalTokens  int       `json:"total_tokens" gorm:"not null"`
	Cost         float64   `json:"cost" gorm:"type:decimal(10,6)"`
	SessionID    string    `json:"session_id" gorm:"index"`
	UserID       *string   `json:"user_id,omitempty" gorm:"index"`
	Timestamp    time.Time `json:"timestamp" gorm:"not null;index"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (u *UsageMetric) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}

// ToolCall represents individual tool invocations with metrics
type ToolCall struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	ToolName     string    `json:"tool_name" gorm:"not null;index"`
	ServerName   string    `json:"server_name" gorm:"not null;index"`
	Model        string    `json:"model" gorm:"not null;index"`
	ClientType   string    `json:"client_type" gorm:"not null;index"`
	InputTokens  int       `json:"input_tokens"`
	OutputTokens int       `json:"output_tokens"`
	ResponseTime int       `json:"response_time"` // milliseconds
	Success      bool      `json:"success" gorm:"not null;index"`
	Error        *string   `json:"error,omitempty"`
	SessionID    string    `json:"session_id" gorm:"index"`
	UserID       *string   `json:"user_id,omitempty" gorm:"index"`
	Timestamp    time.Time `json:"timestamp" gorm:"not null;index"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// Relations
	UsageMetricID *uuid.UUID   `json:"usage_metric_id,omitempty" gorm:"type:uuid;index"`
	UsageMetric   *UsageMetric `json:"usage_metric,omitempty" gorm:"foreignKey:UsageMetricID"`
}

func (t *ToolCall) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return nil
}

// ServerMetric represents server-level performance metrics
type ServerMetric struct {
	ID                uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	ServerName        string    `json:"server_name" gorm:"not null;index"`
	TotalCalls        int       `json:"total_calls" gorm:"default:0"`
	SuccessfulCalls   int       `json:"successful_calls" gorm:"default:0"`
	FailedCalls       int       `json:"failed_calls" gorm:"default:0"`
	AvgResponseTime   float64   `json:"avg_response_time"` // milliseconds
	TotalTokens       int       `json:"total_tokens" gorm:"default:0"`
	TotalCost         float64   `json:"total_cost" gorm:"type:decimal(10,6);default:0"`
	LastCallTime      *time.Time `json:"last_call_time"`
	Date              time.Time  `json:"date" gorm:"not null;index"` // Daily aggregation
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

func (s *ServerMetric) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return nil
}

// ModelMetric represents model-level usage statistics
type ModelMetric struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Model            string    `json:"model" gorm:"not null;index"`
	ClientType       string    `json:"client_type" gorm:"not null;index"`
	TotalCalls       int       `json:"total_calls" gorm:"default:0"`
	TotalInputTokens int       `json:"total_input_tokens" gorm:"default:0"`
	TotalOutputTokens int      `json:"total_output_tokens" gorm:"default:0"`
	TotalCost        float64   `json:"total_cost" gorm:"type:decimal(10,6);default:0"`
	AvgResponseTime  float64   `json:"avg_response_time"`
	PopularTools     string    `json:"popular_tools" gorm:"type:text"` // JSON array of tool names
	Date             time.Time `json:"date" gorm:"not null;index"`     // Daily aggregation
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (m *ModelMetric) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return nil
}

// ClientMetric represents client-level usage patterns
type ClientMetric struct {
	ID                uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	ClientType        string    `json:"client_type" gorm:"not null;index"`
	TotalSessions     int       `json:"total_sessions" gorm:"default:0"`
	TotalToolCalls    int       `json:"total_tool_calls" gorm:"default:0"`
	TotalTokens       int       `json:"total_tokens" gorm:"default:0"`
	TotalCost         float64   `json:"total_cost" gorm:"type:decimal(10,6);default:0"`
	UniqueServers     int       `json:"unique_servers" gorm:"default:0"`
	UniqueTools       int       `json:"unique_tools" gorm:"default:0"`
	AvgSessionLength  float64   `json:"avg_session_length"` // minutes
	PreferredModels   string    `json:"preferred_models" gorm:"type:text"` // JSON array
	Date              time.Time `json:"date" gorm:"not null;index"`        // Daily aggregation
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (c *ClientMetric) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return nil
}

// CostSummary represents aggregated cost information for billing
type CostSummary struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Period       string    `json:"period" gorm:"not null;index"` // daily, weekly, monthly
	StartDate    time.Time `json:"start_date" gorm:"not null;index"`
	EndDate      time.Time `json:"end_date" gorm:"not null;index"`
	TotalCost    float64   `json:"total_cost" gorm:"type:decimal(10,6)"`
	TotalTokens  int       `json:"total_tokens"`
	TotalCalls   int       `json:"total_calls"`
	ModelCosts   string    `json:"model_costs" gorm:"type:text"`   // JSON breakdown by model
	ClientCosts  string    `json:"client_costs" gorm:"type:text"`  // JSON breakdown by client
	ServerCosts  string    `json:"server_costs" gorm:"type:text"`  // JSON breakdown by server
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (c *CostSummary) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return nil
}

// Alert represents usage alerts and notifications
type Alert struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Type        string    `json:"type" gorm:"not null;index"` // cost_threshold, usage_spike, error_rate
	Title       string    `json:"title" gorm:"not null"`
	Message     string    `json:"message" gorm:"not null"`
	Severity    string    `json:"severity" gorm:"not null;index"` // low, medium, high, critical
	Threshold   *float64  `json:"threshold,omitempty"`
	CurrentValue *float64  `json:"current_value,omitempty"`
	ResourceType string    `json:"resource_type"` // model, server, client, global
	ResourceID   *string   `json:"resource_id,omitempty"`
	Acknowledged bool      `json:"acknowledged" gorm:"default:false"`
	Resolved     bool      `json:"resolved" gorm:"default:false"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (a *Alert) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return nil
}