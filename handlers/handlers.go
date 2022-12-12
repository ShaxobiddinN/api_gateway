package handlers

import (
	"blogpost/api_gateway/clients"
	"blogpost/api_gateway/config"
)

// Handler...
type handler struct {
	cfg         config.Config
	grpcClients *clients.GrpcClients
}

// NewHandler ...
func NewHandler(cfg config.Config, grpcClients *clients.GrpcClients) handler {
	return handler{
		cfg:         cfg,
		grpcClients: grpcClients,
	}
}
