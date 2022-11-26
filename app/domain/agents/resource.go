package agents

import (
	"io"
	"net/http"

	"github.com/space-fold-technologies/aurora-agent/app/core/server/http/controllers"
	"github.com/space-fold-technologies/aurora-agent/app/core/server/registry"
	"google.golang.org/protobuf/proto"
)

const BASE_PATH = "/api/v1/aurora-agent/agents"

type AgentController struct {
	*controllers.ControllerBase
	service *AgentService
}

func NewController(service *AgentService) controllers.HTTPController {
	return &AgentController{service: service}
}

func (nc *AgentController) Name() string {
	return "agent-controller"
}

func (nc *AgentController) Initialize(RouteRegistry registry.RouterRegistry) {
	RouteRegistry.AddRestricted(
		BASE_PATH+"/join",
		[]string{"nodes.create"},
		"PUT",
		nc.join,
	)
	RouteRegistry.AddRestricted(
		BASE_PATH+"/{service-identifier}/containers",
		[]string{""},
		"GET",
		nc.containers,
	)
	RouteRegistry.AddRestricted(
		BASE_PATH+"/leave",
		[]string{"nodes.remove"},
		"PUT",
		nc.leave,
	)
}

func (ac *AgentController) join(w http.ResponseWriter, r *http.Request) {
	order := &RegisterAgent{}
	if data, err := io.ReadAll(r.Body); err != nil {
		ac.BadRequest(w, err)
	} else if err := proto.Unmarshal(data, order); err != nil {
		ac.BadRequest(w, err)
	} else if err := ac.service.Join(order); err != nil {
		ac.ServiceFailure(w, err)
	} else {
		ac.OKNoResponse(w)
	}
}

func (ac *AgentController) leave(w http.ResponseWriter, r *http.Request) {
	order := &RemoveAgent{}
	if data, err := io.ReadAll(r.Body); err != nil {
		ac.BadRequest(w, err)
	} else if err := proto.Unmarshal(data, order); err != nil {
		ac.BadRequest(w, err)
	} else if err := ac.service.Remove(order); err != nil {
		ac.ServiceFailure(w, err)
	} else {
		ac.OKNoResponse(w)
	}
}

func (ac *AgentController) containers(w http.ResponseWriter, r *http.Request) {
	serviceIdentifier := ac.GetVar("service-identifier", r)
	if results, err := ac.service.ServiceContainers(serviceIdentifier); err != nil {
		ac.ServiceFailure(w, err)
	} else {
		ac.OK(w, results)
	}
}
