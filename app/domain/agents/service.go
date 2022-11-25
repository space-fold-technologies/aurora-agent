package agents

import "github.com/space-fold-technologies/aurora-agent/app/core/providers"

type AgentService struct {
	provider providers.Provider
}

func NewService(provider providers.Provider) *AgentService {
	instance := new(AgentService)
	instance.provider = provider
	return instance
}

func (as *AgentService) Join(order *RegisterAgent) error {
	return as.provider.Join(&providers.JoinOrder{
		Name:        order.GetName(),
		Token:       order.GetToken(),
		CaptainAddr: order.GetAddress(),
	})
}

func (as *AgentService) Remove(order *RemoveAgent) error {
	return as.provider.Leave(&providers.LeaveOrder{ID: order.GetId()})
}
