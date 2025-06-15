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
	return nil
}
