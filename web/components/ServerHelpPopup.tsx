'use client'

import { Fragment, useState } from 'react'
import { Dialog, Transition } from '@headlessui/react'
import { XMarkIcon, QuestionMarkCircleIcon } from '@heroicons/react/24/outline'

interface ServerHelpPopupProps {
  isOpen: boolean
  onClose: () => void
}

export default function ServerHelpPopup({ isOpen, onClose }: ServerHelpPopupProps) {
  const [activeTab, setActiveTab] = useState('stdio')

  const examples = {
    stdio: {
      title: 'Standard Input/Output (Stdio)',
      description: 'Direct process communication using stdin/stdout. Most common for local MCP servers.',
      examples: [
        {
          name: 'Python MCP Server',
          command: 'python',
          args: ['-m', 'mcp_server'],
          environment: { PYTHONPATH: '/path/to/server' },
          description: 'Run a Python-based MCP server module'
        },
        {
          name: 'Node.js MCP Server',
          command: 'node',
          args: ['server.js'],
          environment: { NODE_ENV: 'production' },
          description: 'Execute a Node.js MCP server script'
        },
        {
          name: 'Go Binary MCP Server',
          command: '/usr/local/bin/mcp-server',
          args: ['--config', '/etc/mcp/config.json'],
          environment: { LOG_LEVEL: 'info' },
          description: 'Run a compiled Go MCP server binary'
        }
      ]
    },
    http: {
      title: 'HTTP Transport',
      description: 'Streamable HTTP transport for remote MCP servers over HTTP/HTTPS.',
      examples: [
        {
          name: 'Local HTTP Server',
          url: 'http://localhost:8080/mcp/server',
          bearer_token: '',
          description: 'Connect to a local HTTP MCP server'
        },
        {
          name: 'Remote HTTPS Server',
          url: 'https://api.example.com/mcp/v1',
          bearer_token: 'your-auth-token-here',
          description: 'Connect to a remote HTTPS MCP server with authentication'
        },
        {
          name: 'Docker Container',
          url: 'http://mcp-container:3000/mcp',
          bearer_token: '',
          description: 'Connect to MCP server running in Docker container'
        }
      ]
    },
    sse: {
      title: 'Server-Sent Events (SSE)',
      description: 'Real-time communication using Server-Sent Events over HTTP.',
      examples: [
        {
          name: 'SSE Stream Server',
          url: 'http://localhost:8080/events',
          bearer_token: '',
          description: 'Connect to SSE-based MCP server for real-time updates'
        },
        {
          name: 'Authenticated SSE',
          url: 'https://stream.example.com/mcp/events',
          bearer_token: 'sse-auth-token',
          description: 'Connect to authenticated SSE MCP server'
        },
        {
          name: 'WebSocket Alternative',
          url: 'http://realtime.local:9000/sse',
          bearer_token: '',
          description: 'SSE as lightweight alternative to WebSocket'
        }
      ]
    }
  }

  return (
    <Transition.Root show={isOpen} as={Fragment}>
      <Dialog as="div" className="relative z-50" onClose={onClose}>
        <Transition.Child
          as={Fragment}
          enter="ease-out duration-300"
          enterFrom="opacity-0"
          enterTo="opacity-100"
          leave="ease-in duration-200"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
        >
          <div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
        </Transition.Child>

        <div className="fixed inset-0 z-10 overflow-y-auto">
          <div className="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
            <Transition.Child
              as={Fragment}
              enter="ease-out duration-300"
              enterFrom="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
              enterTo="opacity-100 translate-y-0 sm:scale-100"
              leave="ease-in duration-200"
              leaveFrom="opacity-100 translate-y-0 sm:scale-100"
              leaveTo="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
            >
              <Dialog.Panel className="relative transform overflow-hidden rounded-lg bg-white px-4 pb-4 pt-5 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-4xl sm:p-6">
                <div className="absolute right-0 top-0 hidden pr-4 pt-4 sm:block">
                  <button
                    type="button"
                    className="rounded-md bg-white text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-jungle-500 focus:ring-offset-2"
                    onClick={onClose}
                  >
                    <span className="sr-only">Close</span>
                    <XMarkIcon className="h-6 w-6" aria-hidden="true" />
                  </button>
                </div>
                
                <div className="sm:flex sm:items-start">
                  <div className="mt-3 text-center sm:ml-4 sm:mt-0 sm:text-left w-full">
                    <Dialog.Title as="h3" className="text-lg font-semibold leading-6 text-gray-900 mb-4">
                      MCP Server Configuration Guide
                    </Dialog.Title>
                    
                    <div className="mb-6">
                      <nav className="flex space-x-8" aria-label="Tabs">
                        {Object.entries(examples).map(([key, config]) => (
                          <button
                            key={key}
                            onClick={() => setActiveTab(key)}
                            className={`whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm ${
                              activeTab === key
                                ? 'border-jungle-500 text-jungle-600'
                                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                            }`}
                          >
                            {config.title}
                          </button>
                        ))}
                      </nav>
                    </div>

                    <div className="space-y-6">
                      <div>
                        <h4 className="text-base font-medium text-gray-900 mb-2">
                          {examples[activeTab as keyof typeof examples].title}
                        </h4>
                        <p className="text-sm text-gray-600 mb-4">
                          {examples[activeTab as keyof typeof examples].description}
                        </p>
                      </div>

                      <div className="space-y-4">
                        <h5 className="text-sm font-medium text-gray-900">Configuration Examples:</h5>
                        
                        {examples[activeTab as keyof typeof examples].examples.map((example, index) => (
                          <div key={index} className="border border-gray-200 rounded-lg p-4 bg-gray-50">
                            <h6 className="font-medium text-gray-900 mb-2">{example.name}</h6>
                            <p className="text-sm text-gray-600 mb-3">{example.description}</p>
                            
                            <div className="space-y-2 text-sm">
                              {activeTab === 'stdio' && (
                                <>
                                  <div className="flex items-center">
                                    <span className="font-medium text-gray-700 w-20">Command:</span>
                                    <code className="bg-white px-2 py-1 rounded border text-jungle-600">
                                      {(example as any).command}
                                    </code>
                                  </div>
                                  {(example as any).args && (example as any).args.length > 0 && (
                                    <div className="flex items-start">
                                      <span className="font-medium text-gray-700 w-20 mt-1">Args:</span>
                                      <div className="flex flex-wrap gap-1">
                                        {(example as any).args.map((arg: string, argIndex: number) => (
                                          <code key={argIndex} className="bg-white px-2 py-1 rounded border text-jungle-600">
                                            {arg}
                                          </code>
                                        ))}
                                      </div>
                                    </div>
                                  )}
                                  {(example as any).environment && Object.keys((example as any).environment).length > 0 && (
                                    <div className="flex items-start">
                                      <span className="font-medium text-gray-700 w-20 mt-1">Env:</span>
                                      <div className="space-y-1">
                                        {Object.entries((example as any).environment).map(([key, value]) => (
                                          <code key={key} className="block bg-white px-2 py-1 rounded border text-jungle-600">
                                            {key}={value as string}
                                          </code>
                                        ))}
                                      </div>
                                    </div>
                                  )}
                                </>
                              )}
                              
                              {(activeTab === 'http' || activeTab === 'sse') && (
                                <>
                                  <div className="flex items-center">
                                    <span className="font-medium text-gray-700 w-20">URL:</span>
                                    <code className="bg-white px-2 py-1 rounded border text-jungle-600">
                                      {(example as any).url}
                                    </code>
                                  </div>
                                  {(example as any).bearer_token && (
                                    <div className="flex items-center">
                                      <span className="font-medium text-gray-700 w-20">Token:</span>
                                      <code className="bg-white px-2 py-1 rounded border text-jungle-600">
                                        {(example as any).bearer_token}
                                      </code>
                                    </div>
                                  )}
                                </>
                              )}
                            </div>
                          </div>
                        ))}
                      </div>
                      
                      <div className="bg-blue-50 border border-blue-200 rounded-lg p-4">
                        <h6 className="font-medium text-blue-900 mb-2">💡 Tips for {examples[activeTab as keyof typeof examples].title}:</h6>
                        <ul className="text-sm text-blue-800 space-y-1">
                          {activeTab === 'stdio' && (
                            <>
                              <li>• Ensure the command is in your PATH or use absolute path</li>
                              <li>• Test your command manually first to verify it works</li>
                              <li>• Use environment variables for configuration</li>
                              <li>• Check server logs if connection fails</li>
                            </>
                          )}
                          {activeTab === 'http' && (
                            <>
                              <li>• Use HTTPS in production for security</li>
                              <li>• Ensure the server supports MCP over HTTP</li>
                              <li>• Test connectivity with curl first</li>
                              <li>• Keep bearer tokens secure and rotate regularly</li>
                            </>
                          )}
                          {activeTab === 'sse' && (
                            <>
                              <li>• SSE provides real-time updates from server</li>
                              <li>• Ensure your server implements SSE properly</li>
                              <li>• Check CORS settings if connecting from browser</li>
                              <li>• Monitor connection stability for long-running streams</li>
                            </>
                          )}
                        </ul>
                      </div>
                    </div>
                    
                    <div className="mt-8 flex justify-end">
                      <button
                        type="button"
                        className="inline-flex justify-center rounded-md bg-jungle-600 px-4 py-2 text-sm font-semibold text-white shadow-sm hover:bg-jungle-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-jungle-600"
                        onClick={onClose}
                      >
                        Got it, thanks!
                      </button>
                    </div>
                  </div>
                </div>
              </Dialog.Panel>
            </Transition.Child>
          </div>
        </div>
      </Dialog>
    </Transition.Root>
  )
}

export function ServerHelpButton({ onClick }: { onClick: () => void }) {
  return (
    <button
      type="button"
      onClick={onClick}
      className="inline-flex items-center justify-center rounded-full p-1 text-gray-400 hover:text-jungle-600 hover:bg-jungle-50 focus:outline-none focus:ring-2 focus:ring-jungle-500 focus:ring-offset-2"
      title="Server configuration help"
    >
      <QuestionMarkCircleIcon className="h-5 w-5" aria-hidden="true" />
    </button>
  )
}