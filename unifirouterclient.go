package main

import (
	"net"

	unifi "github.com/dim13/unifi"
)

type UnifiRouterClient struct {
	controller *unifi.Unifi
	//router *netgear.Client
}

func NewUnifiRouterClient() RouterClient {
	client := UnifiRouterClient{}
	return &client
}

func (client *UnifiRouterClient) Login(host, port, user, pass string) error {
	var err error
	client.controller, err = unifi.Login(user, pass, host, "8443", "dubious-machines.com", 5)
	return err
}
func (client *UnifiRouterClient) GetAttachedDeviceList() ([]RouterDevice, error) {

	site, err := client.controller.Site("dubious-machines.com")
	if err != nil {
		return nil, err
	}
	devices, err := client.controller.DeviceMap(site)
	if err != nil {
		return nil, err
	}

	sta, err := client.controller.Sta(site)
	if err != nil {
		return nil, err
	}

	return client.transformDevices(sta, devices)
}

func (client *UnifiRouterClient) transformDevices(sta []unifi.Sta, devices unifi.DeviceMap) ([]RouterDevice, error) {
	rdlist := make([]RouterDevice, 0)

	for _, s := range sta {
		deviceMac := ""

		if s.ApMac != "" {
			deviceMac = s.ApMac
		} else if s.SwMac != "" {
			deviceMac = s.SwMac
		}
		//deviceName := devices[deviceMac].DeviceName()

		macaddr, err := net.ParseMAC(deviceMac)
		if err != nil {
			macaddr, _ = net.ParseMAC("01:23:45:67:89:ab")
		}

		rd := RouterDevice{
			IP:     net.ParseIP(s.IP),
			Name:   s.Name(),
			MAC:    macaddr,
			Signal: s.Signal,
		}
		rdlist = append(rdlist, rd)
	}
	return rdlist, nil
}
