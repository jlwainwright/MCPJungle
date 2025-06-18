package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/duaraghav8/mcpjungle/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnalyticsService struct {
	db *gorm.DB
}

func NewAnalyticsService(db *gorm.DB) *AnalyticsService {
	return &AnalyticsService{db: db}
}

// RecordUsage records token usage and cost metrics
func (s *AnalyticsService) RecordUsage(usage model.UsageMetric) error {
	usage.TotalTokens = usage.InputTokens + usage.OutputTokens
	return s.db.Create(&usage).Error
}

// RecordToolCall records individual tool invocation metrics
func (s *AnalyticsService) RecordToolCall(toolCall model.ToolCall) error {
	return s.db.Create(&toolCall).Error
}

// GetUsageByTimeframe returns usage statistics for a given timeframe
func (s *AnalyticsService) GetUsageByTimeframe(timeframe string, clientType *string) (map[string]interface{}, error) {
	var startDate time.Time
	now := time.Now()

	switch timeframe {
	case "day":
		startDate = now.AddDate(0, 0, -1)
	case "week":
		startDate = now.AddDate(0, 0, -7)
	case "month":
		startDate = now.AddDate(0, -1, 0)
	case "year":
		startDate = now.AddDate(-1, 0, 0)
	default:
		return nil, fmt.Errorf("invalid timeframe: %s", timeframe)
	}

	query := s.db.Model(&model.UsageMetric{}).Where("timestamp >= ?", startDate)
	if clientType != nil {
		query = query.Where("client_type = ?", *clientType)
	}

	var totalTokens int64
	var totalCost float64
	var totalCalls int64

	// Get aggregated statistics
	err := query.Select("SUM(total_tokens) as total_tokens, SUM(cost) as total_cost, COUNT(*) as total_calls").
		Row().Scan(&totalTokens, &totalCost, &totalCalls)
	if err != nil {
		return nil, err
	}

	// Get top models
	var topModels []struct {
		Model      string  `json:"model"`
		CallCount  int     `json:"call_count"`
		TotalCost  float64 `json:"total_cost"`
		TotalTokens int    `json:"total_tokens"`
	}

	err = query.Select("model, COUNT(*) as call_count, SUM(cost) as total_cost, SUM(total_tokens) as total_tokens").
		Group("model").Order("call_count DESC").Limit(10).Scan(&topModels)
	if err != nil {
		return nil, err
	}

	// Get usage trends (daily breakdown)
	var trends []struct {
		Date        string  `json:"date"`
		TotalTokens int     `json:"total_tokens"`
		TotalCost   float64 `json:"total_cost"`
		CallCount   int     `json:"call_count"`
	}

	err = query.Select("DATE(timestamp) as date, SUM(total_tokens) as total_tokens, SUM(cost) as total_cost, COUNT(*) as call_count").
		Group("DATE(timestamp)").Order("date").Scan(&trends)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"summary": map[string]interface{}{
			"total_tokens": totalTokens,
			"total_cost":   totalCost,
			"total_calls":  totalCalls,
		},
		"top_models": topModels,
		"trends":     trends,
	}, nil
}

// GetToolUsageStats returns tool usage statistics
func (s *AnalyticsService) GetToolUsageStats(timeframe string) (map[string]interface{}, error) {
	var startDate time.Time
	now := time.Now()

	switch timeframe {
	case "day":
		startDate = now.AddDate(0, 0, -1)
	case "week":
		startDate = now.AddDate(0, 0, -7)
	case "month":
		startDate = now.AddDate(0, -1, 0)
	case "year":
		startDate = now.AddDate(-1, 0, 0)
	default:
		return nil, fmt.Errorf("invalid timeframe: %s", timeframe)
	}

	// Get top tools
	var topTools []struct {
		ToolName     string  `json:"tool_name"`
		ServerName   string  `json:"server_name"`
		CallCount    int     `json:"call_count"`
		SuccessRate  float64 `json:"success_rate"`
		AvgResponseTime float64 `json:"avg_response_time"`
	}

	err := s.db.Model(&model.ToolCall{}).
		Select("tool_name, server_name, COUNT(*) as call_count, AVG(CASE WHEN success THEN 1.0 ELSE 0.0 END) as success_rate, AVG(response_time) as avg_response_time").
		Where("timestamp >= ?", startDate).
		Group("tool_name, server_name").
		Order("call_count DESC").
		Limit(20).
		Scan(&topTools)
	if err != nil {
		return nil, err
	}

	// Get server performance
	var serverStats []struct {
		ServerName   string  `json:"server_name"`
		CallCount    int     `json:"call_count"`
		SuccessRate  float64 `json:"success_rate"`
		AvgResponseTime float64 `json:"avg_response_time"`
	}

	err = s.db.Model(&model.ToolCall{}).
		Select("server_name, COUNT(*) as call_count, AVG(CASE WHEN success THEN 1.0 ELSE 0.0 END) as success_rate, AVG(response_time) as avg_response_time").
		Where("timestamp >= ?", startDate).
		Group("server_name").
		Order("call_count DESC").
		Scan(&serverStats)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"top_tools":    topTools,
		"server_stats": serverStats,
	}, nil
}

