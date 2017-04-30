package qframe_filter_inventory

/*
import "github.com/docker/docker/api/types"

//***** Example Interface String
//To use the qframe-inventory the Struct has to implement Equal()


type ContainerRequest struct {
	Name string
	ID string
	IPs []string
}

func (this ContainerRequest) Equal(other types.ContainerJSON) bool {
	matchIP := false
	if len(this.IPs) != 0 && other.NetworkSettings.Networks != nil {
		for _, net := range other.NetworkSettings.Networks {
			for _, ip := range this.IPs {
				if ip == net.IPAddress {
					matchIP = true
				}
			}
		}
	}
	return this.ID == other.ID || this.Name == other.Name || matchIP
}
*/
