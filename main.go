package main

import (
	"time"
	"log"
	"fmt"
	"golang.org/x/net/context"

	"github.com/docker/docker/client"
	dt "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/zpatrick/go-config"

	"github.com/qnib/qframe-types"
	"github.com/qnib/qframe-utils"
	"github.com/qnib/qframe-filter-inventory/lib"
	"github.com/qnib/qframe-collector-docker-events/lib"
	"github.com/qnib/qframe-inventory/lib"
)

const (
	dockerHost = "unix:///var/run/docker.sock"
	dockerAPI = "v1.29"
)

func Run(qChan qtypes.QChan, cfg config.Config, name string) {
	p := qframe_filter_inventory.New(qChan, cfg, name)
	p.Run()
}

func initConfig() (config *container.Config) {
	return &container.Config{Image: "alpine", Volumes: nil, Cmd: []string{"/bin/sleep", "5"}, AttachStdout: false}
}

func hConfig() (config *container.HostConfig) {
	return &container.HostConfig{AutoRemove: true}
}

func startCnt(cli *client.Client, name string, sec int) {
	time.Sleep(time.Duration(sec)*time.Second)
	// Start container
	create, err := cli.ContainerCreate(context.Background(), initConfig(), hConfig(), nil, name)
	if err != nil {
		fmt.Println(err)
	}
	err = cli.ContainerStart(context.Background(), create.ID, dt.ContainerStartOptions{})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	myId := qutils.GetGID()
	qChan := qtypes.NewQChan()
	qChan.Broadcast()
	cfgMap := map[string]string{
		"filter.inventory.inputs": "docker-events",
		"log.level": "debug",
		"filter.inventory.ticker-ms": "2500",
	}

	cfg := config.NewConfig(
		[]config.Provider{
			config.NewStatic(cfgMap),
		},
	)
	// Setup engineCli
	engineCli, err := client.NewClient(dockerHost, dockerAPI, nil, nil)
	if err != nil {
		log.Println("Could not connect to /var/run/docker.sock")
	}
	// Inventory Filter
	p := qframe_filter_inventory.New(qChan, *cfg, "inventory")
	go p.Run()
	// Start docker-events
	pde, _ := qframe_collector_docker_events.New(qChan, *cfg, "docker-events")
	go pde.Run()
	cnts := map[int]string{
		0: fmt.Sprintf("TestCnt1%d", time.Now().Unix()),
		1: fmt.Sprintf("TestCnt2%d", time.Now().Unix()),
	}
	go startCnt(engineCli, cnts[0], 1)
	time.Sleep(time.Second)
	// Create Request to Inventory
	req := qframe_inventory.NewNameContainerRequest(cnts[0])
	p.Log("debug", fmt.Sprintf("SearcRequest for name %s", req.Name))
	p.Inventory.ServeRequest(req)
	req2 := qframe_inventory.NewNameContainerRequest(cnts[1])
	p.Log("debug", fmt.Sprintf("SearcRequest for name %s", req2.Name))
	p.Inventory.ServeRequest(req2)
	// Fire Up second container
	go startCnt(engineCli, cnts[1], 3)
	dc := qChan.Data.Join()
	done := []string{}
	for {
		select {
		case msg := <-dc.Read:
			qm := msg.(qtypes.QMsg)
			if qm.SourceID == myId {
				continue
			}
			p.Log("debug" , fmt.Sprintf("#### Received message on Data-channel: %s\n", qm.Msg))
		case res := <- req.Back:
			p.Log("info", fmt.Sprintf(" SUCCESS > Request: %s (length of PendingPendingRequests: %d)", res.Name, len(p.Inventory.PendingRequests)))
			done = append(done, res.Name)
		case res := <- req2.Back:
			p.Log("info", fmt.Sprintf(" SUCCESS > Request: %s (length of PendingPendingRequests: %d)", res.Name, len(p.Inventory.PendingRequests)))
			done = append(done, res.Name)
		}
		if len(done) == 2 {
			p.Log("debug", fmt.Sprintf("PendingRequests has length: %d", len(p.Inventory.PendingRequests)))
			break
		}
	}

}

