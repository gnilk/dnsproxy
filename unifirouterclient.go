package main

import (
	"fmt"
	"log"
	"net"
	"time"

	unifi "github.com/dim13/unifi"
)

type UnifiRouterClient struct {
	config     Router
	controller *unifi.Unifi
	//router *netgear.Client
	//host, port, user, pass string
}

func NewUnifiRouterClient(config Router) RouterClient {
	client := UnifiRouterClient{
		config: config,
	}
	return &client
}

func (client *UnifiRouterClient) Login(host, port, user, pass string) error {
	var err error
	log.Printf("UnifiRouterClient::Login initated to: %s:%s\n", host, port)
	client.controller, err = unifi.Login(user, pass, host, "8443", "dubious-machines.com", 5)
	if err == nil {
		// Save these, as we might need to re-login
		// client.host = host
		// client.port = port
		// client.user = user
		// client.pass = pass
	}
	return err
}

func (client *UnifiRouterClient) AsyncLogin(host, port, user, pass string) error {
	log.Printf("UnifiRouterClient::AsyncLogin, host: %s, port: %s\n", host, port)
	ch := make(chan error, 1)
	go func() {
		err := client.Login(host, port, user, pass)
		ch <- err
	}()

	tout := time.Duration(client.config.TimeoutSec) * time.Second
	if tout < 5 {
		log.Printf("[WARN] Router timeout not specified or too low, setting to 5 seconds\n")
		tout = 5 * time.Second
	}

	select {
	case err := <-ch:
		{
			return err
		}
	case <-time.After(tout):
		{
			return fmt.Errorf("Timeout while connecting to router at host '%s'", client.config.Host)
		}
	}
}

func (client *UnifiRouterClient) GetAttachedDeviceList() ([]RouterDevice, error) {

	if client.controller == nil {
		log.Printf("[INFO] Unifi router disconnected, re-login initated\n")
		//err := client.Login(client.host, client.port, client.user, client.pass)
		log.Printf("UnifiRouterClient::GetAttachedDeviceList, client.config.host: %s\n", client.config.Host)
		err := client.AsyncLogin(client.config.Host, client.config.Port, client.config.User, client.config.Password)
		if err != nil {
			return nil, err
		}
		log.Printf("[INFO] Logged in to Unifi router\n")
	}

	site, err := client.controller.Site("dubious-machines.com")
	if err != nil {
		client.controller = nil
		return nil, err
	}
	devices, err := client.controller.DeviceMap(site)
	if err != nil {
		client.controller = nil
		return nil, err
	}

	sta, err := client.controller.Sta(site)
	if err != nil {
		client.controller = nil
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
