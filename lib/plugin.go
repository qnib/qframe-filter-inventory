package qframe_filter_inventory

import (
	"C"
	"fmt"
	"time"
	"reflect"

	"github.com/qnib/qframe-types"
	"github.com/qnib/qframe-utils"
	"github.com/qnib/qframe-inventory/lib"
	"github.com/zpatrick/go-config"
)

const (
	version = "0.1.1"
	pluginTyp = qtypes.FILTER
	pluginPkg = "inventory"
)

type Plugin struct {
	qtypes.Plugin
	Inventory qframe_inventory.Inventory
}

func New(qChan qtypes.QChan, cfg config.Config, name string) Plugin {
	return Plugin{
		Plugin: qtypes.NewNamedPlugin(qChan, cfg, pluginTyp, pluginPkg, name, version),
		Inventory: qframe_inventory.NewInventory(),
	}
}

// Run fetches everything from the Data channel and flushes it to stdout
func (p *Plugin) Run() {
	p.Log("info", fmt.Sprintf("Start inventory v%s", p.Version))
	myId := qutils.GetGID()
	dc := p.QChan.Data.Join()
	tickerTime := p.CfgIntOr("ticker-ms", 2500)
	ticker := time.NewTicker(time.Millisecond * time.Duration(tickerTime)).C
	for {
		select {
		case val := <-dc.Read:
			switch val.(type) {
			case qtypes.QMsg:
				qm := val.(qtypes.QMsg)
				if qm.SourceID == myId {
					continue
				}
				switch qm.Data.(type) {
				case qtypes.ContainerEvent:
					ce := qm.Data.(qtypes.ContainerEvent)
					p.Log("debug", fmt.Sprintf("Received Event: %s.%s",ce.Event.Type, ce.Event.Action))
					if ce.Event.Type == "container" && ce.Event.Action == "start" {
						p.Inventory.SetItem(ce.Container.ID, ce.Container)
					}
				default:
					p.Log("debug", fmt.Sprintf("Received qm.Data: %s", reflect.TypeOf(qm.Data)))
				}
			case qframe_inventory.ContainerRequest:
				req := val.(qframe_inventory.ContainerRequest)
				p.Log("info", fmt.Sprintf("Received InventoryRequest for %v", req))
				p.Inventory.ServeRequest(req)
			}
		case <- ticker:
			p.Log("debug", "Ticker came along: p.Inventory.CheckRequests()")
			p.Inventory.CheckRequests()
			continue
		}
	}
}
