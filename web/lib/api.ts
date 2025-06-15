const API_BASE_URL = process.env.NODE_ENV === 'development' 
  ? 'http://localhost:8080/api/v0'
  : '/api/v0'

export interface MCPServer {
  id: string
  name: string
  description: string
  url: string
  status: 'active' | 'inactive'
  tools: string[]
}

export interface ClientConfig {
  name: string
  displayName: string
  icon: string
  configPath: string
  enabled: boolean
}

export interface ClientServerMatrix {
  [clientName: string]: {
    [serverId: string]: boolean
  }
}

export const api = {
  async getServers(): Promise<MCPServer[]> {
    const response = await fetch(`${API_BASE_URL}/servers`)
    if (!response.ok) throw new Error('Failed to fetch servers')
    return response.json()
  },

  async getClients(): Promise<ClientConfig[]> {
    const response = await fetch(`${API_BASE_URL}/clients`)
    if (!response.ok) throw new Error('Failed to fetch clients')
    return response.json()
  },

  async getClientServerMatrix(): Promise<ClientServerMatrix> {
    const response = await fetch(`${API_BASE_URL}/client-server-matrix`)
    if (!response.ok) throw new Error('Failed to fetch client server matrix')
    return response.json()
  },

  async toggleServerForClient(clientType: string, serverId: string): Promise<void> {
    const response = await fetch(`${API_BASE_URL}/clients/${clientType}/servers/${serverId}/toggle`, {
      method: 'POST'
    })
    if (!response.ok) throw new Error('Failed to toggle server for client')
  },

  async generateClientConfig(clientType: string): Promise<object> {
    const response = await fetch(`${API_BASE_URL}/clients/${clientType}/config`)
    if (!response.ok) throw new Error('Failed to generate client config')
    return response.json()
  }
}