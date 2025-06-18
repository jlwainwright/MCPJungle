package api

import (
	"fmt"
	"github.com/duaraghav8/mcpjungle/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/mark3labs/mcp-go/server"
)

const V0PathPrefix = "/api/v0"

type Server struct {
	port           string
	router         *gin.Engine
	mcpProxyServer *server.MCPServer
	mcpService     *service.MCPService
	clientService  *service.ClientService
}

// NewServer initializes a new Gin server for MCPJungle registry and MCP proxy
func NewServer(port string, mcpProxyServer *server.MCPServer, mcpService *service.MCPService, clientService *service.ClientService) (*Server, error) {
	r, err := newRouter(mcpProxyServer, mcpService, clientService)
	if err != nil {
		return nil, err
	}
	s := &Server{
		port:           port,
		router:         r,
		mcpProxyServer: mcpProxyServer,
		mcpService:     mcpService,
		clientService:  clientService,
	}
	return s, nil
}

// Start runs the Gin server (blocking call)
func (s *Server) Start() error {
	if err := s.router.Run(":" + s.port); err != nil {
		return fmt.Errorf("failed to run the server: %w", err)
	}
	return nil
}

// newRouter sets up the Gin router with the MCP proxy server and API endpoints.
func newRouter(mcpProxyServer *server.MCPServer, mcpService *service.MCPService, clientService *service.ClientService) (*gin.Engine, error) {
	r := gin.Default()

	// Enable CORS for web interface
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	r.GET(
		"/health",
		func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		},
	)

	// Set up the MCP proxy server on /mcp
	streamableHttpServer := server.NewStreamableHTTPServer(mcpProxyServer)
	r.Any("/mcp", gin.WrapH(streamableHttpServer))

	// Setup API endpoints
	apiV0 := r.Group(V0PathPrefix)
	{
		apiV0.POST("/servers", registerServerHandler(mcpService))
		apiV0.DELETE("/servers/:name", deregisterServerHandler(mcpService))
		apiV0.GET("/servers", listServersHandler(mcpService))
		apiV0.GET("/tools", listToolsHandler(mcpService))
		apiV0.POST("/tools/invoke", invokeToolHandler(mcpService))
		apiV0.GET("/tool", getToolHandler(mcpService))
		
		// Client management endpoints
		apiV0.GET("/clients", listClientsGinHandler(clientService))
		apiV0.GET("/clients/:clientType/servers", getClientServersGinHandler(clientService))
		apiV0.POST("/clients/:clientType/servers/:serverId/toggle", toggleServerForClientGinHandler(clientService))
		apiV0.GET("/clients/:clientType/config", generateClientConfigGinHandler(clientService))
		apiV0.GET("/client-server-matrix", getClientServerMatrixGinHandler(clientService))
	}

	return r, nil
}
