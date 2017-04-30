package qframe_filter_inventory


/******************** Inventory Request
 Sends a query for a key or an IP and provides a back-channel, so that the requesting partner can block on the request
 until it arrives - honouring a timeout...
*/

/*
import (
	"time"
	"github.com/docker/docker/api/types"
)

type InventoryRequest struct {
	Filter interface{}
	Key string
	KeyIsIp bool
	Timeout time.Duration
	Back chan types.ContainerJSON
}

func NewInvReq(filter interface{}, tout time.Duration) InventoryRequest {
	return InventoryRequest{
		Filter: filter,
		Timeout: tout,
		Back: make(chan types.ContainerJSON),
	}
}

func NewInvReqIP(ip string, tout time.Duration, back chan types.ContainerJSON) InventoryRequest {
	return InventoryRequest{
		Filter: NewFilterIP(ip),
		Timeout: tout,
		Back: back,
	}
}

func NewInvReqCntName(name string, tout time.Duration, back chan types.ContainerJSON) InventoryRequest {
	return InventoryRequest{
		Filter: NewFilterCntName(name),
		Timeout: tout,
		Back: back,
	}
}

//********* Filter IP

type FilterIP struct {
	IP string
}

func NewFilterIP(ip string) FilterIP {
	return FilterIP{
		IP: ip,
	}
}

//********* Filter Name

type FilterCntName struct {
	Name string
}

func NewFilterCntName(name string) FilterCntName {
	return FilterCntName{
		Name: name,
	}
}
*/
