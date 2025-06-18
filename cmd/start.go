package cmd

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"github.com/duaraghav8/mcpjungle/internal/api"
	"github.com/duaraghav8/mcpjungle/internal/db"
	"github.com/duaraghav8/mcpjungle/internal/migrations"
	"github.com/duaraghav8/mcpjungle/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

const (
	BindPortEnvVar  = "PORT"
	BindPortDefault = "8080"
)

var startServerCmdBindPort string

var startServerCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the MCP registry server",
	RunE:  runStartServer,
}

func init() {
	startServerCmd.Flags().StringVar(
		&startServerCmdBindPort,
		"port",
		"",
		fmt.Sprintf("port to bind the server to (overrides env var %s)", BindPortEnvVar),
	)
	rootCmd.AddCommand(startServerCmd)
}

func runStartServer(cmd *cobra.Command, args []string) error {
	_ = godotenv.Load()

	// configure debug level and logging
	configureDebugLevel()

	// connect to the DB and run migrations
	dsn := os.Getenv("DATABASE_URL")
	dbConn, err := db.NewDBConnection(dsn)
	if err != nil {
		return err
	}
	if err := migrations.Migrate(dbConn); err != nil {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	// determine the port to bind the server to
	port := startServerCmdBindPort
	if port == "" {
		port = os.Getenv(BindPortEnvVar)
	}
	if port == "" {
		port = BindPortDefault
	}

	// create the MCP proxy server
	mcpProxyServer := server.NewMCPServer(
		"MCPJungle Proxy MCP Server",
		"0.0.1",
		server.WithToolCapabilities(true),
	)

	mcpService, err := service.NewMCPService(dbConn, mcpProxyServer)
	if err != nil {
		return fmt.Errorf("failed to create MCP service: %v", err)
	}

	// create the client service
	clientService := service.NewClientService(dbConn)
	
	// initialize default clients
	if err := clientService.InitializeDefaultClients(); err != nil {
		return fmt.Errorf("failed to initialize default clients: %v", err)
	}

	// initialize example servers
	if err := mcpService.InitializeExampleServers(); err != nil {
		return fmt.Errorf("failed to initialize example servers: %v", err)
	}

	// create the health service
	healthService := service.NewHealthService(dbConn)

	// create the API server
	s, err := api.NewServer(port, mcpProxyServer, mcpService, clientService, healthService)
	if err != nil {
		return fmt.Errorf("failed to create server: %v", err)
	}

	fmt.Printf("MCPJungle server listening on :%s", port)
	if err := s.Start(); err != nil {
		return fmt.Errorf("failed to run the server: %v", err)
	}

	return nil
}

// configureDebugLevel sets up logging and debug level based on environment variables
func configureDebugLevel() {
	debugLevel := strings.ToLower(os.Getenv("DEBUG_LEVEL"))
	if debugLevel == "" {
		debugLevel = "info"
	}

	switch debugLevel {
	case "debug":
		gin.SetMode(gin.DebugMode)
		log.Println("[DEBUG] Debug mode enabled")
	case "info":
		gin.SetMode(gin.ReleaseMode)
		log.Println("[INFO] Info mode enabled")
	case "warn", "warning":
		gin.SetMode(gin.ReleaseMode)
		log.Println("[WARN] Warning mode enabled")
	case "error":
		gin.SetMode(gin.ReleaseMode)
		log.Println("[ERROR] Error mode enabled")
	default:
		gin.SetMode(gin.ReleaseMode)
		log.Printf("[WARN] Unknown debug level '%s', defaulting to info mode", debugLevel)
	}

	// Enable memory profiling if requested
	enablePprof := strings.ToLower(os.Getenv("ENABLE_PPROF"))
	if enablePprof == "true" {
		pprofPort := os.Getenv("PPROF_PORT")
		if pprofPort == "" {
			pprofPort = "6060"
		}
		log.Printf("[DEBUG] Memory profiling enabled on port %s", pprofPort)
		log.Println("[DEBUG] Access profiling at http://localhost:" + pprofPort + "/debug/pprof/")
		
		// Start pprof server in a goroutine
		go func() {
			log.Println(http.ListenAndServe("localhost:"+pprofPort, nil))
		}()
	}
}
