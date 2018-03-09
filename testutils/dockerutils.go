package testutils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/docker/docker/pkg/namesgenerator"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/satori/go.uuid"
)

const (
	propsTokenRepo = "https://github.com/PROPSProject/props-token-distribution.git#docker"
	rabbitmq       = "https://github.com/PROPSProject/dockerfiles.git#master:rabbitmq"
	host           = "127.0.0.1"
)

func getContainerTag(tagPrefix string) string {
	tagPrefix = strings.TrimLeft(tagPrefix, "/")
	return fmt.Sprintf("%v-%v", tagPrefix, uuid.NewV1().String())
}

// StartPropsContainerV1 starts a docker container with ganache-cli and deployed props token contract
func StartPropsContainerV1(acct1, acct2, acct3 string) *chan bool {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	tag := getContainerTag("go-utils_testutils_propscontainer")
	opts := types.ImageBuildOptions{
		Remove:        true,
		Tags:          []string{tag},
		RemoteContext: propsTokenRepo,
	}

	resp, err := cli.ImageBuild(ctx, bytes.NewReader(make([]byte, 1000)), opts)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	log.Println("==================================================== From Docker ====================================================")
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println("====================================================================================================================")

	httpPort, _ := nat.NewPort("tcp", "8000")
	ganachePort, _ := nat.NewPort("tcp", "8545")
	infoSvrPort, _ := nat.NewPort("tcp", "3000")
	ports := nat.PortMap{
		ganachePort: []nat.PortBinding{nat.PortBinding{
			HostIP:   host,
			HostPort: ganachePort.Port(),
		}},
		infoSvrPort: []nat.PortBinding{nat.PortBinding{
			HostIP:   host,
			HostPort: infoSvrPort.Port(),
		}},
	}
	hostConfig := &container.HostConfig{
		PortBindings: ports,
	}
	portSet, _, _ := nat.ParsePortSpecs([]string{httpPort.Port(), ganachePort.Port(), infoSvrPort.Port()})
	containerConfig := &container.Config{
		Image:        tag,
		ExposedPorts: portSet,
		Entrypoint:   []string{"./run.sh", acct1, acct2, acct3},
	}

	createResp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, tag)
	if err != nil {
		panic(err)
	}
	quit := make(chan bool)

	go runContainer(&quit, cli, createResp.ID)

	return &quit
}

// StartRabbitmqContainerV1 starts a docker container with rabbitmq running
func StartRabbitmqContainerV1() *chan bool {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	tag := getContainerTag("rabbitmq")
	opts := types.ImageBuildOptions{
		Remove:        true,
		Tags:          []string{tag},
		RemoteContext: rabbitmq,
		NoCache:       true,
	}

	resp, err := cli.ImageBuild(ctx, bytes.NewReader(make([]byte, 1000)), opts)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	log.Println("==================================================== From Docker ====================================================")
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println("====================================================================================================================")

	rabbitPort, _ := nat.NewPort("tcp", "15672")
	ports := nat.PortMap{
		rabbitPort: []nat.PortBinding{nat.PortBinding{
			HostIP:   host,
			HostPort: rabbitPort.Port(),
		}},
	}
	hostConfig := &container.HostConfig{
		PortBindings: ports,
	}
	portSet, _, _ := nat.ParsePortSpecs([]string{rabbitPort.Port()})
	containerConfig := &container.Config{
		Image:        tag,
		ExposedPorts: portSet,
	}
	createResp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, namesgenerator.GetRandomName(30))
	if err != nil {
		panic(err)
	}

	quit := make(chan bool)

	go runContainer(&quit, cli, createResp.ID)

	return &quit
}

func runContainer(quit *chan bool, cli *client.Client, id string) {
	ctx := context.Background()

	go func(cli *client.Client, id string) {
		log.Println(fmt.Sprintf("Attempting to start container %v", id))
		err := cli.ContainerStart(ctx, id, types.ContainerStartOptions{})
		if err != nil {
			panic(err)
		}
		log.Println(fmt.Sprintf("Container %v started", id))
	}(cli, id)

	out, err := cli.ContainerLogs(ctx, id, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)

	for {

		select {
		case <-*quit:
			log.Println(fmt.Sprintf("Stopping container %v", id))
			err := cli.ContainerStop(ctx, id, nil)
			if err != nil {
				panic(err)
			}

			log.Println(fmt.Sprintf("Removing container %v", id))
			err = cli.ContainerRemove(ctx, id, types.ContainerRemoveOptions{})
			if err != nil {
				panic(err)
			}
		}
	}
}
