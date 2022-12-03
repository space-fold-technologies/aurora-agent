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

func (as *AgentService) ServiceContainers(identifier string) (*ContainerReport, error) {
	if containers, err := as.provider.ServiceContainers(identifier); err != nil {
		return nil, err
	} else {
		report := &ContainerReport{Containers: make([]*ContainerReport_Container, 0)}
		for _, container := range containers {
			report.Containers = append(report.Containers, &ContainerReport_Container{
				Identifier:        container.ID,
				NodeIdentifier:    container.NodeID,
				TaskIdentifier:    container.TaskID,
				ServiceIdentifier: container.ServiceID,
				IpAddress:         container.IPAddress,
				AddressFamily:     int32(container.AddressFamily),
			})
		}
		return report, nil
	}
}
