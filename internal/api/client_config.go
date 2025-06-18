package api

import (
	"encoding/json"
	"net/http"

	"github.com/duaraghav8/mcpjungle/internal/model"
	"github.com/duaraghav8/mcpjungle/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

type ClientConfigAPI struct {
	clientService *service.ClientService
}

func NewClientConfigAPI(clientService *service.ClientService) *ClientConfigAPI {
	return &ClientConfigAPI{
		clientService: clientService,
	}
}

func (api *ClientConfigAPI) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/clients", api.ListClients).Methods("GET")
	router.HandleFunc("/clients/{clientType}/servers", api.GetClientServers).Methods("GET")
	router.HandleFunc("/clients/{clientType}/servers/{serverId}/toggle", api.ToggleServerForClient).Methods("POST")
	router.HandleFunc("/clients/{clientType}/config", api.GenerateClientConfig).Methods("GET")
	router.HandleFunc("/client-server-matrix", api.GetClientServerMatrix).Methods("GET")
}

func (api *ClientConfigAPI) ListClients(w http.ResponseWriter, r *http.Request) {
	clients, err := api.clientService.ListClients()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)
}

func (api *ClientConfigAPI) GetClientServers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientType := model.ClientType(vars["clientType"])

	servers, err := api.clientService.GetClientServers(clientType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(servers)
}

func (api *ClientConfigAPI) ToggleServerForClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientType := model.ClientType(vars["clientType"])
	serverID := vars["serverId"]

	err := api.clientService.ToggleServerForClient(clientType, serverID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func (api *ClientConfigAPI) GenerateClientConfig(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientType := model.ClientType(vars["clientType"])

	config, err := api.clientService.GenerateClientConfig(clientType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config)
}

func (api *ClientConfigAPI) GetClientServerMatrix(w http.ResponseWriter, r *http.Request) {
	matrix, err := api.clientService.GetClientServerMatrix()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matrix)
}

// Gin handler wrappers
func listClientsGinHandler(clientService *service.ClientService) gin.HandlerFunc {
	return func(c *gin.Context) {
		clients, err := clientService.ListClients()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, clients)
	}
}

func getClientServersGinHandler(clientService *service.ClientService) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientType := model.ClientType(c.Param("clientType"))
		servers, err := clientService.GetClientServers(clientType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, servers)
	}
}

func toggleServerForClientGinHandler(clientService *service.ClientService) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientType := model.ClientType(c.Param("clientType"))
		serverID := c.Param("serverId")
		
		err := clientService.ToggleServerForClient(clientType, serverID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}

func generateClientConfigGinHandler(clientService *service.ClientService) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientType := model.ClientType(c.Param("clientType"))
		config, err := clientService.GenerateClientConfig(clientType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, config)
	}
}

func getClientServerMatrixGinHandler(clientService *service.ClientService) gin.HandlerFunc {
	return func(c *gin.Context) {
		matrix, err := clientService.GetClientServerMatrix()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, matrix)
	}
}