// GetClientUsageBreakdown returns usage breakdown by client type
func (s *AnalyticsService) GetClientUsageBreakdown(timeframe string) ([]map[string]interface{}, error) {
	var startDate time.Time
	now := time.Now()

	switch timeframe {
	case "day":
		startDate = now.AddDate(0, 0, -1)
	case "week":
		startDate = now.AddDate(0, 0, -7)
	case "month":
		startDate = now.AddDate(0, -1, 0)
	case "year":
		startDate = now.AddDate(-1, 0, 0)
	default:
		return nil, fmt.Errorf("invalid timeframe: %s", timeframe)
	}

	var clientBreakdown []struct {
		ClientType  string  `json:"client_type"`
		TotalTokens int     `json:"total_tokens"`
		TotalCost   float64 `json:"total_cost"`
		CallCount   int     `json:"call_count"`
	}

	err := s.db.Model(&model.UsageMetric{}).
		Select("client_type, SUM(total_tokens) as total_tokens, SUM(cost) as total_cost, COUNT(*) as call_count").
		Where("timestamp >= ?", startDate).
		Group("client_type").
		Order("total_cost DESC").
		Scan(&clientBreakdown)
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, len(clientBreakdown))
	for i, cb := range clientBreakdown {
		result[i] = map[string]interface{}{
			"client_type":  cb.ClientType,
			"total_tokens": cb.TotalTokens,
			"total_cost":   cb.TotalCost,
			"call_count":   cb.CallCount,
		}
	}

	return result, nil
}

// GenerateCostSummary creates aggregated cost reports for billing
func (s *AnalyticsService) GenerateCostSummary(period string, startDate, endDate time.Time) (*model.CostSummary, error) {
	// Get total metrics
	var totalCost float64
	var totalTokens int64
	var totalCalls int64

	err := s.db.Model(&model.UsageMetric{}).
		Select("SUM(cost) as total_cost, SUM(total_tokens) as total_tokens, COUNT(*) as total_calls").
		Where("timestamp BETWEEN ? AND ?", startDate, endDate).
		Row().Scan(&totalCost, &totalTokens, &totalCalls)
	if err != nil {
		return nil, err
	}

	// Get model breakdown
	var modelBreakdown []struct {
		Model     string  `json:"model"`
		Cost      float64 `json:"cost"`
		Tokens    int     `json:"tokens"`
		CallCount int     `json:"call_count"`
	}

	err = s.db.Model(&model.UsageMetric{}).
		Select("model, SUM(cost) as cost, SUM(total_tokens) as tokens, COUNT(*) as call_count").
		Where("timestamp BETWEEN ? AND ?", startDate, endDate).
		Group("model").Scan(&modelBreakdown)
	if err != nil {
		return nil, err
	}

	modelCosts, _ := json.Marshal(modelBreakdown)

	// Get client breakdown
	clientBreakdown, err := s.GetClientUsageBreakdown(period)
	if err != nil {
		return nil, err
	}
	clientCosts, _ := json.Marshal(clientBreakdown)

	// Get server breakdown from tool calls
	var serverBreakdown []struct {
		ServerName string `json:"server_name"`
		CallCount  int    `json:"call_count"`
	}

	err = s.db.Model(&model.ToolCall{}).
		Select("server_name, COUNT(*) as call_count").
		Where("timestamp BETWEEN ? AND ?", startDate, endDate).
		Group("server_name").Scan(&serverBreakdown)
	if err != nil {
		return nil, err
	}
	serverCosts, _ := json.Marshal(serverBreakdown)

	summary := &model.CostSummary{
		Period:      period,
		StartDate:   startDate,
		EndDate:     endDate,
		TotalCost:   totalCost,
		TotalTokens: int(totalTokens),
		TotalCalls:  int(totalCalls),
		ModelCosts:  string(modelCosts),
		ClientCosts: string(clientCosts),
		ServerCosts: string(serverCosts),
	}

	err = s.db.Create(summary).Error
	return summary, err
}

// CreateAlert creates usage alerts and notifications
func (s *AnalyticsService) CreateAlert(alertType, title, message, severity string, threshold, currentValue *float64, resourceType, resourceID *string) error {
	alert := model.Alert{
		Type:         alertType,
		Title:        title,
		Message:      message,
		Severity:     severity,
		Threshold:    threshold,
		CurrentValue: currentValue,
		ResourceType: resourceType,
		ResourceID:   resourceID,
	}

	return s.db.Create(&alert).Error
}

// GetActiveAlerts returns unresolved alerts
func (s *AnalyticsService) GetActiveAlerts() ([]model.Alert, error) {
	var alerts []model.Alert
	err := s.db.Where("resolved = ?", false).Order("created_at DESC").Find(&alerts).Error
	return alerts, err
}

// CheckThresholds monitors usage and creates alerts when thresholds are exceeded
func (s *AnalyticsService) CheckThresholds() error {
	// TODO: Implement threshold monitoring
	// Check daily/monthly cost limits
	// Check token usage spikes
	// Check error rates
	// Create alerts when thresholds exceeded
	return nil
}

// ExportUsageReport generates usage reports for external systems
func (s *AnalyticsService) ExportUsageReport(startDate, endDate time.Time, format string) ([]byte, error) {
	// TODO: Implement report export
	// Support CSV, JSON, PDF formats
	// Include usage breakdowns, costs, trends
	// Return formatted report data
	return nil, fmt.Errorf("export functionality not implemented yet")
}