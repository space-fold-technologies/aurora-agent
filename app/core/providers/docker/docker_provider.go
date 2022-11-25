package docker

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/space-fold-technologies/aurora-agent/app/core/logging"
	"github.com/space-fold-technologies/aurora-agent/app/core/providers"
)

var (
	MAX_RETRIES = 10
)

type DockerProvider struct {
	dkr           *client.Client
	advertiseAddr string
	listenAddr    string
}

func NewProvider(advertiseAddr, listenAddr string) providers.Provider {
	instance := new(DockerProvider)
	if err := instance.initialize(); err != nil {
		logging.GetInstance().Error(err)
		os.Exit(-1)
	}
	instance.advertiseAddr = advertiseAddr
	instance.listenAddr = listenAddr
	return instance
}

func (dp *DockerProvider) initialize() error {
	if cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation()); err != nil {
		return err
	} else {
		dp.dkr = cli
	}
	return nil
}

func (dp *DockerProvider) Join(order *providers.JoinOrder) error {
	logging.GetInstance().Infof("ADDING NODE WITH IP: %s TO CLUSTER WITH ADDR: %s", dp.advertiseAddr, order.CaptainAddr)
	ctx := context.Background()
	retries := 0
	return dp.join(ctx, order.CaptainAddr, order.Token, &retries)
}

func (dp *DockerProvider) Leave(order *providers.LeaveOrder) error {
	ctx := context.Background()
	defer ctx.Done()
	return dp.dkr.NodeRemove(ctx, order.ID, types.NodeRemoveOptions{Force: true})
}

func (dp *DockerProvider) join(ctx context.Context, captainIP, token string, retries *int) error {
	logging.GetInstance().Infof("WORKER TOKEN: %s WORKER IP: %s CAPTAIN IP: %s", token, dp.advertiseAddr, captainIP)

	if err := dp.dkr.SwarmJoin(ctx, swarm.JoinRequest{
		ListenAddr:    fmt.Sprintf("%s:2377", dp.listenAddr),
		AdvertiseAddr: fmt.Sprintf("%s:2377", dp.advertiseAddr),
		RemoteAddrs:   []string{fmt.Sprintf("%s:2377", captainIP)},
		DataPathAddr:  dp.advertiseAddr,
		JoinToken:     token,
		Availability:  swarm.NodeAvailabilityActive,
	}); err != nil {
		if strings.Contains(err.Error(), "This node is already part of a swarm") {
			logging.GetInstance().Infof("Probably a retry call from: %s", captainIP)
			return nil
		} else if strings.Contains(strings.ToLower(err.Error()), "could not find worker node on time") && *retries < MAX_RETRIES {
			*retries++
			return dp.join(ctx, captainIP, token, retries)
		}
		return err
	}
	return nil
}
