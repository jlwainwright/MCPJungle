const API_BASE_URL = process.env.NODE_ENV === 'development' 
  ? 'http://localhost:8080/api/v0'
  : '/api/v0'

// Console logging for API calls
const logApiCall = (method: string, url: string, data?: any) => {
  console.group(`🌳 MCPJungle API: ${method.toUpperCase()} ${url}`)
  console.log('🕐 Timestamp:', new Date().toISOString())
  if (data) {
    console.log('📤 Request Data:', data)
  }
  console.groupEnd()
}

const logApiResponse = (method: string, url: string, response: Response, data?: any) => {
  console.group(`🌳 MCPJungle API Response: ${method.toUpperCase()} ${url}`)
  console.log('🕐 Timestamp:', new Date().toISOString())
  console.log('📊 Status:', response.status, response.statusText)
  console.log('📍 URL:', response.url)
  if (data) {
    console.log('📥 Response Data:', data)
  }
  console.groupEnd()
}

const logApiError = (method: string, url: string, error: any) => {
  console.group(`🚨 MCPJungle API Error: ${method.toUpperCase()} ${url}`)
  console.log('🕐 Timestamp:', new Date().toISOString())
  console.error('❌ Error:', error)
  console.groupEnd()
}

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
    const url = `${API_BASE_URL}/servers`
    logApiCall('GET', url)
    try {
      const response = await fetch(url)
      const data = await response.json()
      logApiResponse('GET', url, response, data)
      if (!response.ok) throw new Error('Failed to fetch servers')
      return data
    } catch (error) {
      logApiError('GET', url, error)
      throw error
    }
  },

  async getClients(): Promise<ClientConfig[]> {
    const url = `${API_BASE_URL}/clients`
    logApiCall('GET', url)
    try {
      const response = await fetch(url)
      const data = await response.json()
      logApiResponse('GET', url, response, data)
      if (!response.ok) throw new Error('Failed to fetch clients')
      return data
    } catch (error) {
      logApiError('GET', url, error)
      throw error
    }
  },

  async getClientServerMatrix(): Promise<ClientServerMatrix> {
    const url = `${API_BASE_URL}/client-server-matrix`
    logApiCall('GET', url)
    try {
      const response = await fetch(url)
      const data = await response.json()
      logApiResponse('GET', url, response, data)
      if (!response.ok) throw new Error('Failed to fetch client server matrix')
      return data
    } catch (error) {
      logApiError('GET', url, error)
      throw error
    }
  },

  async toggleServerForClient(clientType: string, serverId: string): Promise<void> {
    const url = `${API_BASE_URL}/clients/${clientType}/servers/${serverId}/toggle`
    logApiCall('POST', url, { clientType, serverId })
    try {
      const response = await fetch(url, {
        method: 'POST'
      })
      const data = response.status !== 204 ? await response.json() : null
      logApiResponse('POST', url, response, data)
      if (!response.ok) throw new Error('Failed to toggle server for client')
    } catch (error) {
      logApiError('POST', url, error)
      throw error
    }
  },

  async generateClientConfig(clientType: string): Promise<object> {
    const url = `${API_BASE_URL}/clients/${clientType}/config`
    logApiCall('GET', url)
    try {
      const response = await fetch(url)
      const data = await response.json()
      logApiResponse('GET', url, response, data)
      if (!response.ok) throw new Error('Failed to generate client config')
      return data
    } catch (error) {
      logApiError('GET', url, error)
      throw error
    }
  }
}