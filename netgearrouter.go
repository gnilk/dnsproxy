// deprecated!
package main

import (
	netgear "github.com/DRuggeri/netgear_client"
	"strings"
)

type NetGearRouterClient struct {
	config *Router
	router *netgear.NetgearClient
}

func NewNetGearRouterClient(config *Router) RouterClient {
	client := NetGearRouterClient{
		config: config,
	}
	client.Login(config.Host, config.Port, config.User, config.Password)
	return &client
}

func (client *NetGearRouterClient) Login(host, port, user, pass string) error {
	client.router, _ = netgear.NewNetgearClient(host, true, user, pass)
	return client.router.LogIn()
}
func (client *NetGearRouterClient) GetAttachedDeviceList() ([]RouterDevice, error) {
	devices, err := client.router.G
	if err != nil {
		return nil, err
	}

	return client.transformDevices(devices)
}

func (client *NetGearRouterClient) transformDevices(devices []netgear.AttachedDevice) ([]RouterDevice, error) {
	rdlist := make([]RouterDevice, 0)
	for _, d := range devices {

		// Replace unknown's with 'Mac' addresses
		if d.Name == "<unknown>" {
			d.Name = d.MAC.String()
		}

		rd := RouterDevice{
			IP:       d.IP,
			Name:     strings.ToLower(d.Name),
			MAC:      d.MAC,
			Type:     d.Type,
			LinkRate: d.LinkRate,
			Signal:   d.Signal,
		}
		rdlist = append(rdlist, rd)
	}
	return rdlist, nil
}
