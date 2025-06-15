'use client'

import { useState, useEffect } from 'react'
import { Switch } from '@headlessui/react'
import { 
  ServerIcon, 
  CogIcon, 
  DocumentDuplicateIcon,
  CheckCircleIcon,
  ChevronDownIcon,
  ChevronRightIcon 
} from '@heroicons/react/24/outline'
import clsx from 'clsx'

interface MCPServer {
  id: string
  name: string
  description: string
  url: string
  status: 'active' | 'inactive'
  tools: string[]
}

interface ClientConfig {
  name: string
  displayName: string
  icon: string
  configPath: string
  enabled: boolean
}

const clients: ClientConfig[] = [
  { name: 'claude', displayName: 'Claude Desktop', icon: '🧠', configPath: '~/Library/Application Support/Claude/claude_desktop_config.json', enabled: true },
  { name: 'cursor', displayName: 'Cursor', icon: '↗️', configPath: '~/.cursor/mcp_config.json', enabled: true },
  { name: 'windsurf', displayName: 'Windsurf', icon: '🏄', configPath: '~/.windsurf/mcp_config.json', enabled: true },
  { name: 'cline', displayName: 'Cline', icon: '📋', configPath: '~/.cline/mcp_config.json', enabled: true },
]

export default function Dashboard() {
  const [servers, setServers] = useState<MCPServer[]>([])
  const [clientServerMatrix, setClientServerMatrix] = useState<Record<string, Record<string, boolean>>>({})
  const [loading, setLoading] = useState(true)
  const [expandedServers, setExpandedServers] = useState<Set<string>>(new Set())
  const [clientToolMatrix, setClientToolMatrix] = useState<Record<string, Record<string, Record<string, boolean>>>>({}) // client -> server -> tool -> enabled

  useEffect(() => {
    const fetchData = async () => {
      try {
        // For now, use mock data since the API endpoints might not match exactly
        const mockServers: MCPServer[] = [
          {
            id: '1',
            name: 'github',
            description: 'GitHub integration tools',
            url: 'http://localhost:3000/mcp',
            status: 'active',
            tools: ['create_repo', 'list_repos', 'create_issue']
          },
          {
            id: '2', 
            name: 'calculator',
            description: 'Basic math operations',
            url: 'http://localhost:8000/mcp',
            status: 'active',
            tools: ['add', 'subtract', 'multiply', 'divide']
          },
          {
            id: '3',
            name: 'weather',
            description: 'Weather information',
            url: 'http://localhost:9000/mcp',
            status: 'inactive',
            tools: ['get_weather', 'get_forecast']
          }
        ]

        const initialMatrix: Record<string, Record<string, boolean>> = {}
        const initialToolMatrix: Record<string, Record<string, Record<string, boolean>>> = {}
        
        clients.forEach(client => {
          initialMatrix[client.name] = {}
          initialToolMatrix[client.name] = {}
          mockServers.forEach(server => {
            initialMatrix[client.name][server.id] = client.name === 'claude' // Default: only Claude has all servers enabled
            initialToolMatrix[client.name][server.id] = {}
            server.tools.forEach(tool => {
              initialToolMatrix[client.name][server.id][tool] = client.name === 'claude'
            })
          })
        })

        setServers(mockServers)
        setClientServerMatrix(initialMatrix)
        setClientToolMatrix(initialToolMatrix)
      } catch (error) {
        console.error('Failed to fetch data:', error)
      } finally {
        setLoading(false)
      }
    }
    
    fetchData()
  }, [])

  const toggleServerForClient = (clientName: string, serverId: string) => {
    setClientServerMatrix(prev => ({
      ...prev,
      [clientName]: {
        ...prev[clientName],
        [serverId]: !prev[clientName][serverId]
      }
    }))
  }

  const toggleToolForClient = (clientName: string, serverId: string, toolName: string) => {
    setClientToolMatrix(prev => ({
      ...prev,
      [clientName]: {
        ...prev[clientName],
        [serverId]: {
          ...prev[clientName][serverId],
          [toolName]: !prev[clientName][serverId][toolName]
        }
      }
    }))
  }

  const toggleServerExpansion = (serverId: string) => {
    setExpandedServers(prev => {
      const newSet = new Set(prev)
      if (newSet.has(serverId)) {
        newSet.delete(serverId)
      } else {
        newSet.add(serverId)
      }
      return newSet
    })
  }

  const generateConfig = (clientName: string) => {
    const enabledServers = servers.filter(server => 
      clientServerMatrix[clientName]?.[server.id]
    )
    
    // Generate client-specific config
    const config = {
      mcpServers: enabledServers.reduce((acc, server) => {
        acc[server.name] = {
          url: `http://localhost:8080/mcp/${server.name}`
        }
        return acc
      }, {} as Record<string, any>)
    }

    navigator.clipboard.writeText(JSON.stringify(config, null, 2))
  }

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="animate-spin rounded-full h-32 w-32 border-b-2 border-jungle-500"></div>
      </div>
    )
  }

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      {/* Header */}
      <div className="mb-8">
        <div className="flex items-center space-x-3">
          <div className="text-4xl">🌳</div>
          <div>
            <h1 className="text-3xl font-bold text-gray-900">MCPJungle Dashboard</h1>
            <p className="text-gray-600">Manage MCP servers across your AI clients</p>
          </div>
        </div>
      </div>

      {/* Stats Cards */}
      <div className="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
        <div className="bg-white rounded-lg shadow p-6">
          <div className="flex items-center">
            <ServerIcon className="h-8 w-8 text-jungle-500" />
            <div className="ml-4">
              <p className="text-sm font-medium text-gray-600">Total Servers</p>
              <p className="text-2xl font-bold text-gray-900">{servers.length}</p>
            </div>
          </div>
        </div>
        
        <div className="bg-white rounded-lg shadow p-6">
          <div className="flex items-center">
            <CheckCircleIcon className="h-8 w-8 text-green-500" />
            <div className="ml-4">
              <p className="text-sm font-medium text-gray-600">Active Servers</p>
              <p className="text-2xl font-bold text-gray-900">
                {servers.filter(s => s.status === 'active').length}
              </p>
            </div>
          </div>
        </div>

        <div className="bg-white rounded-lg shadow p-6">
          <div className="flex items-center">
            <CogIcon className="h-8 w-8 text-blue-500" />
            <div className="ml-4">
              <p className="text-sm font-medium text-gray-600">Total Tools</p>
              <p className="text-2xl font-bold text-gray-900">
                {servers.reduce((acc, s) => acc + s.tools.length, 0)}
              </p>
            </div>
          </div>
        </div>

        <div className="bg-white rounded-lg shadow p-6">
          <div className="flex items-center">
            <div className="text-2xl">🤖</div>
            <div className="ml-4">
              <p className="text-sm font-medium text-gray-600">AI Clients</p>
              <p className="text-2xl font-bold text-gray-900">{clients.length}</p>
            </div>
          </div>
        </div>
      </div>

      {/* Server-Client Matrix */}
      <div className="bg-white rounded-lg shadow">
        <div className="px-6 py-4 border-b border-gray-200">
          <h2 className="text-lg font-medium text-gray-900">Server Access Matrix</h2>
          <p className="text-sm text-gray-600">Toggle which servers are available to each AI client</p>
        </div>
        
        <div className="p-6">
          <div className="overflow-x-auto">
            <table className="min-w-full">
              <thead>
                <tr>
                  <th className="text-left text-sm font-medium text-gray-500 pb-4">MCP Server</th>
                  {clients.map(client => (
                    <th key={client.name} className="text-center text-sm font-medium text-gray-500 pb-4 px-4">
                      <div className="flex flex-col items-center space-y-1">
                        <span className="text-2xl">{client.icon}</span>
                        <span>{client.displayName}</span>
                      </div>
                    </th>
                  ))}
                  <th className="text-center text-sm font-medium text-gray-500 pb-4">Actions</th>
                </tr>
              </thead>
              <tbody className="divide-y divide-gray-200">
                {servers.map(server => (
                  <>
                    <tr key={server.id} className="hover:bg-gray-50">
                      <td className="py-4">
                        <div className="flex items-center space-x-3">
                          <div className={clsx(
                            'h-3 w-3 rounded-full',
                            server.status === 'active' ? 'bg-green-400' : 'bg-red-400'
                          )} />
                          <div>
                            <div className="text-sm font-medium text-gray-900">{server.name}</div>
                            <div className="text-sm text-gray-500">{server.description}</div>
                            <div className="text-xs text-gray-400">
                              {server.tools.length} tool{server.tools.length !== 1 ? 's' : ''}
                            </div>
                          </div>
                        </div>
                      </td>
                      {clients.map(client => (
                        <td key={client.name} className="py-4 text-center px-4">
                          <Switch
                            checked={clientServerMatrix[client.name]?.[server.id] || false}
                            onChange={() => toggleServerForClient(client.name, server.id)}
                            className={clsx(
                              clientServerMatrix[client.name]?.[server.id] 
                                ? 'bg-jungle-600' 
                                : 'bg-gray-200',
                              'relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-jungle-500 focus:ring-offset-2'
                            )}
                          >
                            <span
                              className={clsx(
                                clientServerMatrix[client.name]?.[server.id] 
                                  ? 'translate-x-6' 
                                  : 'translate-x-1',
                                'inline-block h-4 w-4 transform rounded-full bg-white transition-transform'
                              )}
                            />
                          </Switch>
                        </td>
                      ))}
                      <td className="py-4 text-center">
                        <button 
                          onClick={() => toggleServerExpansion(server.id)}
                          className="flex items-center space-x-2 text-jungle-600 hover:text-jungle-900 text-sm font-medium mx-auto"
                        >
                          {expandedServers.has(server.id) ? (
                            <ChevronDownIcon className="h-4 w-4" />
                          ) : (
                            <ChevronRightIcon className="h-4 w-4" />
                          )}
                          <span>Configure</span>
                        </button>
                      </td>
                    </tr>
                    
                    {/* Expanded tool configuration */}
                    {expandedServers.has(server.id) && (
                      <tr key={`${server.id}-tools`} className="bg-gray-50">
                        <td className="py-4 pl-12">
                          <div className="space-y-2">
                            <div className="text-sm font-medium text-gray-700 mb-3">Tools in {server.name}:</div>
                            {server.tools.map(tool => (
                              <div key={tool} className="flex items-center space-x-3">
                                <div className="h-2 w-2 rounded-full bg-blue-400" />
                                <span className="text-sm text-gray-600">{tool}</span>
                              </div>
                            ))}
                          </div>
                        </td>
                        {clients.map(client => (
                          <td key={`${client.name}-tools`} className="py-4 text-center px-4">
                            <div className="space-y-3">
                              {server.tools.map(tool => (
                                <div key={tool} className="flex justify-center">
                                  <Switch
                                    checked={clientToolMatrix[client.name]?.[server.id]?.[tool] || false}
                                    onChange={() => toggleToolForClient(client.name, server.id, tool)}
                                    className={clsx(
                                      clientToolMatrix[client.name]?.[server.id]?.[tool] 
                                        ? 'bg-blue-600' 
                                        : 'bg-gray-200',
                                      'relative inline-flex h-5 w-9 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2'
                                    )}
                                  >
                                    <span
                                      className={clsx(
                                        clientToolMatrix[client.name]?.[server.id]?.[tool] 
                                          ? 'translate-x-5' 
                                          : 'translate-x-1',
                                        'inline-block h-3 w-3 transform rounded-full bg-white transition-transform'
                                      )}
                                    />
                                  </Switch>
                                </div>
                              ))}
                            </div>
                          </td>
                        ))}
                        <td className="py-4 text-center">
                          <div className="text-xs text-gray-500">Individual tool toggles</div>
                        </td>
                      </tr>
                    )}
                  </>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>

      {/* Client Configurations */}
      <div className="mt-8 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        {clients.map(client => {
          const enabledCount = servers.filter(server => 
            clientServerMatrix[client.name]?.[server.id]
          ).length
          
          return (
            <div key={client.name} className="bg-white rounded-lg shadow p-6">
              <div className="flex items-center space-x-3 mb-4">
                <span className="text-2xl">{client.icon}</span>
                <div>
                  <h3 className="text-lg font-medium text-gray-900">{client.displayName}</h3>
                  <p className="text-sm text-gray-500">{enabledCount} servers enabled</p>
                </div>
              </div>
              
              <div className="space-y-2">
                <button
                  onClick={() => generateConfig(client.name)}
                  className="w-full flex items-center justify-center space-x-2 bg-jungle-600 text-white px-4 py-2 rounded-md hover:bg-jungle-700 transition-colors"
                >
                  <DocumentDuplicateIcon className="h-4 w-4" />
                  <span>Copy Config</span>
                </button>
                
                <div className="text-xs text-gray-500 truncate" title={client.configPath}>
                  {client.configPath}
                </div>
              </div>
            </div>
          )
        })}
      </div>
    </div>
  )
}