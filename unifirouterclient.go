package main

import (
	"log"
	"net"

	unifi "github.com/dim13/unifi"
)

type UnifiRouterClient struct {
	controller *unifi.Unifi
	//router *netgear.Client
	host, port, user, pass string
}

func NewUnifiRouterClient() RouterClient {
	client := UnifiRouterClient{}
	return &client
}

func (client *UnifiRouterClient) Login(host, port, user, pass string) error {
	var err error
	client.controller, err = unifi.Login(user, pass, host, "8443", "dubious-machines.com", 5)
	if err == nil {
		// Save these, as we might need to re-login
		client.host = host
		client.port = port
		client.user = user
		client.pass = pass
	}
	return err
}

func (client *UnifiRouterClient) GetAttachedDeviceList() ([]RouterDevice, error) {

	if unifi.Connected != unifi.Connected {
		log.Printf("unifi router disconnected, re-login initated\n")
		err := client.Login(client.host, client.port, client.user, client.pass)
		if err != nil {
			return nil, err
		}
		log.Printf("Logged in to Unifi router\n")
	}

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
