package providers

type JoinOrder struct {
	Name        string
	Token       string
	CaptainAddr string
}

type LeaveOrder struct {
	ID string
}

type NodeDetails struct {
	ID   string
	Addr string
}

type ContainerDetails struct {
	ID            string
	NodeID        string
	TaskID        string
	ServiceID     string
	IPAddress     string
	AddressFamily uint
}
