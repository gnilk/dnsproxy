package main

import (
	netgear "github.com/EvanPurkhiser/netgear-go"
)

type NetGearRouterClient struct {
	router *netgear.Client
}

func NewNetGearRouterClient() RouterClient {
	client := NetGearRouterClient{}
	return &client
}

func (client *NetGearRouterClient) Login(host, port, user, pass string) error {
	client.router = netgear.NewClient(host, user, pass)
	return client.router.Login()
}
func (client *NetGearRouterClient) GetAttachedDeviceList() ([]RouterDevice, error) {
	devices, err := client.router.Devices()
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
			Name:     d.Name,
			MAC:      d.MAC,
			Type:     d.Type,
			LinkRate: d.LinkRate,
			Signal:   d.Signal,
		}
		rdlist = append(rdlist, rd)
	}
	return rdlist, nil
}
