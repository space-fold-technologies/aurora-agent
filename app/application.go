package app

import (
	"github.com/space-fold-technologies/aurora-agent/app/core/configuration"
	"github.com/space-fold-technologies/aurora-agent/app/core/providers"
	"github.com/space-fold-technologies/aurora-agent/app/core/providers/docker"
	"github.com/space-fold-technologies/aurora-agent/app/core/security"
	"github.com/space-fold-technologies/aurora-agent/app/core/server"
	"github.com/space-fold-technologies/aurora-agent/app/core/server/http/controllers"
	"github.com/space-fold-technologies/aurora-agent/app/domain/agents"
)

type ServiceResources struct {
	server       *server.ServerCore
	parameters   configuration.Configuration
	tokenHandler security.TokenHandler
	provider     providers.Provider
}

func ProduceServiceResources(
	server *server.ServerCore,
	parameters configuration.Configuration,
	tokenHandler security.TokenHandler) *ServiceResources {
	return &ServiceResources{
		server:       server,
		parameters:   parameters,
		tokenHandler: tokenHandler,
	}
}

func (sr *ServiceResources) Initialize() {
	sr.provider = sr.providers(sr.parameters.Provider)
	sr.setupControllers(sr.server.GetRegistry())
}

func (sr *ServiceResources) providers(name string) providers.Provider {
	if name == "DOCKER-SWARM" {
		return docker.NewProvider(sr.parameters.AdvertiseAddress, sr.parameters.ListenAddress)
	}
	return nil
}

func (sr *ServiceResources) setupControllers(registry *controllers.HTTPControllerRegistry) {
	//TODO: Register all repositories and inject inject into services and controllers
	registry.AddController(agents.NewController(agents.NewService(sr.provider)))
}
