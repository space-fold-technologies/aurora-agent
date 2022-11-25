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
