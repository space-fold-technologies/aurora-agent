package providers

type Provider interface {
	Join(order *JoinOrder) error
	Leave(order *LeaveOrder) error
	ServiceContainers(identifier string) ([]*ContainerDetails, error)
}
