package server

import "github.com/mark3labs/mcp-go/server"

// This file provides the expected API for main.go by acting as an
// adapter to the implementations in server.go

// NewService creates a new MCP service for gopls integration
func NewService(s *server.MCPServer) (*Service, error) {
	logFile, err := setupLogger()
	if err != nil {
		return nil, err
	}

	svc := &Service{
		logFile: logFile,
	}

	if err := svc.initLSPClient(); err != nil {
		if logFile != nil {
			logFile.Close()
		}
		return nil, err
	}

	if s != nil {
		svc.server = s
	} else {
		svc.server = setupServer()
	}
	return svc, nil
}
