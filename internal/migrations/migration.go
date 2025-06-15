package migrations

import (
	"fmt"
	"github.com/duaraghav8/mcpjungle/internal/model"
	"gorm.io/gorm"
)

// Migrate performs the database migration for the application.
func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.McpServer{}); err != nil {
		return fmt.Errorf("auto‑migration failed for McpServer model: %v", err)
	}
	if err := db.AutoMigrate(&model.Tool{}); err != nil {
		return fmt.Errorf("auto‑migration failed for Tool model: %v", err)
	}
	if err := db.AutoMigrate(&model.ClientConfig{}); err != nil {
		return fmt.Errorf("auto‑migration failed for ClientConfig model: %v", err)
	}
	if err := db.AutoMigrate(&model.ClientServerMapping{}); err != nil {
		return fmt.Errorf("auto‑migration failed for ClientServerMapping model: %v", err)
	}
	if err := db.AutoMigrate(&model.UsageMetric{}); err != nil {
		return fmt.Errorf("auto‑migration failed for UsageMetric model: %v", err)
	}
	if err := db.AutoMigrate(&model.ToolCall{}); err != nil {
		return fmt.Errorf("auto‑migration failed for ToolCall model: %v", err)
	}
	if err := db.AutoMigrate(&model.ServerMetric{}); err != nil {
		return fmt.Errorf("auto‑migration failed for ServerMetric model: %v", err)
	}
	if err := db.AutoMigrate(&model.ModelMetric{}); err != nil {
		return fmt.Errorf("auto‑migration failed for ModelMetric model: %v", err)
	}
	if err := db.AutoMigrate(&model.ClientMetric{}); err != nil {
		return fmt.Errorf("auto‑migration failed for ClientMetric model: %v", err)
	}
	if err := db.AutoMigrate(&model.CostSummary{}); err != nil {
		return fmt.Errorf("auto‑migration failed for CostSummary model: %v", err)
	}
	if err := db.AutoMigrate(&model.Alert{}); err != nil {
		return fmt.Errorf("auto‑migration failed for Alert model: %v", err)
	}
	return nil
}
