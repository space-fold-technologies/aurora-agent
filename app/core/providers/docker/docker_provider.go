package docker

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
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
	logger := logging.GetInstance()
	logger.Infof("ADDING NODE WITH IP: %s TO CLUSTER WITH ADDR: %s", dp.advertiseAddr, order.CaptainAddr)
	ctx := context.Background()
	defer ctx.Done()
	retries := 0
	return dp.join(ctx, order.CaptainAddr, order.Token, &retries)
}

func (dp *DockerProvider) Leave(order *providers.LeaveOrder) error {
	ctx := context.Background()
	defer ctx.Done()
	return dp.dkr.NodeRemove(ctx, order.ID, types.NodeRemoveOptions{Force: true})
}

func (dp *DockerProvider) ServiceContainers(identifier string) ([]*providers.ContainerDetails, error) {
	ctx := context.Background()
	defer ctx.Done()
	retries := 0
	if containers, err := dp.queryContainers(ctx, identifier, &retries); err != nil {
		return nil, err
	} else {
		return containers, nil
	}
}

func (dp *DockerProvider) join(ctx context.Context, captainIP, token string, retries *int) error {
	logger := logging.GetInstance()
	logger.Infof("WORKER TOKEN: %s WORKER IP: %s CAPTAIN IP: %s", token, dp.advertiseAddr, captainIP)
	cmd := fmt.Sprintf("docker swarm join --token %s %s:2377", token, captainIP)
	if _, err := exec.Command("/bin/sh", "-c", cmd).Output(); err != nil {
		return err
	}
	//TODO: Closing off this section of the code in favor of using the cli command for correct network set up
	/*
		if err := dp.dkr.SwarmJoin(ctx, swarm.JoinRequest{
			ListenAddr:    fmt.Sprintf("%s:2377", dp.listenAddr),
			AdvertiseAddr: fmt.Sprintf("%s:2377", dp.advertiseAddr),
			RemoteAddrs:   []string{fmt.Sprintf("%s:2377", captainIP)},
			DataPathAddr:  dp.advertiseAddr,
			JoinToken:     token,
			Availability:  swarm.NodeAvailabilityActive,
		}); err != nil {
			if strings.Contains(err.Error(), "This node is already part of a swarm") {
				logger.Infof("Probably a retry call from: %s", captainIP)
				return nil
			} else if strings.Contains(strings.ToLower(err.Error()), "could not find worker node on time") && *retries < MAX_RETRIES {
				*retries++
				return dp.join(ctx, captainIP, token, retries)
			}
			return err
		}
	*/
	return nil
}

func (dp *DockerProvider) queryContainers(ctx context.Context, identifier string, retries *int) ([]*providers.ContainerDetails, error) {
	details := make([]*providers.ContainerDetails, 0)
	filter := filters.NewArgs()
	filter.Add("label", fmt.Sprintf("com.docker.swarm.service.id=%s", identifier))
	if containers, err := dp.dkr.ContainerList(ctx, types.ContainerListOptions{Filters: filter}); err != nil {
		return nil, err
	} else if len(containers) == 0 && *retries < MAX_RETRIES {
		time.Sleep(5 * time.Second)
		*retries++
		return dp.queryContainers(ctx, identifier, retries)
	} else {
		for _, container := range containers {
			details = append(details, &providers.ContainerDetails{
				ID:            container.ID,
				NodeID:        container.Labels["com.docker.swarm.node.id"],
				TaskID:        container.Labels["com.docker.swarm.task.id"],
				ServiceID:     container.Labels["com.docker.swarm.service.id"],
				IPAddress:     container.NetworkSettings.Networks["aurora-default"].IPAddress,
				AddressFamily: 4,
			})
		}
		return details, nil
	}
}
